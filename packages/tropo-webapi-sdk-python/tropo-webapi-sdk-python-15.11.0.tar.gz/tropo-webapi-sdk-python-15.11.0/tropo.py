"""
The TropoPython module. This module implements a set of classes and methods for manipulating the Voxeo Tropo WebAPI.

Usage:

----
from tropo import Tropo

tropo = Tropo()
tropo.say("Hello, World")
json = tropo.RenderJson()
----

You can write this JSON back to standard output to get Tropo to perform
the action. For example, on Google Appengine you might write something like:

handler.response.out.write(json)

Much of the time, a you will interact with Tropo by  examining the Result
object and communicating back to Tropo via the Tropo class methods, such
as "say". In some cases, you'll want to build a class object directly such as in :

    choices = tropo.Choices("[5 digits]").obj

    tropo.ask(choices,
              say="Please enter your 5 digit zip code.",
              attempts=3, bargein=True, name="zip", timeout=5, voice="dave")
    ...

NOTE: This module requires python 2.5 or higher.

"""

try:
    import cjson as jsonlib
    jsonlib.dumps = jsonlib.encode
    jsonlib.loads = jsonlib.decode
except ImportError:
    try:
        from django.utils import simplejson as jsonlib
    except ImportError:
        try:
            import simplejson as jsonlib
        except ImportError:
            import json as jsonlib

class TropoAction(object):
    """
    Class representing the base Tropo action.
    Two properties are provided in order to avoid defining the same attributes for every action.
    """
    @property
    def json(self):
        return self._dict

    @property
    def obj(self):
        return {self.action: self._dict}

class Ask(TropoAction):
    """
    Class representing the "ask" Tropo action. Builds an "ask" JSON object.
    Class constructor arg: choices, a Choices object
    Convenience function: Tropo.ask()
    Class constructor options: attempts, bargein, choices, minConfidence, name, recognizer, required, say, timeout, voice

    Request information from the caller and wait for a response.
    (See https://www.tropo.com/docs/webapi/ask)

        { "ask": {
            "attempts": Integer,
            "allowSignals": String or Array,
            "bargein": Boolean,
            "choices": Object, #Required
            "interdigitTimeout": Float, 
            "minConfidence": Integer, 
            "name": String,
            "recognizer": String,
            "required": Boolean,
            "say": Object,
            "sensitivity": Float, 
            "speechCompleteTimeout": Float, 
            "speechIncompleteTimeout": Float,
            "timeout": Float,
            "voice": String,
             
             } }

    """
    action = 'ask'
    options_array = ['attempts', 'allowSignals', 'bargein', 'choices', 'interdigitTimeout', 'minConfidence', 'name', 'recognizer', 'required', 'say', 'sensitivity', 'speechCompleteTimeout', 'speechIncompleteTimeout', 'timeout', 'voice', 'promptLogSecurity', 'asrLogSecurity', 'maskTemplate']

    def __init__(self, choices, **options):
        self._dict = {}
        if (isinstance(choices, basestring)):
            self._dict['choices'] = Choices(choices).json
        else:
#            self._dict['choices'] = choices['choices']
            self._dict['choices'] = choices.json
        for opt in self.options_array:
            if opt in options:
                if ((opt == 'say') and (isinstance(options['say'], basestring))):
                    self._dict['say'] = Say(options['say']).json
                else:
                    self._dict[opt] = options[opt]

class Answer(TropoAction):
    """
    Class representing the "answer" Tropo action. Builds a "answer" JSON object.
    Class constructor arg: headers, a Object
    Class constructor options: headers
    Convenience function: Tropo.answer()

    (See https://www.tropo.com/docswebapi/answer)

    { "answer": {
        "headers": Object  } }
    """
    action = 'answer'
    options_array = []

    def __init__(self, headers, **options):
        self._dict = {'headers': headers}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

                
