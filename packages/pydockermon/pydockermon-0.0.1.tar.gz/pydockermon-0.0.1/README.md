# A python module to interact with ha-dockermon.

#### Notes
This has been tested with python 3.6  
This module require a Cloudflare username and API key.  
This module uses these external libararies:
- requests

#### Install
```bash
pip install pydockermon
```

#### Usage:
```python
from pydockermon import Dockermon

dm = Dockermon()
host = '192.168.1.3'
port = 8126 #Optional
exclude = ['MyFirstContainer', 'MySecondContainer'] #Optional

containers = dm.listContainers(host)
print(containers) #Returns a list of all containers with state, or False if it fails.
containerstate = dm.getContainerState('MyContainer', host, port)
print(containerstate) #Returns the container state, or False if it fails.
containerstats = dm.getContainerStats('MyContainer', host, port)
print(containerstats) #Returns stats about the container, or False if it fails.
stopcontainer = dm.stopContainer('MyContainer', host, port)
print(stopcontainer) #Returns True/false.
startcontariner = dm.startContainer('MyContainer', host, port)
print(startcontariner) #Returns True/False.
```