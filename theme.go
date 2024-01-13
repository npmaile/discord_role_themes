package main

type Theme struct {
	Restricted   []int
	Unrestricted []int
}

var schemes map[string]Theme

func init(){
	schemes = map[string]Theme{
		"TokyoNight": Theme{
			Restricted:   []int{0x1ABC9C},
			Unrestricted: []int{0x57F287,0x11806A},
		},
	}
}
