// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0
// SPDX-License-Identifier: BlueOak-1.0.0

package jsf32

import (
	"math/rand"
	"testing"
)

func BenchmarkUint32(b *testing.B) {
	src := New(0)
	for i := 0; i < b.N; i++ {
		src.Uint32()
	}
}

func BenchmarkUint64(b *testing.B) {
	src := New(0)
	for i := 0; i < b.N; i++ {
		src.Uint64()
	}
}

type int31Test struct {
	Seed   int64
	Values []int32
}

var int31Tests = []int31Test{
	{
		Seed: 0,
		Values: []int32{
			0x1A9B6C07,
			0x1A550895,
			0x712BE876,
			0x902BA19,
			0x20F1A244,
			0x32BC5D2,
			0xBFDB9A1,
			0x7384175A,
			0x16A0F7E5,
			0x470AD8F6,
		},
	},
	{
		Seed: 1,
		Values: []int32{
			0x225132F4,
			0x1EFA0761,
			0x332B56B3,
			0x51AEDB87,
			0x4C4D7156,
			0x3663157A,
			0x1B0A0C8A,
			0x173762FE,
			0x5DE060EC,
			0x17E08DEC,
		},
	},
	{
		Seed: -1,
		Values: []int32{
			0x3EA8325D,
			0x3428F0F3,
			0x61294FA5,
			0x5E2DD8D2,
			0x5555D2D6,
			0x45161F91,
			0x60883E94,
			0x58EB3176,
			0x1613C4D1,
			0x70AA1C8C,
		},
	},
}

func TestJSF32_Int31(t *testing.T) {
	for _, test := range int31Tests {
		src := New(test.Seed)
		for _, expected := range test.Values {
			u := src.Int31()
			if u != expected {
				t.Fatalf("seed %d, expected 0x%X, got 0x%X", test.Seed, expected, u)
			}
		}
	}
}

type int63Test struct {
	Seed   int64
	Values []int64
}

var int63Tests = []int63Test{
	{
		Seed: 0,
		Values: []int64{
			0x1A9B6C079A550895,
			0x712BE8760902BA19,
			0x20F1A244832BC5D2,
			0xBFDB9A17384175A,
			0x16A0F7E5470AD8F6,
		},
	},
	{
		Seed: 1,
		Values: []int64{
			0x225132F41EFA0761,
			0x332B56B3D1AEDB87,
			0x4C4D7156B663157A,
			0x1B0A0C8A973762FE,
			0x5DE060EC17E08DEC,
		},
	},
	{
		Seed: -1,
		Values: []int64{
			0x3EA8325DB428F0F3,
			0x61294FA5DE2DD8D2,
			0x5555D2D6C5161F91,
			0x60883E9458EB3176,
			0x1613C4D170AA1C8C,
		},
	},
}

func TestJSF32_Int63(t *testing.T) {
	for _, test := range int63Tests {
		var src rand.Source64 = New(test.Seed)
		rng := rand.New(src)
		for _, expected := range test.Values {
			i := rng.Int63()
			if i != expected {
				t.Fatalf("seed %d, expected 0x%X, got 0x%X", test.Seed, expected, i)
			}
		}
	}
}

type uint32Test struct {
	Seed   int64
	Values []uint32
}

var uint32Tests = []uint32Test{
	{
		Seed: 0,
		Values: []uint32{
			0x1A9B6C07,
			0x9A550895,
			0xF12BE876,
			0x902BA19,
			0x20F1A244,
			0x832BC5D2,
			0xBFDB9A1,
			0x7384175A,
			0x96A0F7E5,
			0x470AD8F6,
		},
	},
	{
		Seed: 1,
		Values: []uint32{
			0xA25132F4,
			0x1EFA0761,
			0x332B56B3,
			0xD1AEDB87,
			0x4C4D7156,
			0xB663157A,
			0x9B0A0C8A,
			0x973762FE,
			0xDDE060EC,
			0x17E08DEC,
		},
	},
	{
		Seed: -1,
		Values: []uint32{
			0xBEA8325D,
			0xB428F0F3,
			0x61294FA5,
			0xDE2DD8D2,
			0x5555D2D6,
			0xC5161F91,
			0x60883E94,
			0x58EB3176,
			0x9613C4D1,
			0x70AA1C8C,
		},
	},
}

func TestJSF32_Uint32(t *testing.T) {
	for _, test := range uint32Tests {
		src := New(test.Seed)
		for _, expected := range test.Values {
			u := src.Uint32()
			if u != expected {
				t.Fatalf("seed %d, expected 0x%X, got 0x%X", test.Seed, expected, u)
			}
		}
	}
}

type uint64Test struct {
	Seed   int64
	Values []uint64
}

var uint64Tests = []uint64Test{
	{
		Seed: 0,
		Values: []uint64{
			0x1A9B6C079A550895,
			0xF12BE8760902BA19,
			0x20F1A244832BC5D2,
			0x0BFDB9A17384175A,
			0x96A0F7E5470AD8F6,
		},
	},
	{
		Seed: 1,
		Values: []uint64{
			0xA25132F41EFA0761,
			0x332B56B3D1AEDB87,
			0x4C4D7156B663157A,
			0x9B0A0C8A973762FE,
			0xDDE060EC17E08DEC,
		},
	},
	{
		Seed: -1,
		Values: []uint64{
			0xBEA8325DB428F0F3,
			0x61294FA5DE2DD8D2,
			0x5555D2D6C5161F91,
			0x60883E9458EB3176,
			0x9613C4D170AA1C8C,
		},
	},
}

func TestJSF32_Uint64(t *testing.T) {
	for _, test := range uint64Tests {
		var src rand.Source64 = New(test.Seed)
		rng := rand.New(src)
		for _, expected := range test.Values {
			u := rng.Uint64()
			if u != expected {
				t.Fatalf("seed %d, expected 0x%X, got 0x%X", test.Seed, expected, u)
			}
		}
	}
}
