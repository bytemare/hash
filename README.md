# github.com/bytemare/hash

[![CI](https://github.com/bytemare/hash/actions/workflows/wf-tests.yaml/badge.svg)](https://github.com/bytemare/hash/actions/workflows/wf-tests.yaml)
[![Analysis](https://github.com/bytemare/hash/actions/workflows/wf-analysis.yaml/badge.svg)](https://github.com/bytemare/hash/actions/workflows/wf-analysis.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bytemare/hash.svg)](https://pkg.go.dev/github.com/bytemare/hash)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/bytemare/hash/badge)](https://securityscorecards.dev/viewer/?uri=github.com/bytemare/hash)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

High-level Go wrappers for standard hash functions, unifying set of common cryptographic hash operations behind a single API, allowing callers to switch between algorithms with minimal code changes while relying on well-maintained and secure backends.
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

## API

- `hash.Hash` enumerates supported algorithms and exposes metadata helpers (`Hash.Available`, `Hash.Type`, etc.) that route calls to registered constructors (`hash.go:24`, `hash.go:85`).
- `Fixed` wraps Merkle–Damgård style hashes and adds HMAC/HKDF utilities that stay within Go's approved primitives (`fixed.go:64`, `fixed.go:135`).
- `ExtendableHash` provides SHAKE/BLAKE2X streaming support with configurable output sizes (`extensible.go:71`, `extensible.go:127`).

When no lengths are provided, the standard lengths are used.

**Panics occur in the following scenarios, to indicate the caller doesn't adhere to secure usage patterns:**
- requested output lengths for XOFs it too short (avoid truncation of security guarantees)
- Supplying HMAC keys longer than the output length (clip or hash your keys first if they overflow)

These panics must be considered as bugs from the caller, wrap them if this is unacceptable.

## Addressing Common Weaknesses

| Weakness                                                      | Implemented countermeasures                                                                                                      |
|---------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------|
| Improper key length enforcement (CWE-325)                     | HMAC helper panics on oversized keys and tests ensure panic occurs (`fixed.go:135`, `tests/fixed_test.go:33`).                   |
| Insufficient output length for cryptographic hashes (CWE-326) | Extendable hash `Read` verifies requested size to prevent truncated outputs (`extensible.go:95`).                                |
| Exceeding HKDF block limits (CWE-331)                         | HKDF methods propagate upstream errors and tests enforce failure on excessive length (`fixed.go:147`, `tests/fixed_test.go:64`). |
| Unregistered algorithm usage (CWE-693)                        | Registry gate ensures only compiled algorithms are exposed, preventing accidental downgrade to weaker hashes (`hash.go:60`).     |

## Versioning and Compatibility

The module follows [Semantic Versioning](https://semver.org/) for its public API. New hashes or helper APIs as well as breaking changes ship in minor releases.

## License

Released under the [MIT License](LICENSE). See `LICENSE` for full terms.
