package arg

import "flag"

var visual *bool
var rounds *int
var delayms *int
var verbose *bool
var maxperms *int
var file *string
var nomoredirt *bool
var manual *bool

func BindArgs() {
	visual = flag.Bool("visual", false, "Draw the environment graphically")
	rounds = flag.Int("rounds", 1000, "The number of ticks in each simulation")
	delayms = flag.Int("delay", 500, "The delay between each tick in visual mode")
	verbose = flag.Bool("verbose", false, "Verbose output to STDOUT")
	maxperms = flag.Int("maxperm", -1, "The maximum number of permutations (defaults to N*2^N)")
	file = flag.String("file", "default", "The environment file to load from")
	nomoredirt = flag.Bool("nmd", false, "Dirt will not reappear in tiles")
	manual = flag.Bool("manual", false, "Visual mode will not progress until enter is pressed")
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

func File() string {
	return *file
}

func NoMoreDirt() bool {
	return *nomoredirt
}

func Manual() bool {
	return *manual
}

