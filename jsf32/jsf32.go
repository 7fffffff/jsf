// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0
// SPDX-License-Identifier: BlueOak-1.0.0

// Package jsf32 provides the 32-bit version of Bob Jenkins'
// public-domain small noncryptographic PRNG
//
// The original C implementation can be found here:
// https://burtleburtle.net/bob/rand/smallprng.html
package jsf32

// JSF32 implements the math/rand.Source64 interface
//
// Callers must seed a generator through New() or Seed() before
// generating any numbers. An unseeded JSF32 (the zero value) will
// always produce a 0
type JSF32 struct {
	a uint32
	b uint32
	c uint32
	d uint32
}

// New returns a seeded generator
func New(seed int64) *JSF32 {
	j := &JSF32{}
	j.Seed(seed)
	return j
}

// Int31 returns a pseudorandom value in [0, 1<<31)
func (j *JSF32) Int31() int32 {
	u := j.Uint32()
	return int32(u & 0x7FFFFFFF)
}

// Int63 returns a pseudorandom value in [0, 1<<63)
func (j *JSF32) Int63() int64 {
	u := j.Uint64()
	return int64(u & 0x7FFFFFFFFFFFFFFF)
}

// Seed (re)initializes the generator.
// Only the lower 32 bits of seed are used.
func (j *JSF32) Seed(seed int64) {
	j.a = 0xf1ea5eed
	j.b = uint32(seed)
	j.c = uint32(seed)
	j.d = uint32(seed)
	for i := 0; i < 20; i++ {
		j.Uint32()
	}
}

// Uint32 returns a pseudorandom value in [0, 1<<32)
func (j *JSF32) Uint32() uint32 {
	e := j.a - rot32l(j.b, 27)
	j.a = j.b ^ rot32l(j.c, 17)
	j.b = j.c + j.d
	j.c = j.d + e
	j.d = e + j.a
	return j.d
}

// Uint64 returns a pseudorandom value in [0, 1<<64)
func (j *JSF32) Uint64() uint64 {
	x := uint64(j.Uint32())
	y := uint64(j.Uint32())
	return (x << 32) | y
}

func rot32l(x, k uint32) uint32 {
	return (x << k) | (x >> (32 - k))
}
