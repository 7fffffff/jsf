// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0
// SPDX-License-Identifier: BlueOak-1.0.0

// Package jsf provides the 64-bit, 3-rotation variant of Bob
// Jenkins' public-domain small noncryptographic PRNG
//
// The original C implementation can be found here:
// https://burtleburtle.net/bob/rand/smallprng.html
package jsf

// JSF64 implements the math/rand.Source64 interface
//
// Callers must seed a generator through New() or Seed() before
// generating any numbers. An unseeded JSF64 (the zero value) will
// always produce a 0
type JSF64 struct {
	a uint64
	b uint64
	c uint64
	d uint64
}

// New returns a seeded generator
func New(seed int64) *JSF64 {
	j := &JSF64{}
	j.Seed(seed)
	return j
}

// Int31 returns a pseudorandom value in [0, 1<<31)
func (j *JSF64) Int31() int32 {
	u := j.Uint64()
	return int32(u & 0x7FFFFFFF)
}

// Int63 returns a pseudorandom value in [0, 1<<63)
func (j *JSF64) Int63() int64 {
	u := j.Uint64()
	return int64(u & 0x7FFFFFFFFFFFFFFF)
}

// Seed (re)initializes the generator
func (j *JSF64) Seed(seed int64) {
	j.a = 0xf1ea5eed
	j.b = uint64(seed)
	j.c = uint64(seed)
	j.d = uint64(seed)
	for i := 0; i < 20; i++ {
		j.Uint64()
	}
}

// Uint32 returns a pseudorandom value in [0, 1<<32)
func (j *JSF64) Uint32() uint32 {
	return uint32(j.Uint64())
}

// Uint64 returns a pseudorandom value in [0, 1<<64)
func (j *JSF64) Uint64() uint64 {
	e := j.a - (rot64l(j.b, 7))
	j.a = j.b ^ rot64l(j.c, 13)
	j.b = j.c + rot64l(j.d, 37)
	j.c = j.d + e
	j.d = e + j.a
	return j.d
}

func rot64l(x, k uint64) uint64 {
	return (x << k) | (x >> (64 - k))
}
