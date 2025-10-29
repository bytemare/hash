# Design and Security Overview

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

| Weakness                                                      | Implemented countermeasures                                                                                   |
|---------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| Improper key length enforcement (CWE-325)                     | HMAC helper panics on oversized keys and tests ensure panic occurs.                                           |
| Insufficient output length for cryptographic hashes (CWE-326) | Extendable hash `Read` verifies requested size to prevent truncated outputs.                                  |
| Exceeding HKDF block limits (CWE-331)                         | HKDF methods propagate upstream errors and tests enforce failure on excessive length.                         |
| Unregistered algorithm usage (CWE-693)                        | Registry gate ensures only compiled algorithms are exposed, preventing accidental downgrade to weaker hashes. |
