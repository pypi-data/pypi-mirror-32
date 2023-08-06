# Behavioral Signals CLI

Command Line Interface for Behavioral Signals' Emotion and Behavior Recognition in the Cloud


* Free software: MIT license

## Install
```
pip install behavioral_signals_cli
```

## Getting Started

* First request your account id and token for the Behavioral Signals Web API by sending an email to nassos@behavioralsignals.com
* Export the following as environmental variables (see below for ore elaborate configuration):
```
   export BEST_API_ID=your_id_on_service_api
   export BEST_API_TOKEN=your_token_for_service_api
```

* Run the CLI to submit your audio files (use bsi-cli for less typing):
```
   behaviorals_signals_cli send_audio [csv_file] [pids_log]
```
   The .csv file must have the following form (order matters):
path/to/file, number of channels, call direction, agentId, agentTeam, campaign Id, calltype, calltime, timezone, ANI. The [pids_log] file
is an empty file where the process ids of the created jobs will be written.

* Run the CLI to get the emotion/behavior recognition, diarization and other results:
```
   behaviorals_signals_cli get_results [pids_log] [results_dir]
```
   The results will be written as .json files inside [results_dir] (polling may be performed if results
   are not readily available).
   
* Run the CLI to get ASR results:
```
   behaviorals_signals_cli get_results_asr [pids_log] [results_dir]
```
   The results will be written as "[filename]_[pid]_words.json" files inside [results_dir] (polling may be performed if results
   are not readily available).

you may use bsi-cli instead of behaviorals_signals_cli 
for less typing :-)

Type:
```
   behavioral_signals_cli --help 
```
for more info.
   

Features
--------
The CLI allows you to easily:

- Submit multiple audio files to API,
- Get behavior and emotion recognition results
- Get speech recognition results

* TODO

Configuration
--------
The CLI allows you to use a flexible configuration, since it accepts
with increasing priority:

-  internal defaults (specified in the source code of the cli)
-  environment variables (for the security tokens)
-  external configuration file with sections (see demo for examples). If not specified, it will look for ./bsi-cli.conf and ~/.bsi-cli.conf)
-  command line opions (see bsi-cli -h)

the subcommand: "bsi-cli config" displays the current configuration values.
For example:
```
*** bsi-cli configuration
           apiid : yyyy
        apitoken : xxxxx
          apiurl : http://api.bsis.me:8080
  configLocation : bsi-cli.conf
      configfile : None
             log : WARNING
            stag : default
***
```

Main configuration variables
---------------------

- BSI_API_ID (apiid, or --apiid)      : application id -- provided by BSI
- BSI_API_TOKEN (apitoken, --apitoken) : application token -- provided by BSI
- apiurl, --apiurl: usually http://api.bsis.me:8080: the address of the caller service

BSI_* are environment variables
api* are configuration file parameters
--api* are command line parameters


Credits
---------

This package was created with Cookiecutter_ and the `audreyr/cookiecutter-pypackage`_ project template.

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`audreyr/cookiecutter-pypackage`: https://github.com/audreyr/cookiecutter-pypackage

