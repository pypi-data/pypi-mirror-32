## Dvina

Active Learning library in Python. Dvina is a library of components for Active Learning 
and Assisted Ground Truth Retrieval.

## Prerequisites
The library depends on some basic components for Machine Learning, such as
```
pip install future sklearn numpy scipy jsonpickle
```
and works for both Python 2.7 and 3.6.
## Installation

```
pip install dvina
```

## Documentation
The detailed documentation of the library and examples
can be found here: [https://awslabs.github.io/dvina/](https://awslabs.github.io/dvina/).

To rebuild it, use
```
sphinx-build -b html doc docs
```

## Running the tests
Go to the root `dvina` of the library, and run
```
pytest
```

## Contributing
Contributions are welcome.

## Authors
* Fedor Zhdanov - [fedorzh](https://github.com/fedorzh)

Ideas and contributions came from the Amazon Core Machine Learning team.

## License

This library is licensed under the Apache 2.0 License. 
