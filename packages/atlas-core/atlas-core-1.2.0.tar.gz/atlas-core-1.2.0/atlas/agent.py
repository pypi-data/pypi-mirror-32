import os, logging
from .version import __version__
from .client import AgentClient
from .utils import generate_hash
from atlas_sdk import BrokerConfig
from atlas_sdk.request import SID_KEY, UID_KEY, ENV_KEY, VERSION_KEY, LANG_KEY, \
  CID_KEY
from .interpreters import Interpreter
from transitions import Machine, EventData, MachineError
from fuzzywuzzy import process

PREFIX_ATLAS = 'atlas/'
STATE_ASLEEP = '%sasleep' % PREFIX_ATLAS
STATE_ASK = '%sask' % PREFIX_ATLAS
STATE_CANCEL = '%scancel' % PREFIX_ATLAS
STATE_FALLBACK = '%sfallback' % PREFIX_ATLAS

def is_builtin(state):
  """Checks if the given state is a builtin one.

  :param state: State to check
  :type state: str
  :rtype: bool

  """

  return state.startswith(PREFIX_ATLAS)

def resolve_parametric_intent_name(name, slots):
  """Format the intent name with slots dict to handle parametric intent name.

  With this feature, you can create generic intent and output a single intent name
  based on slots values.

  :param name: Name of the intent, may contain placeholders such as {slotName}
  :type name: str
  :param slots: Dictionary of slots as retrieved by the NLU
  :type slots: dict

  """

  # TODO what to do if a slot value is an array?

  slots_flattened = { k: v[0].get('value') for (k, v) in slots.items() }

  try:
    return name.format(**slots_flattened)
  except KeyError:
    return name

try:
  # Try to load the GraphMachine
  import pygraphviz
  from transitions.extensions import GraphMachine as Machine
except:
  pass

