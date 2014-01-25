package arg

import "flag"

var visual *bool
var rounds *int
var delayms *int
var verbose *bool
var maxperms *int

func BindArgs() {
	visual = flag.Bool("visual", false, "")
	rounds = flag.Int("rounds", 1000, "")
	delayms = flag.Int("delay", 500, "")
	verbose = flag.Bool("verbose", false, "")
	maxperms = flag.Int("maxperm", -1, "")
	flag.Parse()
}

func Visual() bool {
	return *visual
}

func NumRounds() int {
	return *rounds
}

func DelayMS() int {
	return *delayms
}

func Verbose() bool {
	return *verbose
}

func MaxPermutations() int {
	return *maxperms
}
