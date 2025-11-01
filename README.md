# Hash - High-level Cryptographic Hash Wrappers for Go [![Go Reference](https://pkg.go.dev/badge/github.com/bytemare/hash.svg)](https://pkg.go.dev/github.com/bytemare/hash)

This package provides High-level Go wrappers for standard hash functions, unifying a set of common cryptographic hash operations behind a single API, allowing callers to switch between algorithms with minimal code changes while relying on well-maintained and secure backends.
It unifies Merkle–Damgård and extendable-output functions behind a single API, providing HMAC and HKDF helpers along with consistent metadata (block sizes, security strength, availability). Hashers are immutable handles that produce fresh state,
preventing cross-call contamination and enforcing sane defaults for sizes.

## Import

```
import "github.com/bytemare/hash"
```

## Quick Start

- Quickly hash data:

```go
message := []byte("message")

output := hash.SHA256.Hash(message)
```

- HMAC

```go
message := []byte("message")
key := []byte("key")

hmac := hash.SHA256.GetHashFunction().Hmac(message, key)
```

- If you have an identifier from the crypto package, it's easy:
```go
id := crypto.SHA256
message := []byte("message")

output := hash.FromCrypto(id).Hash(message)
```

See [`examples_test.go`](examples_test.go) for additional HMAC, HKDF extract/expand, and metadata usage patterns.

## Versioning and Compatibility

The module follows [Semantic Versioning](https://semver.org/) for its public API. New hashes or helper APIs as well as breaking changes ship in minor releases.

## License

Released under the [MIT License](LICENSE). See `LICENSE` for full terms.
