# userpass

-----

**Table of Contents**

* [Installation](#installation)
* [Usage](#usage)
* [License](#license)

## Installation

userpass is distributed on [PyPI](https://pypi.org) as a universal
wheel and is available on Linux/macOS and Windows and supports
Python 2.7/3.5+ and PyPy.

```bash
$ pip install userpass
```

## Usage

```python
from userpass import User
# Userinfo file will be stored in `~/.user` file.
user = User('.user')
username = user.username
password = user.password
user.del_user()
```

## License

userpass is distributed under the terms of the
[MIT License](https://choosealicense.com/licenses/mit).
