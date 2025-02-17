// SPDX-License-Identifier: MIT
//
// Copyright (C) 2024 Daniel Bourdrez. All Rights Reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree or at
// https://spdx.org/licenses/MIT.html

package hash

import (
	"errors"
	"io"

	"crypto/sha3"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

const (
	// string IDs for the hash functions.
	shake128 = "SHAKE128"
	shake256 = "SHAKE256"
	blake2xb = "BLAKE2XB"
	blake2xs = "BLAKE2XS"

	// block size in bytes.
	blockSHAKE128 = (1600 - 256) / 8
	blockSHAKE256 = (1600 - 512) / 8
)

var errSmallOutputSize = errors.New("requested output size too small")

// xof defines the interface to hash functions that support arbitrary-length output.
type xof interface {
	// Writer Write absorbs more data into the hash's state. It panics if called
	// after Read.
	io.Writer

	// Reader Read reads more output from the hash. It returns io.EOF if the limit
	// has been reached.
	io.Reader

	// Reset resets the XOF to its initial state.
	Reset()
}

func newXOF(hid Hash, size int) newHash {
	return func() Hasher {
		ext := &ExtendableHash{
			xof:  nil,
			id:   hid,
			size: outputSizes[hid],
		}

		switch hid {
		case SHAKE128:
			ext.xof = sha3.NewSHAKE128()
		case SHAKE256:
			ext.xof = sha3.NewSHAKE256()
		case BLAKE2XB:
			ext.xof, _ = blake2b.NewXOF(uint32(size), nil)
		case BLAKE2XS:
			ext.xof, _ = blake2s.NewXOF(uint16(size), nil)
		}

		return ext
	}
}

// ExtendableHash offers easy an easy-to-use API for common
// cryptographic hash operations of extendable output functions.
type ExtendableHash struct {
	xof
	id   Hash
	size int
}

// Algorithm returns the Hash function identifier.
func (h *ExtendableHash) Algorithm() Hash {
	return h.id
}

// Hash returns the hash of the input argument with size output length.
func (h *ExtendableHash) Hash(input ...[]byte) []byte {
	h.Reset()

	for _, i := range input {
		_, _ = h.Write(i)
	}

	return h.Read(h.size)
}

// Read consumes and returns size bytes from the current hash.
func (h *ExtendableHash) Read(size int) []byte {
	if size != 0 && size < h.Size() {
		panic(errSmallOutputSize)
	}

	output := make([]byte, size)
	_, _ = h.xof.Read(output)

	return output
}

// Write implements io.Writer.
func (h *ExtendableHash) Write(input []byte) (int, error) {
	return h.xof.Write(input)
}

// Sum appends the current hash to b and returns the resulting slice.
func (h *ExtendableHash) Sum(prefix []byte) []byte {
	output := make([]byte, h.id.Size()+len(prefix))
	copy(output, prefix)
	_, _ = h.xof.Read(output[len(prefix):])

	return output
}

// Reset resets the hash to its initial state.
func (h *ExtendableHash) Reset() {
	h.xof.Reset()
}

// SetOutputSize sets the output size for an ExtendableOutputFunction, and is a no-op for fixed hashing.
func (h *ExtendableHash) SetOutputSize(size int) {
	switch h.id {
	case BLAKE2XB:
		x, err := blake2b.NewXOF(uint32(size), nil)
		if err != nil {
			panic(err)
		}

		h.xof = x
	case BLAKE2XS:
		x, err := blake2s.NewXOF(uint16(size), nil)
		if err != nil {
			panic(err)
		}

		h.xof = x
	}

	h.size = size
}

// Size returns the number of bytes Hash will return.
func (h *ExtendableHash) Size() int {
	return h.size
}

// BlockSize returns the hash's underlying block size.
func (h *ExtendableHash) BlockSize() int {
	return h.id.BlockSize()
}

// GetHashFunction returns nil.
func (h *ExtendableHash) GetHashFunction() *Fixed {
	return nil
}

// GetXOF returns the underlying ExtendableHash Hasher.
func (h *ExtendableHash) GetXOF() *ExtendableHash {
	return h
}
