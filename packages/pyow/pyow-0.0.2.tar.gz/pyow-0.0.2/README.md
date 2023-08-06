> Function argument validation for humans for Python 3!


## Highlights

- Straight port from the amazing library [`ow`](https://github.com/sindresorhus/ow/) from Sindre Sorhus
- Expressive chainable API
- Lots of built-in validations
- Supports custom validations
- Written in Python3 with Type hinting


## Notes
Since this is a straight up port from the JavaScript library not all features are available.
Partly since this is a port and I haven't caught up and also since Python doesn't support all usecases
as JavaScript does.

## Install

```
$ pip3 install pyow
```


## Usage

```python
from pyow import pyow;

def unicorn(input):
	pyow(input, pyow.string.min_length(5));

	# …

unicorn(3)
>>> ArgumentError: ('Expected argument to be of type `str` but received type `int`')

unicorn('yo');
>>> ArgumentError: ('Expected string to have a minimum length of `3`, got `yo`'
```

## API

### ow(value, predicate)

Test if `value` matches the provided `predicate`.  Throws an `ArgumentError` if the test fails.

### ow.{type}

All the below types return a predicate. Every predicate has some extra operators that you can use to test the value even more fine-grained.

#### Primitives

- [`string`]()
- [`number`]()
- [`boolean`]()

#### Built-in types

- [`list`]()
- [`error`]()
- [`dict`]()
- [`set`]()

## Maintainers

- [Henrik Andersson](https://github.com/limelights)


## Related

- [@sindresorhus/ow](https://github.com/sindresorhus/ow) - Function argument validation for humans


## License

MIT