class Agent:
  """Agents are the core of the dialog engine of atlas.

  They maintain the state of a conversation and act as a relay between skills and channels. They are
  tied to a particular user or device and open a MQTT client themselve.

  """

  def __init__(self, id, uid, interpreter, env, validate_intent=None, data_path=None):
    """Creates a new agent.

    :param id: Channel id
    :type id: str
    :param uid: User ID
    :type uid: str
    :param interpreter: Interpreter to be used
    :type interpreter: Interpreter
    :param env: Configuration parameters for the user
    :type env: dict
    :param validate_intent: Handler to validate the user intent by checking if it's available, 
                            it will take the intent name and should returns a dict containing valid env 
                            variables if the skill exists or None if no skill could be found matching
                            that intent
    :type validate_intent: callable
    :param data_path: Where to store generated data such as the graph
    :type data_path: str

    """

    self._log = logging.getLogger('atlas.agent.%s' % id)
    self._intent_queue = []

    self.validate_intent = validate_intent
    self.interpreter = interpreter
    self.env = env
    self.id = id
    self.uid = uid

    # Configure the client facade

    self._client = AgentClient(id, self.interpreter.lang(),
      on_parse=self.parse,
      on_ask=self.ask,
      on_terminate=self._terminate_from_skill,
      on_show=self.show
    )

    self.reset()

    # Constructs every possible transitions from interpreter metadata

    metadata = { k: v for k, v in self.interpreter.metadata().items() if not is_builtin(k) }
    metadata_states = list(metadata.keys()) + [STATE_FALLBACK]

    states = [STATE_ASLEEP, STATE_CANCEL, STATE_ASK] + metadata_states

    self._machine = Machine(self, 
      states=states, 
      initial=STATE_ASLEEP, 
      send_event=True, 
      after_state_change=self._format_transition)

    self._log.info('Created with states %s' % list(self._machine.states.keys()))

    self._machine.add_transition(STATE_ASLEEP, 
      [STATE_CANCEL] + metadata_states, STATE_ASLEEP, after=self.reset)
    self._machine.add_transition(STATE_CANCEL, 
      [STATE_ASK] + metadata_states, STATE_CANCEL, after=self._call_intent)
    self._machine.add_transition(STATE_ASK, metadata_states, STATE_ASK, after=self._on_ask)

    for intent in metadata_states:
      self._machine.add_transition(intent, [STATE_ASLEEP, STATE_ASK], intent, after=self._call_intent)

    if data_path:
      try:
        self.get_graph().draw(os.path.join(data_path, '%s.png' % self.uid), prog='dot') # pylint: disable=E1101
      except:
        self._log.warning("Could not draw the transitions graph, maybe you'll need pygraphviz installed!")

  def _format_transition(self, e):
    dest = e.transition.dest
    msg = '⚡ %s: %s -> %s' % (e.event.name, e.transition.source, dest)

    if dest == STATE_ASK:
      msg += ' (slot: %s, choices: %s)' % (self._cur_asked_slot, self._cur_choices)

    self._log.info(msg)

  def reset(self, event=None):
    """Resets the current agent states.

    :param event: Optional event parameters
    :type event: EventData

    """

    self._cur_asked_slot = None
    self._cur_intent = None
    self._cur_conversation_id = None
    self._cur_slots = {}
    self._cur_choices = None

    self._client.terminate()

    self._process_next_intent()

  def go(self, trigger_name, **kwargs):
    """Safely call a trigger and catch errors

    :param trigger_name: Name of the trigger
    :type trigger_name: str

    """

    try:
      self.trigger(trigger_name, **kwargs) # pylint: disable=E1101
    except Exception as err:
      self._log.error('Could not trigger "%s": %s' % (trigger_name, err))

  def _call_intent(self, event):
    """Call the intent with current slot values.

    :param event: Machine event
    :type event: EventData

    """

    dest = event.transition.dest

    # Generates a new conversation id if needed
    if dest != self._cur_intent:
      self._cur_conversation_id = generate_hash()
      self._log.info('💬 New "%s" conversation started with id %s' % (dest, self._cur_conversation_id))

    # Here we need to keep track of the original intent name otherwise transitions
    # will be broke
    self._cur_intent = dest

    # But we use the resolved name for the intent validation to check if a skill can
    # take care of it
    resolved_intent_name = resolve_parametric_intent_name(self._cur_intent, self._cur_slots)
    valid_keys = self.env.keys()

    # Try to validate if the intent could be reached

    if self.validate_intent:
      valid_keys = self.validate_intent(resolved_intent_name)

      if valid_keys == None:
        self._log.warn('Intent %s could not be reached, skipping now' % resolved_intent_name)
        return self.go(STATE_ASLEEP)
    
    self._client.work()

    # Constructs the message payload
    data = {
      CID_KEY: self._cur_conversation_id,
      SID_KEY: self.id,
      UID_KEY: self.uid,
      LANG_KEY: self.interpreter.lang(),
      VERSION_KEY: __version__,
      ENV_KEY: { k: self.env[k] for k in valid_keys },
    }

    data.update(self._cur_slots)

    self._log.debug('Calling resolved intent "%s (%s)" with params %s' % (resolved_intent_name, self._cur_intent, data))

    self._client.intent(resolved_intent_name, data)

  def _on_ask(self, event):
    """Entered in ask state, save current asked param.

    :param event: Machine event
    :type event: EventData

    """

    self._cur_asked_slot = event.kwargs.get('slot')
    self._cur_choices = event.kwargs.get('choices')

    payload = event.kwargs.get('payload')

    self._log.debug('Asking request with payload %s' % payload)

    self._client.ask(payload)

  def extract_choices(self, values):
    """Extract choices from given values. It will use fuzzy match to determine
    the first valid one.

    :param values: Values to use when searching
    :type values: list
    :rtype: list

    """

    # No choices available, just return the values
    if not self._cur_choices:
      return values

    # Else, try to find one value matching with current choices
    for value in values:
      # Retrieve the inner slot value
      slot_value = value.get('value')

      if slot_value:
        match = process.extractOne(slot_value, self._cur_choices, score_cutoff=60)

        # If match, update the value prop for this slot but keep other informations
        if match:
          return [{ **value, 'value': match[0] }]

    return None

  def parse(self, msg):
    """Parse a raw message.

    :param msg: Message to parse
    :type msg: str

    """

    self._log.debug('Parsing "%s"' % msg)

    data = self.interpreter.parse(msg)
    data_without_cancel = [d for d in data if d['intent'] != STATE_CANCEL]

    # If intent is "cancel", go to the cancel state immediately. This will generates a new
    # conversation id so old request will be dismissed if they tried to do something
    if len(data_without_cancel) != len(data) and self.state != STATE_ASLEEP: # pylint: disable=E1101
      self.go(STATE_CANCEL)
    else:
      if self.state == STATE_ASK and self._cur_asked_slot: # pylint: disable=E1101
        values = self.interpreter.parse_entity(msg, self._cur_intent, self._cur_asked_slot)

        matched_values = self.extract_choices(values)

        if matched_values:
          self._cur_slots[self._cur_asked_slot] = matched_values
          self.go(self._cur_intent)
        else:
          self._log.warning('Not a valid input "%s", choices are %s' % (msg, self._cur_choices))
      else:
        self._intent_queue.extend(data_without_cancel)

        if self.state == STATE_ASLEEP: # pylint: disable=E1101
          if len(data_without_cancel) == 0:
            self._cur_slots = { 'text': [{ 'value': msg }] }

            self.go(STATE_FALLBACK)
          else:
            self._process_next_intent()

  def _process_next_intent(self):
    """Process the intent queue if any left.
    """

    if len(self._intent_queue) > 0:
      intent = self._intent_queue.pop(0)
      self._cur_slots = intent['slots']
      self.go(intent['intent'])

  def _is_valid_request(self, data, required_keys=[]):
    """Checks if a request is valid in the current context.

    :param data: Request data
    :type data: dict
    :param required_keys: Keys required in the data to validate the request
    :type required_keys: list

    """

    conversation_id = data.get(CID_KEY)

    if conversation_id == None or conversation_id != self._cur_conversation_id:
      self._log.warning('Wrong conversation id, dismissing')
      
      return False

    if not all(elem in data and data[elem] != None for elem in required_keys):
      self._log.warning('One of required keys are not present or its value is equal to None in the data, required keys were %s' % required_keys)

      return False

    return True

  def ask(self, data, raw_msg):
    """Ask required by the skill intent.

    :param data: Data sent by the intent, some keys such as "slot" or "confirm" represents specific cases
    :type data: dict
    :param raw_msg: Raw payload received
    :type raw_msg: str

    """

    if self._is_valid_request(data, ['text', 'slot']):
      self.go(STATE_ASK, slot=data.get('slot'), choices=data.get('choices'), payload=raw_msg)

  def show(self, data, raw_msg):
    """Show simply pass message to the client channel.

    :param data: Data sent by the intent
    :type data: dict
    :param raw_msg: Raw payload received
    :type raw_msg: str

    """

    if self._is_valid_request(data, ['text']):
      self._client.show(raw_msg)

  def _terminate_from_skill(self, data, raw_msg):
    """Called by a skill when it has ended its work.

    :param data: Data sent by the intent
    :type data: dict
    :param raw_msg: Raw payload received
    :type raw_msg: str

    """

    if self._is_valid_request(data):
      self.terminate()

  def terminate(self):
    """Terminates a dialog and returns to the asleep state.
    """

    self.go(STATE_ASLEEP)

  def start(self, config):
    """Starts this agent.

    :param config: Broker configuration
    :type config: BrokerConfig

    """

    self._client.start(config)

  def cleanup(self):
    """Cleanup the agent.
    """

    # TODO Find a way to remove it cleanly
    # self._machine.remove_model(self)

    self._client.destroyed()
    self._client.stop()

  def __str__(self):
    return 'Agent %s - %s' % (self.id, self.interpreter.lang())
