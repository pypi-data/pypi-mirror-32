Feature: Bubble custom local src client with query method
Scenario: load mysrclient.py in the Bubble directory, with query method
  Given a file named ".bubble" with:
            """
            bubble=0.1.1
            """
    And a file named "./config/config.yaml" with:
            """
            ---
            CFG:
                BUBBLE:
                    DEBUG: True
                    VERBOSE: True
                    STORAGE_TYPE: json
                DEV:
                    SOURCE:    #pull
                        CLIENT: ./mysrcclient.py
                        URL: http://localhost:8001
                    TRANSFORM:
                        RULES: config/rules.bubble
                    TARGET:    #push
                        CLIENT: dummy
                        URL: http://localhost:8002
            ...
            """

    And a directory named "./remember/archive"
    And a file named "./mysrcclient.py" with:
            """
            from bubble3 import Bubble
            class BubbleClient(Bubble):
                def __init__(self,cfg={}):
                    self.CFG=cfg
                def pull(self, *args, **kwargs):
                    self.say('pull:args:',stuff=args)
                    self.say('pull:kwargs:',stuff=args)
                    return [{"keyA":"A1"},{"keyA":"A2"}]
                def query(self, id):
                    self.say('query:id:'+id)
                    return {"queryID":"ID_"+id}
            """

  When I run "bubble3 pull"
    Then the command output should contain "pulled [2] objects"
    Then the command output should contain "remember/pulled_DEV.json"
    And the command returncode is "0"
  When I run "bubble3 export --stepresult pulled --select 'keyA' --showkeys -f tab"
    Then the command returncode is "0"
    And the command output should contain:
            """
            keyA
            ----
            A1
            A2
            """
  When I run "bubble3 pull -q some_id"
    Then the command output should contain "pulled [1] objects"
    Then the command output should contain "remember/pulled_DEV.json"
    And the command returncode is "0"
  When I run "bubble3 export --stepresult pulled --select 'queryID' -kv -f tab"
    Then the command returncode is "0"
    And the command output should contain:
            """
            queryID
            ----------
            ID_some_id
            """