class Call(TropoAction):
    """
    Class representing the "call" Tropo action. Builds a "call" JSON object.
    Class constructor arg: to, a String
    Class constructor options: answerOnMedia, channel, from, headers, name, network, required, timeout, machineDetection
    Convenience function: Tropo.call()

    (See https://www.tropo.com/docswebapi/call)

    { "call": {
        "to": String or Array,#Required
        "answerOnMedia": Boolean,
        "allowSignals": String or Array
        "channel": string,
        "from": string,
        "headers": Object,
        "name": String,
        "network": String,
        "required": Boolean,
        "timeout": Float.
        "machineDetection: Boolean or Object" 
	"voice": string default "".
        "callbackUrl": string.
        "promptLogSecurity": string default none.
        "label: string"  } }
    """
    action = 'call'
    options_array = ['answerOnMedia', 'allowSignals', 'channel', '_from', 'headers', 'name', 'network', 'required', 'timeout', 'machineDetection', 'voice', 'callbackUrl', 'promptLogSecurity', 'label']

    def __init__(self, to, **options):
        self._dict = {'to': to}
        for opt in self.options_array:
            if opt in options:
                if (opt == "_from"):
                    self._dict['from'] = options[opt]
                else:
                    self._dict[opt] = options[opt]

                

class Choices(TropoAction):
    """
    Class representing choice made by a user. Builds a "choices" JSON object.
    Class constructor options: terminator, mode

    (See https://www.tropo.com/docs/webapi/ask)
    """
    action = 'choices'
    options_array = ['terminator', 'mode']

    def __init__(self, value, **options):
        self._dict = {'value': value}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class TransferOnChoices(TropoAction):
    """
    Class representing choice made by a user. Builds a "choices" JSON object.
    Class constructor options: terminator, mode

    (See https://www.tropo.com/docs/webapi/ask)
    """
    action = 'choices'
    options_array = ['terminator']

    def __init__(self, **options):
        self._dict = {}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class Conference(TropoAction):
    """
    Class representing the "conference" Tropo action. Builds a "conference" JSON object.
    Class constructor arg: id, a String
    Convenience function: Tropo.conference()
    Class constructor options: mute, name, playTones, required, terminator

    (See https://www.tropo.com/docs/webapi/conference)

    { "conference": {
        "id": String,#Required
        "allowSignals": String or Array,
        "interdigitTimeout":Integer,
        "mute": Boolean,
        "name": String,
        "playTones": Boolean,
        "required": Boolean,
        "terminator": String,
        "joinPrompt": Object,
        "promptLogSecurity": string,
        "leavePrompt": Object } }
    """
    action = 'conference'
    options_array = ['allowSignals', 'interdigitTimeout', 'mute', 'name', 'playTones', 'required', 'terminator', 'joinPrompt', 'leavePrompt', 'promptLogSecurity']

    def __init__(self, id, **options):
        self._dict = {'id': id}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class Hangup(TropoAction):
    """
    Class representing the "hangup" Tropo action. Builds a "hangup" JSON object.
    Class constructor arg:
    Class constructor options:
    Convenience function: Tropo.hangup()

    (See https://www.tropo.com/docs/webapi/hangup)

    { "hangup": { } }
    """
    action = 'hangup'

    def __init__(self):
        self._dict = None

class JoinPrompt(TropoAction):
    """
    Class representing join prompts for the conference method. Builds a "joinPrompt" JSON object.
    Class constructor options: value, voice

    (See https://www.tropo.com/docs/webapi/conference)
    """
    action = 'joinPrompt'
    options_array = ['voice']

    def __init__(self, value, **options):
        self._dict = {'value': value}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class LeavePrompt(TropoAction):
    """
    Class representing leave prompts for the conference method. Builds a "leavePrompt" JSON object.
    Class constructor options: value, voice

    (See https://www.tropo.com/docs/webapi/conference)
    """
    action = 'leavePrompt'
    options_array = ['voice']

    def __init__(self, value, **options):
        self._dict = {'value': value}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]


class MachineDetection(TropoAction):
    """
    Class representing machine detection for the call method. Builds a "machineDetection" JSON object.
    Class constructor options: introduction, voice

    (See https://www.tropo.com/docs/webapi/call)
    """
    action = 'machineDetection'
    options_array = ['introduction', 'voice']

    def __init__(self, introduction, **options):
        self._dict = {'introduction': introduction}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]


