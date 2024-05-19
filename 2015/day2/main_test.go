package main

import "testing"

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		desc string
		got  RectagularPrism
		want int
	}{
		{
			desc: "",
			got:  RectagularPrism{2, 3, 4},
			want: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.got.perimeter([]int{2, 3})
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestSmallestPerimeter(t *testing.T) {
	testCases := []struct {
		desc string
		got  RectagularPrism
		want int
	}{
		{
			desc: "",
			got:  RectagularPrism{2, 3, 4},
			want: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.got.smallestPerimeter()
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestSrufaceArea(t *testing.T) {
	testCases := []struct {
		desc string
		got  RectagularPrism
		want int
	}{
		{
			desc: "",
			got:  RectagularPrism{2, 3, 4},
			want: 52,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.got.surfaceArea()
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestSmallestSideArea(t *testing.T) {
	testCases := []struct {
		desc string
		got  RectagularPrism
		want int
	}{
		{
			desc: "Present 2x3x4",
			got:  RectagularPrism{2, 3, 4},
			want: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.got.smallestSideArea()
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestVolume(t *testing.T) {
	testCases := []struct {
		desc string
		got  RectagularPrism
		want int
	}{
		{
			desc: "",
			got:  RectagularPrism{2, 3, 4},
			want: 24,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.got.volume()
			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	testCases := []struct {
		desc string
		got  string
		want int
	}{
		{
			desc: "Present 2x3x4, sq 52, slack 6",
			got:  "2x3x4",
			want: 58,
		},
		{
			desc: "Present 1x1x10, sq 41, slack 1",
			got:  "1x1x10",
			want: 43,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := parseLine(tC.got)

			if got != tC.want {
				t.Errorf("%s, got %d, want %d", tC.desc, got, tC.want)
			}
		})
	}
}
