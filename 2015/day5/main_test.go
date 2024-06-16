package main

import (
	"log"
	"testing"
)

func TestNice(t *testing.T) {
	testCases := []struct {
		desc string
		got  string
		want bool
	}{
		{
			desc: "Naughty String, with bad strings",
			got:  "haegwjzuvuyypxyu",
			want: false,
		},
		{
			desc: "Naughty String, with not enough vowels",
			got:  "dvszwmarrgswjxmb",
			want: false,
		},
		{
			desc: "Naughty String, with no double letter",
			got:  "jchzalrnumimnmhp",
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := nice(tC.got)
			if got != tC.want {
				log.Fatalf("%s, got %v, want %v\n", tC.desc, got, tC.want)
			}
		})
	}
}

func TestCheckVowels(t *testing.T) {
	testCases := []struct {
		desc string
		got  string
		want bool
	}{
		{
			desc: "Naughty String, with not enough vowels",
			got:  "dvszwmarrgswjxmb",
			want: false,
		},
		{
			desc: "nice String, with enough vowels",
			got:  "ugknbfddgicrmopn",
			want: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := checkVowels(tC.got)
			if got != tC.want {
				log.Fatalf("%s, got %v, want %v\n", tC.desc, got, tC.want)
			}
		})
	}
}
