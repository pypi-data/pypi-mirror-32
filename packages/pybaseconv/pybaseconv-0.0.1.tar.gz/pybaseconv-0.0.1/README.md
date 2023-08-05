# pybaseconv

[![Build Status](https://travis-ci.com/El-Sam/pybaseconv.svg?branch=master)](https://travis-ci.com/El-Sam/pybaseconv)

**pybaseconv** is a library that allows you to convert any large number (no fractions) from any base to any other base, thus freedom to choose source and/or destination bases.

There are 5 predefined bases that are supported by **pybaseconv**, which are:

* `DEC` : the decimal base  `"0123456789"`.
* `BIN` : the binary base `"01"`.
* `OCT` : the octal base `"01234567"`.
* `HEX` : the hexadecimal base `"0123456789abcdef"`.
* `FLICKER_BASE_58` & `BITCOIN_BASE_58`: two variations of the **Base58** encoding, more info [here](https://en.wikipedia.org/wiki/Base58).

## Module usage

```python
from pybaseconv import Converter, BASE

dec2hex_converter = Converter(BASE.DEC, BASE.HEX)
print(dec2hex_converter.convert('738653'))  # returns b455d

dec2bitcoin_converter = Converter(BASE.DEC, BASE.BITCOIN_BASE_58)
print(dec2bitcoin_converter.convert('292251'))  # returns 2Vsp

dec2custom_base_converter = Converter(BASE.DEC, '*&@#$')
print(dec2custom_base_converter.convert('539'))  # returns $&@$

bin2dec_converter = Converter(BASE.BIN, BASE.DEC)
print(bin2dec_converter.convert('11112'))  # raises exception because 11112 is not a binary number

```

*PS*: when using custom bases, make sure that the digits are sorted from smallest to the largest, in order to get the appropriate conversion.
