# Design and Security Overview

`github.com/bytemare/hash` provides high-level abstractions for working with multiple hashing functions in Go. It unifies a set of common cryptographic hash operations behind a single API, allowing callers to switch between algorithms with minimal code changes while relying on well-maintained and secure backends.

## Goals

- Offer a single package that can switch between supported groups with minimal call-site changes.
- Provide consistent error handling and encoding semantics across backends.
- Keep dependencies limited to built-in cryptographic libraries.

## API

- `hash.Hash` enumerates supported algorithms and exposes metadata helpers (`Hash.Available`, `Hash.Type`, etc.) that route calls to registered constructors (`hash.go:24`, `hash.go:85`).
- `Fixed` wraps Merkle–Damgård style hashes and adds HMAC/HKDF utilities that stay within Go's approved primitives (`fixed.go:64`, `fixed.go:135`).
- `ExtendableHash` provides SHAKE/BLAKE2X streaming support with configurable output sizes (`extensible.go:71`, `extensible.go:127`).

When no lengths are provided, the standard lengths are used.

**Panics occur in the following scenarios, to indicate the caller doesn't adhere to secure usage patterns:**
- requested output lengths for XOFs it too short (avoid truncation of security guarantees)
- Supplying HMAC keys longer than the output length (clip or hash your keys first if they overflow)

These panics must be considered as bugs from the caller, wrap them if this is unacceptable.

## Module Layout

- `hash.go` contains the public `Hash` enumeration, registration tables, and helper metadata methods exposed by the
  package (`Hash.Available`, `Hash.New`, `Hash.SecurityLevel`).
- `fixed.go` implements the `Fixed` hasher wrapper with HMAC and HKDF helpers for Merkle–Damgård algorithms, built on
  Go's `crypto` package primitives.
- `extensible.go` defines the `ExtendableHash` implementation for SHAKE and BLAKE2X extendable-output functions backed
  by `crypto/sha3` and `golang.org/x/crypto`.
- `tests/` exercises both variants (functional, fuzz, HKDF edge cases) ensuring metadata and panic pathways behave as
  documented, while `examples_test.go` provides runnable snippets for the README and GoDoc.

## Compatibility Policy

- Semantic Versioning governs API stability. Breaking changes require a minor version bump.
- Upstream breaking changes in backends are introduced only after evaluation and version pinning.

## Addressing Common Weaknesses

| Weakness                                                      | Implemented countermeasures                                                                                   |
|---------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| Improper key length enforcement (CWE-325)                     | HMAC helper panics on oversized keys and tests ensure panic occurs.                                           |
| Insufficient output length for cryptographic hashes (CWE-326) | Extendable hash `Read` verifies requested size to prevent truncated outputs.                                  |
| Exceeding HKDF block limits (CWE-331)                         | HKDF methods propagate upstream errors and tests enforce failure on excessive length.                         |
| Unregistered algorithm usage (CWE-693)                        | Registry gate ensures only compiled algorithms are exposed, preventing accidental downgrade to weaker hashes. |
