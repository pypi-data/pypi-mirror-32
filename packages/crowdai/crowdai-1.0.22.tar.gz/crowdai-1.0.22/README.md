# crowdai-client-py
Python client to interact with CrowdAI Grading Server.

# Installation Instruction
```
pip install crowdai
```

```
import sys
import crowdai

api_key = "YOUR-API-KEY-HERE"
challenge = crowdai.Challenge("CHALLENGE_ID", api_key)

data = ...

response = challenge.submit(data)
print response
challenge.disconnect()
```

## Implemented Challenges and associated functions

* OpenSNPChallenge2017
  - `submit` :
    Submit the list of heights for the test set
* CriteoAdPlacementNIPS2017
  - `submit` :
    Submits the path to a prediction file for the [Criteo Ad Placement challenge](https://www.crowdai.org/challenges/nips-17-workshop-criteo-ad-placement-challenge)
* Learning2RunChallengeNIPS2017
  - `submit`:
    Submits a docker tar dump for the [NIPS 2017: Learning to Run Challenge]
* AIGeneratedMusicChallenge
  - `submit`:
    Submits a midi file of length 3600 seconds (at 120 bpm) to the grading interface
* KITEnergyChallenge
  - `submit`:
    Submits a the forecasts for the problem definition in the KITEnergyChallenge
* IEEEInvestmentRankingChallenge
  - `submit`:
    Submits a the predictions for the problem definition in the IEEEInvestmentRankingChallenge
    Args :
      * `filepath` : Filepath of the submission file
      * `round` : 1 **or** 2 depedening on the   
* crowdAIMappingChallenge
  - `submit`:
    Submits a predicted `annotations.json` file for the problem definition in [crowdAI Mapping Challenge](https://www.crowdai.org/challenges/mapping-challenge)
* crowdAIGenericChallenge
  - `submit`:
    Submits a prediction file for any arbitrary challenge on crowdAI.

# Author
S.P. Mohanty <sharada.mohanty@epfl.ch>