class Transcription(TropoAction):
    """
    Class representing transcription for the record method. Builds a "transcription" JSON object.
    Class constructor options: id, url, emailFormat, language

    (See https://www.tropo.com/docs/webapi/record)
    """
    action = 'transcription'
    options_array = ['id', 'url', 'emailFormat', 'language']

    def __init__(self, url, **options):
        self._dict = {'url': url}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]


class RecordUrlTuple(TropoAction):
    """
    Class representing recordURL for the record method. Builds a "recordURL" JSON object.
    Class constructor options: url, username, password, method

    (See https://www.tropo.com/docs/webapi/record)
    """
    action = 'url'
    options_array = ['url', 'username', 'password', 'method']

    def __init__(self, url, **options):
        self._dict = {'url': url}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]


class Message(TropoAction):
    """
    Class representing the "message" Tropo action. Builds a "message" JSON object.
    Class constructor arg: say_obj, a Say object
    Class constructor arg: to, a String
    Class constructor options: answerOnMedia, channel, from, name, network, required, timeout, voice
    Convenience function: Tropo.message()

    (See https://www.tropo.com/docs/webapi/message)
    { "message": {
            "say": Object,#Required
            "to": String or Array,#Required
            "answerOnMedia": Boolean,
            "channel": string,
            "from": String,
            "name": String,
            "network": String,
            "required": Boolean,
            "timeout": Float,
	    "promptLogSecurity": String,
            "voice": String } }
    """
    action = 'message'
    options_array = ['answerOnMedia', 'channel', '_from', 'name', 'network', 'required', 'timeout', 'voice', 'promptLogSecurity']

    def __init__(self, say_obj, to, **options):
        self._dict = {'say': say_obj['say'], 'to': to}
        for opt in self.options_array:
            if opt in options:
                if (opt == "_from"):
                    self._dict['from'] = options[opt]
                else:
                    self._dict[opt] = options[opt]


class On(TropoAction):
    """
    Class representing the "on" Tropo action. Builds an "on" JSON object.
    Class constructor arg: event, a String
    Class constructor options:  name,next,required,say
    Convenience function: Tropo.on()

    (See https://www.tropo.com/docs/webapi/on)

    { "on": {
        "event": String,#Required
        "next": String,
        "say": Object
        "post": String } }
    """
    action = 'on'
    options_array = ['next','say', 'post', 'ask']

    def __init__(self, event, **options):
        self._dict = {}
        for opt in self.options_array:
            if opt in options:
                if ((opt == 'say') and (isinstance(options['say'], basestring))):
                    if('voice' in options):
                      self._dict['say'] = Say(options['say'], voice=options['voice']).json
                    else:
                      self._dict['say'] = Say(options['say']).json
             
                elif ((opt == 'ask') and (isinstance(options['ask'], basestring))):
                  if('voice' in options):
                    self._dict['ask'] = Ask(options['ask'], voice=options['voice']).json
                  else:
                    self._dict['ask'] = Ask(options['ask']).json
                else:
                    self._dict[opt] = options[opt]
                    
        self._dict['event'] = event

class Record(TropoAction):
    """
    Class representing the "record" Tropo action. Builds a "record" JSON object.
    Class constructor arg:
    Class constructor options: attempts, bargein, beep, choices, format, maxSilence, maxTime, method, minConfidence, name, password, required, say, timeout, transcription, url, username
    Convenience function: Tropo.record()

    (See https://www.tropo.com/docs/webapi/record)

        { "record": {
            "asyncUpload": Boolean,
            "allowSignals": Boolean,
            "attempts": Integer,
            "bargein": Boolean,
            "beep": Boolean,
            "choices": Object,
            "format": String,
            "interdigitTimeout": Float,
            "maxSilence": Float,
            "maxTime": Float,
            "method": String,
            "name": String,
            "password": String,
            "required": Boolean,
            "say": Object,
            "timeout": Float,
            "transcription": Array or Object,
            "url": String, #Required 
            "username": String,
            "voice": String, 
            "promptLogSecurity": String,
            "sensitivity": String} }
    """
    action = 'record'
    options_array = ['asyncUpload', 'attempts', 'bargein', 'beep', 'choices', 'format', 'maxSilence', 'maxTime', 'method', 'name', 'password', 'required', 'say', 'timeout', 'transcription', 'username', 'allowSignals', 'voice', 'interdigitTimeout', 'promptLogSecurity', 'sensitivity']
    def __init__(self, url, **options):
        self._dict = {'url': url}
        for opt in self.options_array:
            if opt in options:
                if ((opt == 'say') and (isinstance(options['say'], basestring))):
                    self._dict['say'] = Say(options['say']).json
                else:
                    self._dict[opt] = options[opt]

