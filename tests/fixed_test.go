// SPDX-License-Identifier: MIT
//
// Copyright (C) 2025 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package tests

import (
	"crypto/fips140"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"github.com/bytemare/hash"

	cryptorand "crypto/rand"
)

var errHmacKeySize = errors.New("hmac key length is larger than hash output size")

func TestHmac(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()

			key, _ := hex.DecodeString(testData.key[h.HashID.Size()])
			hmac := hasher.Hmac(testData.message, key)

			if len(hmac) != h.HashID.Size() {
				t.Errorf("#%v : invalid hmac length", h.HashID)
			}
		}
	})
}

func TestLongHmacKey(t *testing.T) {
	longHMACKey := []byte("Length65aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()

			if panics, err := expectPanic(errHmacKeySize, func() {
				_ = hasher.Hmac(testData.message, longHMACKey)
			}); !panics {
				t.Errorf("expected panic: %v", err)
			}
		}
	})
}

func TestHKDF(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()

			for _, l := range []int{0, h.HashID.Size()} {
				key, err := hasher.HKDF(testData.secret, testData.salt, testData.info, l)
				if err != nil {
					t.Fatal(err)
				}

				if len(key) != h.HashID.Size() {
					t.Errorf("#%v : invalid key length (length argument = %d)", h.HashID, l)
				}
			}
		}
	})
}

func TestHKDFExtract(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()

			for _, l := range []int{0, h.HashID.Size()} {
				// Build a pseudorandom key
				prk, err := hasher.HKDFExtract(testData.secret, testData.salt)
				if err != nil {
					t.Fatal(err)
				}

				if len(prk) != h.HashID.Size() {
					t.Errorf("%v : invalid key length (length argument = %d)", h.HashID, l)
				}
			}
		}
	})
}

func TestHKDFExpand(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()

			for _, l := range []int{0, h.HashID.Size()} {
				// Build a pseudorandom key
				prk, err := hasher.HKDFExtract(testData.secret, testData.salt)
				if err != nil {
					t.Fatal(err)
				}

				key, err := hasher.HKDFExpand(prk, testData.info, l)
				if err != nil {
					t.Fatal(err)
				}

				if len(key) != h.HashID.Size() {
					t.Errorf("#%v : invalid key length (length argument = %d)", h.HashID, l)
				}
			}
		}
	})
}

func TestHKDF_LengthTooLong(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()
			limit := hasher.Size() * 255

			_, err := hasher.HKDF(testData.secret, testData.salt, testData.info, limit+1)
			if err == nil {
				t.Fatal("expected error on excessive length")
			}
		}
	})
}

func TestHKDFExtract_KeyTooShort(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()
			shortKey := randomBytes((112 / 8) - 1) // threshold is 14 bytes

			_, err := hasher.HKDFExtract(shortKey, testData.salt)
			if fips140.Enabled() && err == nil {
				t.Fatal("expected error on excessive length")
			}
		}
	})
}

func TestHKDFExtract_NotFIPS140Hash(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()
			switch hasher.Algorithm() {
			case hash.SHA256, hash.SHA384, hash.SHA512:
				return
			}

			_, err := hasher.HKDFExtract(testData.secret, testData.salt)
			if fips140.Enabled() && err == nil {
				t.Fatal("expected error on non FIPS140 hash")
			}
		}
	})
}

func TestHKDFExpand_LengthTooLong(t *testing.T) {
	testAll(t, func(h *testHash) {
		if h.HashType == hash.FixedOutputLength {
			hasher := h.HashID.GetHashFunction()
			limit := hasher.Size() * 255

			_, err := hasher.HKDFExpand(testData.secret, testData.info, limit+1)
			if fips140.Enabled() && err == nil {
				t.Fatal("expected error on excessive length")
			}
		}
	})
}

func randomBytes(length int) []byte {
	random := make([]byte, length)
	if _, err := cryptorand.Read(random); err != nil {
		// We can as well not panic and try again in a loop
		panic(fmt.Errorf("unexpected error in generating random bytes : %w", err))
	}

	return random
}
