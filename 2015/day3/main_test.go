package main

import (
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
			desc: "Single direction, Single slice with 1 element",
			got:  ">",
			want: []string{">"},
		},
		{
			desc: "Four direction, Single slice with 4 elements",
			got:  "^>v<",
			want: []string{"^", ">", "v", "<"},
		},
		{
			desc: "Ten direction, Single slice with 10 elements",
			got:  "^v^v^v^v^v",
			want: []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseLine(tC.got)
			if !slices.Equal(got, tC.want) {
				t.Errorf("%s, got %v, want %v", tC.desc, got, tC.want)
			}
		})
	}
}

func TestCountHouses(t *testing.T) {
	testCases := []struct {
		desc string
		got  string
		want int
	}{
		{
			desc: "One direction is 2 houses",
			got:  ">",
			want: 2,
		},
		{
			desc: "Deliver in sqare",
			got:  "^>v<",
			want: 4,
		},
		{
			desc: "Deliver a bunch of presents to 2 houses",
			got:  "^v^v^v^v^v",
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countHouses(tC.got)
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}
