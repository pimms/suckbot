package arg

import "flag"

var rounds *int
var delayms *int
var verbose *bool
var maxperms *int
var file *string

func BindArgs() {
	rounds = flag.Int("rounds", 1000, "The number of ticks in each simulation")
	delayms = flag.Int("delay", 500, "The delay between each tick in visual mode")
	verbose = flag.Bool("verbose", false, "Verbose output to STDOUT")
	maxperms = flag.Int("maxperm", -1, "The maximum number of permutations (defaults to N*2^N)")
	file = flag.String("file", "default", "The environment file to load from")
	flag.Parse()
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

func File() string {
	return *file
}
