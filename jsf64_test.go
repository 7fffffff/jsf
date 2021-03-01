// Licensed under the Blue Oak Model License 1.0.0:
// https://blueoakcouncil.org/license/1.0.0
// SPDX-License-Identifier: BlueOak-1.0.0

package jsf

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
			0x338FCDF5,
			0x19833F29,
			0x5D0254CE,
			0x454692BB,
			0x6DF72D7B,
			0x70D16D17,
			0x43861DE5,
			0x7E56976D,
			0x5FE204B8,
			0x1EF4D35E,
		},
	},
	{
		Seed: 1,
		Values: []int32{
			0xD060948,
			0x68563732,
			0x6AA2DA64,
			0x6013CC82,
			0x4FD11C9,
			0x5B91B892,
			0x140BF0B1,
			0x4E9766A1,
			0x6D42FFB0,
			0x7C83DD19,
		},
	},
	{
		Seed: -1,
		Values: []int32{
			0x7DC94959,
			0x34EE9963,
			0x788A1CD0,
			0x10D37021,
			0x6D4845C0,
			0x17E4FF4A,
			0x11749778,
			0x5CBE734,
			0x78B9965A,
			0x5CAD4813,
		},
	},
}

func TestJSF64_Int31(t *testing.T) {
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
			0x4B39C42DB38FCDF5,
			0x2EE2C9E919833F29,
			0x30611CD75D0254CE,
			0x7FCFD4F0C54692BB,
			0x358F7AE8EDF72D7B,
			0x4037D431F0D16D17,
			0x582FE1D343861DE5,
			0x3B0CC6607E56976D,
			0x784713275FE204B8,
			0x3025FFDB9EF4D35E,
		},
	},
	{
		Seed: 1,
		Values: []int64{
			0x2E735CA10D060948,
			0xE16AA0268563732,
			0xF061CF1EAA2DA64,
			0x615AE6DDE013CC82,
			0x1FD97CA404FD11C9,
			0x60F02311DB91B892,
			0x63A57B41940BF0B1,
			0x500285784E9766A1,
			0x1FDCAB09ED42FFB0,
			0x1C1A1BE77C83DD19,
		},
	},
	{
		Seed: -1,
		Values: []int64{
			0x28E6401BFDC94959,
			0x467C7D34B4EE9963,
			0x459C60A6F88A1CD0,
			0x6DAB3B4210D37021,
			0x1C790351ED4845C0,
			0xFC36E8D97E4FF4A,
			0x71F33E6291749778,
			0x10D1CB6905CBE734,
			0x2A893C50F8B9965A,
			0x58191B8C5CAD4813,
		},
	},
}

func TestJSF64_Int63(t *testing.T) {
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
			0xB38FCDF5,
			0x19833F29,
			0x5D0254CE,
			0xC54692BB,
			0xEDF72D7B,
			0xF0D16D17,
			0x43861DE5,
			0x7E56976D,
			0x5FE204B8,
			0x9EF4D35E,
		},
	},
	{
		Seed: 1,
		Values: []uint32{
			0xD060948,
			0x68563732,
			0xEAA2DA64,
			0xE013CC82,
			0x4FD11C9,
			0xDB91B892,
			0x940BF0B1,
			0x4E9766A1,
			0xED42FFB0,
			0x7C83DD19,
		},
	},
	{
		Seed: -1,
		Values: []uint32{
			0xFDC94959,
			0xB4EE9963,
			0xF88A1CD0,
			0x10D37021,
			0xED4845C0,
			0x97E4FF4A,
			0x91749778,
			0x5CBE734,
			0xF8B9965A,
			0x5CAD4813,
		},
	},
}

func TestJSF64_Uint32(t *testing.T) {
	for _, test := range uint32Tests {
		var src = New(test.Seed)
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
			0x4B39C42DB38FCDF5,
			0xAEE2C9E919833F29,
			0x30611CD75D0254CE,
			0x7FCFD4F0C54692BB,
			0xB58F7AE8EDF72D7B,
			0x4037D431F0D16D17,
			0x582FE1D343861DE5,
			0xBB0CC6607E56976D,
			0x784713275FE204B8,
			0xB025FFDB9EF4D35E,
		},
	},
	{
		Seed: 1,
		Values: []uint64{
			0xAE735CA10D060948,
			0x8E16AA0268563732,
			0x8F061CF1EAA2DA64,
			0xE15AE6DDE013CC82,
			0x1FD97CA404FD11C9,
			0xE0F02311DB91B892,
			0x63A57B41940BF0B1,
			0x500285784E9766A1,
			0x1FDCAB09ED42FFB0,
			0x1C1A1BE77C83DD19,
		},
	},
	{
		Seed: -1,
		Values: []uint64{
			0xA8E6401BFDC94959,
			0xC67C7D34B4EE9963,
			0x459C60A6F88A1CD0,
			0xEDAB3B4210D37021,
			0x1C790351ED4845C0,
			0x8FC36E8D97E4FF4A,
			0x71F33E6291749778,
			0x90D1CB6905CBE734,
			0x2A893C50F8B9965A,
			0x58191B8C5CAD4813,
		},
	},
}

func TestJSF64_Uint64(t *testing.T) {
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
