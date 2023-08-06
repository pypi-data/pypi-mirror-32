# ATT&CK Python Client

A Python module to access up to date ATT&CK content available in STIX via public TAXII server. This project leverages the python classes and functions of the [cti-python-stix2](https://github.com/oasis-open/cti-python-stix2) and [cti-taxii-client](https://github.com/oasis-open/cti-taxii-client) libraries developed by MITRE.

# Goals

* Provide an easy way to access and interact with up to date ATT&CK content available in STIX via public TAXII server
* Allow security analysts to quickly 
* Allow the integration of ATT&Ck with other platforms to host up to date information from the framework.
* Help security analysts during the transition from the ATT&CK MediaWiki API to the STIX/TAXII 2.0 API
* Learn STIX2 and TAXII Client Python libraries

# Current Status: Beta

The project is currently in a beta stage, which means that the code and the functionality is changing, but the current main functions are stabilising. I would love to get your feedback to make it a better project.

# Resources

* [MITRE CTI](https://github.com/mitre/cti)
* [OASIS CTI TAXII Client](https://github.com/oasis-open/cti-taxii-client)
* [OASIS CTI Python STIX2](https://github.com/oasis-open/cti-python-stix2)
* [MITRE ATT&CK Framework](https://attack.mitre.org/wiki/Main_Page)
* [ATT&CK MediaWiki API](https://attack.mitre.org/wiki/Using_the_API)
* [Invoke-ATTACKAPI](https://github.com/Cyb3rWard0g/Invoke-ATTACKAPI)
* [Mitre-Attack-API](https://github.com/annamcabee/Mitre-Attack-API)

# Getting Started
## Install Requirements

```
pip install -r requirements.txt
```

# Author

* Roberto Rodriguez [@Cyb3rWard0g](https://twitter.com/Cyb3rWard0g)

# Contributors

* Jose Luis Rodriguez [@Cyb3rPandaH](https://twitter.com/Cyb3rPandaH)

# Contributing

# To-Do

* [ ] Revokation logic to update Groups Objects
