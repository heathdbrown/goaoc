package main

import (
	"log"
	"slices"
	"testing"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		desc string
		got  string
		want []string
	}{
		{
			desc: "Single text to single list",
			got:  "^",
			want: []string{"^"},
		},
		{
			desc: "Single text to single list",
			got:  "^>v<",
			want: []string{"^", ">", "v", "<"},
		},
		{
			desc: "Single text to single list",
			got:  "^v^v^v^v^v",
			want: []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseLine(tC.got)

			if !slices.Equal(got, tC.want) {
				log.Fatalf("%s, got %v, want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestCountHouses(t *testing.T) {
	testCases := []struct {
		desc string
		got  []string
		want int
	}{
		{
			desc: "Single direction is two houses",
			got:  parseLine("^"),
			want: 2,
		},
		{
			desc: "square direction is 4 houses",
			got:  parseLine("^>v<"),
			want: 4,
		},
		{
			desc: "Up down direction is 2 houses",
			got:  parseLine("^v^v^v^v^v"),
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countHouses(tC.got)

			if got != tC.want {
				log.Fatalf("%s, got %v, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestCountSantaAndRoboHouses(t *testing.T) {
	testCases := []struct {
		desc string
		got  []string
		want int
	}{
		{
			desc: "Single direction is two houses",
			got:  parseLine("^v"),
			want: 3,
		},
		{
			desc: "square direction is 4 houses",
			got:  parseLine("^>v<"),
			want: 3,
		},
		{
			desc: "Up down direction is 2 houses",
			got:  parseLine("^v^v^v^v^v"),
			want: 11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countSantaAndRoboHouses(tC.got)

			if got != tC.want {
				log.Fatalf("%s, got %v, want %d", tC.desc, got, tC.want)
			}
		})
	}
}
