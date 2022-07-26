// SPDX-License-Identifier: MIT
//
// Copyright (C) 2021 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package hash_test

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/bytemare/hash"
)

var errHmacKeySize = errors.New("hmac key length is larger than hash output size")

func TestLongHmacKey(t *testing.T) {
	longHMACKey := []byte("Length65aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	for _, id := range []hash.Hashing{hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512} {
		h := id.Get()

		if hasPanic, err := expectPanic(errHmacKeySize, func() {
			_ = h.Hmac(testData.message, longHMACKey)
		}); !hasPanic {
			t.Fatalf("expected panic: %v", err)
		}
	}
}

func TestHmac(t *testing.T) {
	for _, id := range []hash.Hashing{hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512} {
		h := id.Get()

		key, _ := hex.DecodeString(testData.key[h.OutputSize()])
		hmac := h.Hmac(testData.message, key)

		if len(hmac) != h.OutputSize() {
			t.Errorf("#%v : invalid hmac length", id)
		}
	}
}

func TestHKDF(t *testing.T) {
	for _, id := range []hash.Hashing{hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512} {
		h := id.Get()

		for _, l := range []int{0, h.OutputSize()} {
			key := h.HKDF(testData.secret, testData.salt, testData.info, l)

			if len(key) != h.OutputSize() {
				t.Errorf("#%v : invalid key length (length argument = %d)", id, l)
			}
		}
	}
}

func TestHKDFExtract(t *testing.T) {
	for _, id := range []hash.Hashing{hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512} {
		h := id.Get()

		for _, l := range []int{0, h.OutputSize()} {
			// Build a pseudorandom key
			prk := h.HKDFExtract(testData.secret, testData.salt)

			if len(prk) != h.OutputSize() {
				t.Errorf("#%v : invalid key length (length argument = %d)", id, l)
			}
		}
	}
}

func TestHKDFExpand(t *testing.T) {
	for _, id := range []hash.Hashing{hash.SHA256, hash.SHA512, hash.SHA3_256, hash.SHA3_512} {
		h := id.Get()

		for _, l := range []int{0, h.OutputSize()} {
			// Build a pseudorandom key
			prk := h.HKDF(testData.secret, testData.salt, testData.info, l)
			key := h.HKDFExpand(prk, testData.info, l)

			if len(key) != h.OutputSize() {
				t.Errorf("#%v : invalid key length (length argument = %d)", id, l)
			}
		}
	}
}
