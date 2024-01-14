package main

type Theme struct {
	Restricted   []int
	Unrestricted []int
}

var schemes map[string]Theme

func init() {
	schemes = map[string]Theme{
		"TokyoNight": {
			Restricted: []int{
				0x803d49,
				0xb9f27c,
				0x618041,
				0x7da6ff,
			},
			Unrestricted: []int{
				0x06080a,
				0x1a1b26,
				0x232433,
				0x2a2b3d,
				0x32344a,
				0x3b3d57,
				0xff7a93,
				0x3e5380,
				0xa9b1d6,
				0xF7768E,
				0xFF9E64,
				0xE0AF68,
				0x9ECE6A,
				0x7AA2F7,
				0xad8ee6,
				0x444B6A},
		},
		"Pastel": {
			Restricted: []int{},
			Unrestricted: []int{
				0xd6e6ff,
				0xd7f9f8,
				0xffffea,
				0xfff0d4,
				0xfbe0e0,
				0xe5d4ef,
			},
		},
		"rose-pine": {
			Restricted: []int{},
			Unrestricted: []int{
				0x16141f,
				0x191724,
				0x1f1d2e,
				0x26233a,
				0x6e6a86,
				0x908caa,
				0xe0def4,
				0xeb6f92,
				0xf6c177,
				0xebbcba,
				0x31748f,
				0x9ccfd8,
				0xc4a7e7,
				0x21202e,
				0x403d52,
				0x524f67,
			},
		},
		"rose-pine-dawn": {
			Restricted: []int{},
			Unrestricted: []int{

				0x575279,
				0xb4637a,
				0xea9d34,
				0xd7827e,
				0x286983,
				0x56949f,
				0x907aa9,
				0xf4ede8,
				0xdfdad9,
				0xf8f0e7,
				0xfaf4ed,
				0xfffaf3,
				0xf2e9e1,
				0x9893a5,
				0x797593,
				0xcecacd,
			},
		},
	}
}
