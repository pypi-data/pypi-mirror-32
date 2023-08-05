# yamole [![Build Status](https://travis-ci.org/YagoGG/yamole.svg?branch=master)](https://travis-ci.org/YagoGG/yamole)

Dig through the JSON references inside a YAML file, the kind of situation
you may run into when parsing [OpenAPI](https://www.openapis.org/) files.

The result is a single, big YAML file with all the references resolved (i.e.
with their contents replaced in the corresponding places).

## Installation

yamole is available as a PyPI module, so you can install it using `pip`:

    $ pip install yamole

## Usage

Using yamole is pretty straightforward. The parser is available through the
`YamoleParser` class:

```python
parser = YamoleParser('input_file.yaml')

output_str = parser.dumps()
```

## Testing

To test that yamole works properly, you can run:

    $ pip install -r requirements.txt
    $ python tests/test.py

This will run the parser against a specific test case that makes use of all of
yamole's features, and will compare the result with a fixture
(`tests/expected.yaml`).

---

(c) 2018 Yago González. All rights reserved