class Redirect(TropoAction):
    """
    Class representing the "redirect" Tropo action. Builds a "redirect" JSON object.
    Class constructor arg: to, a String; name, a String
    Class constructor options:  required
    Convenience function: Tropo.redirect()

    (See https://www.tropo.com/docs/webapi/redirect)

    { "redirect": {
        "to": Object,#Required
        "name": String,,#Required
        "required": Boolean } }
    """
    action = 'redirect'
    options_array = ['required']

    def __init__(self, to, name, **options):
        self._dict = {'to': to, 'name': name}
        #self._dict = {'name': name}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class Reject(TropoAction):
    """
    Class representing the "reject" Tropo action. Builds a "reject" JSON object.
    Class constructor arg:
    Class constructor options:
    Convenience function: Tropo.reject()

    (See https://www.tropo.com/docs/webapi/reject)

    { "reject":  null }
    """
    action = 'reject'

    def __init__(self):
        self._dict = None

class Say(TropoAction):
    """
    Class representing the "say" Tropo action. Builds a "say" JSON object.
    Class constructor arg: message, a String, or a List of Strings
    Class constructor options: attempts, bargein, choices, minConfidence, name, recognizer, required, say, timeout, voice
    Convenience function: Tropo.say()

    (See https://www.tropo.com/docs/webapi/say)

    { "say": {
        "allowSignals": String or Array,
        "voice": String,
        "as": String,
        "name": String,
        "required": Boolean,
        "value": String #Required,
        "promptLogSecurity": String
        } }
    """
    action = 'say'
    # added _as because 'as' is reserved
    options_array = ['_as', 'name', 'required', 'voice', 'allowSignals', 'promptLogSecurity']

    def __init__(self, message, **options):
        dict = {}
        for opt in self.options_array:
            if opt in options:
                if (opt == "_as"):
                    dict['as'] = options[opt]
                else:
                    dict[opt] = options[opt]
        self._list = []
        if (isinstance (message, list)):
            for mess in message:
                new_dict = dict.copy()
                new_dict['value'] = mess
                self._list.append(new_dict)
        else:
            dict['value'] = message
            self._list.append(dict)

    @property
    def json(self):
        return self._list[0] if len(self._list) == 1 else self._list

    @property
    def obj(self):
        return {self.action: self._list[0]} if len(self._list) == 1 else {self.action: self._list}

class StartRecording(TropoAction):
    """
    Class representing the "startRecording" Tropo action. Builds a "startRecording" JSON object.
    Class constructor arg: url, a String
    Class constructor options: format, method, username, password
    Convenience function: Tropo.startRecording()

    (See https://www.tropo.com/docs/webapi/startrecording)

    { "startRecording": {
        "asyncUpload":Boolean,
        "format": String,
        "method": String,
        "url": String,#Required
        "username": String,
        "password": String, 
        "transcriptionID": String
        "transcriptionEmailFormat":String
        "transcriptionLanguage":String
        "transcriptionOutURI": String} }
    """
    action = 'startRecording'
    options_array = ['asyncUpload', 'format', 'method', 'username', 'password', 'transcriptionID', 'transcriptionEmailFormat', 'transcriptionLanguage', 'transcriptionOutURI']
    def __init__(self, url, **options):
        self._dict = {'url': url}
        for opt in self.options_array:
            if opt in options:
                self._dict[opt] = options[opt]

