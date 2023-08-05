# cfgur (configure)

_Load environment vars, and command line arguments in a predictable, standardized way._

This package wraps argparse, loading config from environment variables as well as the usual command line arguments / flags. A config object is passed into init(), configuring argparse. The environment variables are implied from the arg name (stripped, uppercased and underscored), but can be manually defined using the key "env_var" in args config. Once args are loaded, config is accessable using a getter function "get()". To get all variables, use get_all(). The args object keys are passed to parser.add_argument().

### Installation

The cfgur package is available on PyPI which means installation should be as simple as:

```
$ pip install cfgur
```

### Usage

```python
import cfgur

# Initialize once, pass in config
cfgur.init({
    "parser": {
        "description": "this is the service description"
    },
    "args": [
        {
            "name": "-test",
            "help": "Description of test",
            "default": "test-db",
            "required": False,
        }, {
            "name": "--foo",
            "required": True,
        }, {
            "name": "-f",
            "required": False,
            "env_var": "LETTER_F",
        }
    ]
})

# Return value of config. If missing, return None
my_config = cfgur.get_all()

# Return value of config. If missing, return None
test_val = cfgur.get("test")

# Return value of config. If missing, raise KeyError
test_val = cfgur.get("test", True)
```


### License
This software is licensed under the MIT license.

© 2018 Bray Almini.
