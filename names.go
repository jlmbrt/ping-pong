package main

import (
	"math/rand"
)

var names = []string{
	"G3n1us-B0t",
	"R0ck3t-M4x",
	"C0d3-Break3r",
	"Z4p-Unit",
	"Bl4st0-Mat1c",
	"Pr0t0-Sentry",
	"Scr4p-Tr0n",
	"W1r3fr4m3",
	"M3t4l-H3ad",
	"C1rcu1t-Surge",
}

func generateName() string {
	idx := rand.Intn(len(names))

	return names[idx]
}