class StopRecording(TropoAction):
   """
    Class representing the "stopRecording" Tropo action. Builds a "stopRecording" JSON object.
    Class constructor arg:
    Class constructor options:
    Convenience function: Tropo.stopRecording()

   (See https://www.tropo.com/docs/webapi/stoprecording)
      { "stopRecording": null }
   """
   action = 'stopRecording'

   def __init__(self):
       self._dict = None

class Transfer(TropoAction):
    """
    Class representing the "transfer" Tropo action. Builds a "transfer" JSON object.
    Class constructor arg: to, a String, or List
    Class constructor options: answerOnMedia, choices, from, name, required, on
    Convenience function: Tropo.transfer()

    (See https://www.tropo.com/docs/webapi/transfer)
    { "transfer": {
        "to": String or Array,#Required
        "allowSignals":String or Array,
        "answerOnMedia": Boolean,
        "choices": Object,
        "from": String,
        "headers": Object,
        "interdigitTimeout":Float,
        "machineDetection": Boolean or Object
        "name": String,
        "playTones":Boolean,
        "required": Boolean,
        "ringRepeat":Integer,
        "on": Onject,
        "timeout": Float,
        "voice": String,
        "callbackUrl": String,
        "promptLogSecurity": String,
        "label": String
     } }
    """
    action = 'transfer'
    options_array = ['answerOnMedia', 'choices', '_from', 'name', 'on', 'required', 'allowSignals', 'headers', 'interdigitTimeout', 'ringRepeat', 'timeout', 'machineDetection', 'playTones', 'voice', 'callbackUrl', 'promptLogSecurity', 'label']
    def __init__(self, to, **options):
      self._dict = {'to': to}
      for opt in self.options_array:
        if opt in options:
          if (opt == '_from'):
            self._dict['from'] = options['_from']
          else:
              self._dict[opt] = options[opt]

class Wait(TropoAction):
      """
      Class representing the "wait" Tropo action. Builds a "wait" JSON object.
      Class constructor arg: milliseconds, an Integer
      Class constructor options: allowSignals
      Convenience function: Tropo.wait()

      (See https://www.tropo.com/docs/webapi/wait)
      { "wait": {
          "milliseconds": Integer,#Required
          "allowSignals": String or Array
      """
      
      action = 'wait'
      options_array = ['allowSignals']

      def __init__(self, milliseconds, **options):
          self._dict = {'milliseconds': milliseconds}
          for opt in self.options_array:
              if opt in options:
                self._dict[opt] = options[opt]

class GeneralLogSecurity(TropoAction):
      """
      Class representing the "generalLogSecurity" Tropo action. Builds a "generalLogSecurity" JSON object.
      Class constructor arg: state, an String
      Convenience function: Tropo.generalLogSecurity()

      (See https://www.tropo.com/docs/webapi/generallogsecurity)
      { "generallogsecurity": {
          "state": String,#Required
      """
#      action = ''
#      state = ''

      def __init__(self, state):
#          self._dict = {'generalLogSecurity': state}
          self.state = state

      @property
      def obj(self):
	  return {'generalLogSecurity':self.state}


