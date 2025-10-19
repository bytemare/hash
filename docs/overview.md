# Project Overview

`github.com/bytemare/hash` provides high-level abstractions for working with multiple hashing functions in Go. It unifies a set of common cryptographic hash operations behind a single API, allowing callers to switch between algorithms with minimal code changes while relying on well-maintained and secure backends.

## Goals

- Offer a single package that can switch between supported groups with minimal call-site changes.
- Provide consistent error handling and encoding semantics across backends.
- Keep dependencies limited to built-in cryptographic libraries.

## Module Layout

- `hash.go` contains the public `Hash` enumeration, registration tables, and helper metadata methods exposed by the
  package (`Hash.Available`, `Hash.New`, `Hash.SecurityLevel`).
- `fixed.go` implements the `Fixed` hasher wrapper with HMAC and HKDF helpers for Merkle–Damgård algorithms, built on
  Go's `crypto` package primitives.
- `extensible.go` defines the `ExtendableHash` implementation for SHAKE and BLAKE2X extendable-output functions backed
  by `crypto/sha3` and `golang.org/x/crypto`.
- `tests/` exercises both variants (functional, fuzz, HKDF edge cases) ensuring metadata and panic pathways behave as
  documented, while `examples_test.go` provides runnable snippets for the README and GoDoc.

## Supported Go Versions

- `go.mod` targets the latest available Go version, which is the primary development toolchain.
- GitHub Actions run tests against the three latest Go version (`.github/workflows/wf-tests.yaml`).
- Older toolchains may compile, but they are outside the support window and receive no compatibility guarantees.

## Compatibility Policy

- Semantic Versioning governs API stability. Breaking changes require a minor version bump.
- Upstream breaking changes in backends are introduced only after evaluation and version pinning.
