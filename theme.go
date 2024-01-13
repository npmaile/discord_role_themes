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
	}
}
