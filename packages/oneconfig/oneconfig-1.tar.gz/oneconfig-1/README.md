
# OneCofnig

A config lib for python.

## Installation

```console
pip install oneconfig
```

## Usage
For example, there is a config file named `appsettings.Development.json`

```python

from oneconfig.cores import Configuration

class MyConfig(Configuration):
    port = 8080
    domain = "localhost"
    
myconfig = MyConfig()
myconfig.add_file_by_prefix("appsettings") # it will loads all config files startswith `appsettings`
print(MyConfig.domain) # localhost 

myconfig.mode = 'Production'
myconfig.add_file_by_prefix("appsettings", follow_mode=True)
print(MyConfig.domain) # empty
```

## LICENSE

The MIT License (MIT)

Copyright (c) 2018 jeremaihloo
