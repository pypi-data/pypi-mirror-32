# Veriteos Python Library

Official Python client library for Veriteos API.

## Documentation

See the [https://docs.veriteos.com/?python](Python API docs).

## Installation

Install the `veriteos` Python package using `pip`:

```
$ pip install --upgrade veriteos
```

### Requirements

- Python 2.6+ or Python 3.3+

## Usage

```python
import veriteos

client = veriteos.Client('VERITEOS_API_KEY', 'TRANSACTION_FAMILY', '1.0')

# Send a transaction
data = 'hello veriteos'
address = 'my instance'
client.send(address, data)

# Read state
client.read(address)
```

## Troubleshooting

If you notice any problems, please drop us an email at [support@veriteos.com](mailto:support@veriteos.com).
