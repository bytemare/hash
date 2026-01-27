# Hash - High-level Cryptographic Hash Wrappers for Go [![Go Reference](https://pkg.go.dev/badge/github.com/bytemare/hash.svg)](https://pkg.go.dev/github.com/bytemare/hash)

This package provides High-level Go wrappers for standard hash functions, unifying a set of common cryptographic hash operations behind a single API, allowing callers to switch between algorithms with minimal code changes while relying on well-maintained and secure backends.
It unifies Merkle–Damgård and extendable-output functions behind a single API, providing HMAC and HKDF helpers along with consistent metadata (block sizes, security strength, availability). Hashers are immutable handles that produce fresh state,
preventing cross-call contamination and enforcing sane defaults for sizes.

## Import

```go
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

## Versioning

[SemVer](https://semver.org) is used for versioning. For the versions
available, see the [tags on the repository](https://github.com/bytemare/opaque/tags).

## Release Integrity (SLSA Level 3)
Releases are built with the reusable [bytemare/slsa](https://github.com/bytemare/slsa) workflow and ship the evidence required for SLSA Level 3 compliance:

- 📦 Artifacts are uploaded to the release page, and include the deterministic source archive plus subjects.sha256, signed SBOM (sbom.cdx.json), GitHub provenance (*.intoto.jsonl), a reproducibility report (verification.json), and a signed Verification Summary Attestation (verification-summary.attestation.json[.bundle]).
- ✍️ All artifacts are signed using [Sigstore](https://sigstore.dev) with transparency via [Rekor](https://rekor.sigstore.dev).
- ✅ Verification (or see the latest docs at [bytemare/slsa](https://github.com/bytemare/slsa)):
```shell
curl -sSL https://raw.githubusercontent.com/bytemare/slsa/main/verify-release.sh -o verify-release.sh
chmod +x verify-release.sh
./verify-release.sh --repo <owner>/<repo> --tag <tag> --mode full --signer-repo bytemare/slsa
```
Run again with `--mode reproduce` to build in a container, or `--mode vsa` to validate just the verification summary.

## Contributing

Please read [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details on the code
of conduct, and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE)
file for details.
