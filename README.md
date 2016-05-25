# Serializer

This module provides function to go from string to [opaque] string or vice & versa

## Installation

go get github.com/GerardSoleCa/go-serializer

## Usage

* serializer.SecureStringify(str, encryptKey, validationKey string): Returns a string representing the given original string. It is signed and encrypted using the given keys.
* serializer.SecureParse(str, encryptKey, validationKey string): Returns the original string decrypting and validating the secure one

The cypher used is aes256, the crypted data is in hex. The signing process uses HMAC with SHA1.

## Test

```bash
$ go test
```

## Credits

Extracted from [node-serializer](https://github.com/AF83/node-serializer).

Original author: Pierre Ruyssen

## License

MIT

