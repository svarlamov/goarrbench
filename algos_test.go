package main

import (
	"testing"
)

var ids []string
var legitCount int = 5000
var localCount int = 5

func init() {
	baseLocal := "spotify:local:::Diplo+x+Alvaro+-+6th+Gear+FINAL+MASTER"
	baseLegit := "spotify:track:4iV5W9uYEdYUVa79Axb7Rh"
	ids = make([]string, (legitCount + localCount))
	for i := 0; i < legitCount; i++ {
		ids[i] = baseLegit
	}
	for j := legitCount; j < localCount+legitCount; j++ {
		ids[j] = baseLocal
	}
}

func BenchmarkMakeNewIdsArraySameCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = MakeNewIdsArraySameCap(ids)
	}
}

func BenchmarkMakeNewIdsArrayDiffCapWithAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = MakeNewIdsArrayDiffCapWithAppend(ids)
	}
}

func BenchmarkPluckFromBaseArrayWithCounterAndAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = PluckFromBaseArrayWithCounterAndAppend(ids)
	}
}
