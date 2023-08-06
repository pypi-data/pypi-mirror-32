Feature: Bubble storage type:dataset sqlite
Scenario: load mysrclient.py pull and store in dataset default type sqlite
  Given a file named ".bubble" with:
            """
            bubble=0.7.2
            """
    And a file named "./config/config.yaml" with:
            """
            ---
            CFG:
                BUBBLE:
                    STORAGE_TYPE: dataset
                    STORAGE_DATASET_ARGS:
                        DS_TYPE: sqlite
                        DS_BUBBLE_TAG: testdefault
                DEV:
                    SOURCE:    #pull
                        CLIENT: ./mysrcclient.py
            ...
            """
    And a directory named "./remember/archive"
    And a file named "./mysrcclient.py" with:
            """
            from bubble3 import Bubble
            class BubbleClient(Bubble):
                def __init__(self,cfg={}):
                    self.CFG=cfg
                def pull(self, amount=100, index=0):
                    self.say('BC: %d,%d'%(amount,index))
                    for i in range(amount):
                        it={'keyA':'A_'+str(i),
                            'keyB':'B_'+str(i),
                            'keyC':['c',66,{'keyDinList':'D_'+str(i)}]}
                        self.say('BC:inloop:%d %d'%(amount,index),stuff=it,verbosity=100)
                        yield it
                    #return ret
            """
    When I run "bubble3 pull --amount 10"
    Then the command output should contain "saved result in dataset[step:pulled][stage:DEV]"
    And the command returncode is "0"
    When I run "bubble3 export -r pulled -kvp -f tab -c keyA,keyB,keyC.1,keyC.2,keyC.3.keyDinList"
    Then the command returncode is "0"
    And the command output should contain
      """
      BUBBLE_IDX|keyA|keyB|keyC.1|keyC.2|keyC.3.keyDinList
      ----------|----|----|------|------|-----------------
      0         |A_0 |B_0 |c     |66    |D_0              
      1         |A_1 |B_1 |c     |66    |D_1              
      2         |A_2 |B_2 |c     |66    |D_2              
      3         |A_3 |B_3 |c     |66    |D_3              
      4         |A_4 |B_4 |c     |66    |D_4              
      5         |A_5 |B_5 |c     |66    |D_5              
      6         |A_6 |B_6 |c     |66    |D_6              
      7         |A_7 |B_7 |c     |66    |D_7              
      8         |A_8 |B_8 |c     |66    |D_8              
      9         |A_9 |B_9 |c     |66    |D_9
      """
