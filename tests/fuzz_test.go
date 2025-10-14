// SPDX-License-Identifier: MIT
//
// Copyright (C) 2025 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package tests

import (
	"bytes"
	"crypto"
	"testing"

	"github.com/bytemare/hash"
)

func FuzzHash(f *testing.F) {
	f.Fuzz(func(t *testing.T, id uint8, cryptoID uint, size int, prefix, input []byte) {
		_ = hash.FromCrypto(crypto.Hash(cryptoID))
		h := hash.Hash(id)

		if h.Available() {
			_ = h.String()
			_ = h.BlockSize()
			_ = h.Size()
			_ = h.SecurityLevel()

			switch h.Type() {
			case hash.FixedOutputLength:
				if h.GetXOF() != nil {
					t.Fatal("unexpected xof for fixed-length hash ID")
				}
			case hash.ExtendableOutputFunction:
				if h.GetHashFunction() != nil {
					t.Fatal("unexpected fixed-length hash ID for xof")
				}
			default:
				t.Fatal("unknown hash type")
			}

			hasher := h.New()
			if !bytes.Equal(h.Hash(input), hasher.Hash(input)) {
				t.Fatal("unexpected hash output")
			}
		}
	})
}
