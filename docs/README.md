# Documentations for gofiber-template

This folder is dedicated to the documentation . this is the source code for building documentation and you can make the newest version of docs for yourself locally or contribute to this project , which will be really appreciated ;) .

Let's get our hands on !

> [!NOTE]
> This documentation for building docs is designed for Debian distribution of Linux operating system , so If you are on windows you might search for some command-lines on Google .

## Prerequisites

- python 3.12 >=
- python3-venv (for linux)
- sphinx-build >= 5.3.0
- sphinx-autobuild >= 2024.10.03

## Initialization

Open a terminal and navigate to this folder :

```bash
cd docs/
```

Then make a python virtual environment :

```bash
python3.12 -m venv venv
```

Then activate the `environment` :

```bash
source venv/bin/activate
```

Then upgrade the pip :

```bash
(venv) user@hostname:~/path/to/docs
$ pip install --upgrade pip
```

Then install `sphinx-autobuild` :

```bash
(venv) user@hostname:~/path/to/docs
pip install sphinx-autobuild
```

Install `furo` :

```bash
(venv) user@hostname:~/path/to/docs
pip install furo
```

## Development Process

Now that everything is ready you can start development process by running the command :

```bash
(venv) user@hostname:~/path/to/docs
sphinx-autobuild source/ build/
```

Or you can just build the documentation using `make` :

```bash
make html
```

The built documentation is in the `build` folder .