class Result(object):
    """
    Returned anytime a request is made to the Tropo Web API.
    Method: getValue
    (See https://www.tropo.com/docs/webapi/result)

        { "result": {
            "actions": Array or Object,
            "calledId": String,
            "callId": String,
            "complete": Boolean,
            "connectedDuration": Integer,
            "duration": Integer,
            "error": String,
            "sequence": Integer,
            "sessionDuration": Integer,
            "sessionId": String,
            "state": String,
            "userType": String} }
    """
    options_array = ['actions','complete','error','sequence', 'sessionDuration', 'sessionId', 'state', 'userType', 'connectedDuration', 'duration', 'calledID', 'callId']
    def __init__(self, result_json):
        result_data = jsonlib.loads(result_json)
        result_dict = result_data['result']

        for opt in self.options_array:
            if result_dict.get(opt, False):
                setattr(self, '_%s' % opt, result_dict[opt])

    def getActions(self):
        actions = self._actions
        return actions
    
    def getActionsCount(self):
        actions = self._actions
        if actions is None:
            count = 0
            return count
        if (type (actions) is list):
            count = len(actions)
        else:
            count = 1
        return count
    
    

    def getValue(self):
        """
        Get the value of the previously POSTed Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            dict = actions[0]
        else:
            dict = actions
        return dict.get('value', 'NoValue')

    def getIndexedValue(self, index):
        """
        Get the value of the indexed Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            dict = actions[index]
        else:
            dict = actions
        return dict.get('value', 'NoValue')

    def getNamedActionValue(self, name):
        """
        Get the value of the named Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            for a in actions:
                if a.get('name', 'NoValue') == name:
                    dict =a
        else:
            dict = actions
        return dict.get('value', 'NoValue')


    def getUserType(self):
      """
      Get the userType of the previously POSTed Tropo action.
      """
      userType = self._userType
      return userType

# # **Tue May 17 07:17:38 2011** -- egilchri

    def getInterpretation(self):
        """
        Get the value of the previously POSTed Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            dict = actions[0]
        else:
            dict = actions
        return dict.get('interpretation', 'NoValue')

    def getIndexdedInterpretation(self, index):
        """
        Get the value of the indexed Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            dict = actions[index]
        else:
            dict = actions
        return dict.get('interpretation', 'NoValue')

    def getNamedActionInterpretation(self, name):
        """
        Get the value of the named Tropo action.
        """
        actions = self._actions

        if (type (actions) is list):
            for a in actions:
                if a.get('name', 'NoValue') == name:
                    dict =a
        else:
            dict = actions
        return dict.get('interpretation', 'NoValue')

# # **Tue May 17 07:17:38 2011** -- egilchri


class Session(object):
    """
    Session is the payload sent as an HTTP POST to your web application when a new session arrives.
    (See https://www.tropo.com/docs/webapi/session)
    
    Because 'from' is a reserved word in Python, the session object's 'from' property is called
    fromaddress in the Python library
    """
    def __init__(self, session_json):
        session_data = jsonlib.loads(session_json)
        session_dict = session_data['session']
        for key in session_dict:
            val = session_dict[key]
            if key == "from":
                setattr(self, "fromaddress", val)
            else:
                setattr(self, key, val)
	setattr(self, 'dict', session_dict)


class Tropo(object):
    """
      This is the top level class for all the Tropo web api actions.
      The methods of this class implement individual Tropo actions.
      Individual actions are each methods on this class.

      Each method takes one or more required arguments, followed by optional
      arguments expressed as key=value pairs.

      The optional arguments for these methods are described here:
      https://www.tropo.com/docs/webapi/
    """
    def  __init__(self):
        self._steps = []

# # **Sun May 15 21:05:01 2011** -- egilchri
    def setVoice(self, voice):
        self.voice = voice

# # end **Sun May 15 21:05:01 2011** -- egilchri

    def ask(self, choices, **options):
        """
        Sends a prompt to the user and optionally waits for a response.
        Arguments: "choices" is a Choices object
        See https://www.tropo.com/docs/webapi/ask
        """
# # **Sun May 15 21:21:29 2011** -- egilchri

        # Settng the voice in this method call has priority.
	# Otherwise, we can pick up the voice from the Tropo object,
	# if it is set there.
        if hasattr (self, 'voice'):
            if (not 'voice' in options):
                options['voice'] = self.voice
        
# # **Sun May 15 21:21:29 2011** -- egilchri

        self._steps.append(Ask(choices, **options).obj)


    def answer (self, headers, **options):
        """
        Places a call or sends an an IM, Twitter, or SMS message. To start a call, use the Session API headers tell Tropo headers launch your code.
        Arguments: headers is a String.
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/answer
        """
        self._steps.append(Answer (headers, **options).obj)

    def call (self, to, **options):
        """
        Places a call or sends an an IM, Twitter, or SMS message. To start a call, use the Session API to tell Tropo to launch your code.
        Arguments: to is a String.
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/call
        """
        self._steps.append(Call (to, **options).obj)

    def conference(self, id, **options):
        """
        This object allows multiple lines in separate sessions to be conferenced together so that the parties on each line can talk to each other simultaneously.
        This is a voice channel only feature.
        Argument: "id" is a String
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/conference
        """
        self._steps.append(Conference(id, **options).obj)

    def hangup(self):
        """
        This method instructs Tropo to "hang-up" or disconnect the session associated with the current session.
        See https://www.tropo.com/docs/webapi/hangup
        """
        self._steps.append(Hangup().obj)

    def message (self, say_obj, to, **options):
        """
        A shortcut method to create a session, say something, and hang up, all in one step. This is particularly useful for sending out a quick SMS or IM.
        Argument: "say_obj" is a Say object
        Argument: "to" is a String
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/message
        """
        if isinstance(say_obj, basestring):
            say = Say(say_obj).obj
        else:
            say = say_obj
        self._steps.append(Message(say, to, **options).obj)

    def on(self, event, **options):
        """
        Adds an event callback so that your application may be notified when a particular event occurs.
        Possible events are: "continue", "error", "incomplete" and "hangup".
        Argument: event is an event
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/on
        """
        if hasattr (self, 'voice'):
          if (not 'voice' in options):
            options['voice'] = self.voice


        self._steps.append(On(event, **options).obj)

    def record(self, **options):
        """
        Plays a prompt (audio file or text to speech) and optionally waits for a response from the caller that is recorded.
         Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/record
        """
        self._steps.append(Record(**options).obj)

    def redirect(self, id, name, **options):
        """
        Forwards an incoming call to another destination / phone number before answering it.
        Argument: id is a String
        Argument: name is a String
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/redirect
        """
        self._steps.append(Redirect(id, name, **options).obj)

    def reject(self):
        """
        Allows Tropo applications to reject incoming sessions before they are answered.
        See https://www.tropo.com/docs/webapi/reject
        """
        self._steps.append(Reject().obj)

    def say(self, message, **options):
        """
        When the current session is a voice channel this key will either play a message or an audio file from a URL.
        In the case of an text channel it will send the text back to the user via i nstant messaging or SMS.
        Argument: message is a string
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/say
        """
        #voice = self.voice
# # **Sun May 15 21:21:29 2011** -- egilchri

        # Settng the voice in this method call has priority.
	# Otherwise, we can pick up the voice from the Tropo object,
	# if it is set there.
        if hasattr (self, 'voice'):
            if (not 'voice' in options):
                options['voice'] = self.voice
# # **Sun May 15 21:21:29 2011** -- egilchri

        self._steps.append(Say(message, **options).obj)

    def startRecording(self, url, **options):
        """
        Allows Tropo applications to begin recording the current session.
        Argument: url is a string
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/startrecording
        """
        self._steps.append(StartRecording(url, **options).obj)

    def stopRecording(self):
        """
        Stops a previously started recording.
        See https://www.tropo.com/docs/webapi/stoprecording
        """
        self._steps.append(StopRecording().obj)

    def transfer(self, to, **options):
        """
        Transfers an already answered call to another destination / phone number.
        Argument: to is a string
        Argument: **options is a set of optional keyword arguments.
        See https://www.tropo.com/docs/webapi/transfer
        """
        self._steps.append(Transfer(to, **options).obj)
        
    def wait(self, milliseconds, **options):
      """
      Allows the thread to sleep for a given amount of time in milliseconds
      Argument: milliseconds is an Integer
      Argument: **options is a set of optional keyword arguments.
      See https://www.tropo.com/docs/webapi/wait
      """
      self._steps.append(Wait(milliseconds, **options).obj)

    def generalLogSecurity(self, state, **options):
      """
      Turn on/off all logging on the Tropo platform
      Argument: state is a String
      See https://www.tropo.com/docs/webapi/generallogsecurity
      """
      self._steps.append(GeneralLogSecurity(state).obj)



    def RenderJson(self, pretty=False):
        """
        Render a Tropo object into a Json string.
        """
        steps = self._steps
        topdict = {}
        topdict['tropo'] = steps
        if pretty:
            try:
                json = jsonlib.dumps(topdict, indent=4, sort_keys=False)
            except TypeError:
                json = jsonlib.dumps(topdict)
        else:
            json = jsonlib.dumps(topdict)
        return json

if __name__ == '__main__':
    print """

 This is the Python web API for http://www.tropo.com/

 To run the test suite, please run:

    cd test
    python test.py


"""


