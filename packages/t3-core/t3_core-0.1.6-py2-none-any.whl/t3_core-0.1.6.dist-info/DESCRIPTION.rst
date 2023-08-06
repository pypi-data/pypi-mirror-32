# T3 Python Core Library

## Using Forge

### Initialization

```bash
forge setup
```

### cloaked_forge

### Install

```bash
$ pip install -e .
```

### Requirements and Usage

See docs [cloaked_forge](./docs/cloaked_forge/README.md)

## Install

### Setup Virtualenv (optional)
```sh
python -m venv .venv
source .venv/bin/activate
```

### Install Local Copy
```sh
pip install -e .
```

### Install from Gitlab
```sh
pip install git+ssh://git@gitlab.t-3.com/sunoco/sunoco-service-locations.git
```

### Run Migrations
```sh
manage migrate
# OR
cd app
./manage.py migrate
```

## Testing & Linting
### Test
```sh
manage test locations
# OR
cd app
./manage.py test
```

### Coverage Report
```sh
coverage run -m manage test locations && coverage report
# OR
coverage run ./app/manage.py test locations && coverage report
# OR
cd app
coverage run manage.py && coverage report
```

### Lint
```sh
pylama
```


