package cardrank

import (
	"reflect"
	"testing"
)

func TestRazz(t *testing.T) {
	tests := []struct {
		v string
		b string
		u string
		r HandRank
	}{
		{"Kh Qh Jh Th 9h Ks Qs", "Kh Qh Jh Th 9h", "Ks Qs", 7936},
		{"Ah Kh Qh Jh Th Ks Qs", "Kh Qh Jh Th Ah", "Ks Qs", 7681},
		{"2h 2c 2d 2s As Ks Qs", "2h 2c As Ks Qs", "2d 2s", 59569},
		{"Ah Ac Ad Ks Kh Ks Qs", "Ah Ac Ks Kh Qs", "Ad Ks", 63067},
		{"Ah Ac Ad Ks Qh Ks Qs", "Ks Ks Qh Qs Ah", "Ac Ad", 62935},
		{"Kh Kd Qd Qs Jh Ks Js", "Qd Qs Jh Js Kh", "Kd Ks", 62813},
		{"3h 3c Kh Qd Jd Ks Qs", "3h 3c Kh Qd Jd", "Ks Qs", 59734},
		{"2h 2c Kh Qd Jd Ks Qs", "2h 2c Kh Qd Jd", "Ks Qs", 59514},
		{"3h 2c Kh Qd Jd Ks Qs", "Kh Qd Jd 3h 2c", "Ks Qs", 7174},
	}
	for i, test := range tests {
		best, unused := Must(test.b), Must(test.u)
		h := Razz.RankHand(Must(test.v), nil)
		if h.HiRank != test.r {
			t.Errorf("test %d %v expected rank %d, got: %d", i, h.Pocket, test.r, h.HiRank)
		}
		if !reflect.DeepEqual(h.HiBest, best) {
			t.Errorf("test %d %v expected best %v, got: %v", i, h.Pocket, best, h.HiBest)
		}
		if !reflect.DeepEqual(h.HiUnused, unused) {
			t.Errorf("test %d %v expected unused %v, got: %v", i, h.Pocket, unused, h.HiUnused)
		}
	}
}

func TestBadugi(t *testing.T) {
	tests := []struct {
		v string
		b string
		u string
		r HandRank
	}{
		{"Kh Qh Jh Th", "Th", "Kh Qh Jh", 25088},
		{"Kh Qh Jd Th", "Jd Th", "Kh Qh", 17920},
		{"Kh Qc Jd Th", "Qc Jd Th", "Kh", 11776},
		{"Ks Qc Jd Th", "Ks Qc Jd Th", "", 7680},
		{"2h 2c 2d 2s", "2s", "2h 2d 2c", 24578},
		{"Ah Kh Qh Jh", "Ah", "Kh Qh Jh", 24577},
		{"Kh Kd Qd Qs", "Kh Qs", "Kd Qd", 22528},
		{"Ah Ac Ad Ks", "Ks Ah", "Ad Ac", 20481},
		{"3h 3c Kh Qd", "Kh Qd 3c", "3h", 14340},
		{"2h 2c Kh Qd", "Kh Qd 2c", "2h", 14338},
		{"3h 2c Kh Ks", "Ks 3h 2c", "Kh", 12294},
		{"3h 2c Kh Qd", "Qd 3h 2c", "Kh", 10246},
		{"Ah 2c 4s 6d", "6d 4s 2c Ah", "", 43},
		{"Ac 2h 4d 6s", "6s 4d 2h Ac", "", 43},
		{"Ah 2c 3s 6d", "6d 3s 2c Ah", "", 39},
		{"Ah 2c 4s 5d", "5d 4s 2c Ah", "", 27},
		{"Ah 2c 3s 5d", "5d 3s 2c Ah", "", 23},
		{"Ah 2c 3s 4d", "4d 3s 2c Ah", "", 15},
		{"Ac 2h 3s 4d", "4d 3s 2h Ac", "", 15},
	}
	for i, test := range tests {
		best, unused := Must(test.b), Must(test.u)
		h := Badugi.RankHand(Must(test.v), nil)
		if h.HiRank != test.r {
			t.Errorf("test %d %v expected rank %d, got: %d", i, h.Pocket, test.r, h.HiRank)
		}
		if !reflect.DeepEqual(h.HiBest, best) {
			t.Errorf("test %d %v expected best %v, got: %v", i, h.Pocket, best, h.HiBest)
		}
		if !reflect.DeepEqual(h.HiUnused, unused) {
			t.Errorf("test %d %v expected unused %v, got: %v", i, h.Pocket, unused, h.HiUnused)
		}
	}
}

func TestLowball(t *testing.T) {
	tests := []struct {
		v string
		b string
		u string
		r HandRank
	}{
		{"7h 5h 4h 3h 2c", "7h 5h 4h 3h 2c", "", 1},
		{"7h 6h 4h 3h 2c", "7h 6h 4h 3h 2c", "", 2},
		{"7h 6h 5h 3h 2c", "7h 6h 5h 3h 2c", "", 3},
		{"7h 6h 5h 4h 2c", "7h 6h 5h 4h 2c", "", 4},
		{"8h 5h 4h 3h 2c", "8h 5h 4h 3h 2c", "", 5},
		{"8h 6h 4h 3h 2c", "8h 6h 4h 3h 2c", "", 6},
		{"8h 6h 5h 3h 2c", "8h 6h 5h 3h 2c", "", 7},
		{"8h 6h 5h 4h 2c", "8h 6h 5h 4h 2c", "", 8},
		{"8h 6h 5h 4h 3c", "8h 6h 5h 4h 3c", "", 9},
		{"8h 7h 4h 3h 2c", "8h 7h 4h 3h 2c", "", 10},
		{"8h 7h 5h 3h 2c", "8h 7h 5h 3h 2c", "", 11},
		{"8h 7h 5h 4h 2c", "8h 7h 5h 4h 2c", "", 12},
		{"8h 7h 5h 4h 3c", "8h 7h 5h 4h 3c", "", 13},
		{"8h 7h 6h 3h 2c", "8h 7h 6h 3h 2c", "", 14},
		{"8h 7h 6h 4h 2c", "8h 7h 6h 4h 2c", "", 15},
		{"8h 7h 6h 4h 3c", "8h 7h 6h 4h 3c", "", 16},
		{"8h 7h 6h 5h 2c", "8h 7h 6h 5h 2c", "", 17},
		{"8h 7h 6h 5h 3c", "8h 7h 6h 5h 3c", "", 18},
		{"9h 5h 4h 3h 2c", "9h 5h 4h 3h 2c", "", 19},
	}
	for i, test := range tests {
		best, unused := Must(test.b), Must(test.u)
		h := Lowball.RankHand(Must(test.v), nil)
		if h.HiRank != test.r {
			t.Errorf("test %d %v expected rank %d, got: %d", i, h.Pocket, test.r, h.HiRank)
		}
		if !reflect.DeepEqual(h.HiBest, best) {
			t.Errorf("test %d %v expected best %v, got: %v", i, h.Pocket, best, h.HiBest)
		}
		if !reflect.DeepEqual(h.HiUnused, unused) {
			t.Errorf("test %d %v expected unused %v, got: %v", i, h.Pocket, unused, h.HiUnused)
		}
	}
}

func TestTypeHiComp(t *testing.T) {
	tests := []struct {
		typ   Type
		board string
		a     string
		b     string
		j     HandRank
		k     HandRank
		exp   int
	}{
		{Short, "As 7d Ad 6s 6d", "8d Td", "Ac 5h", Flush, FullHouse, -1},
		{Short, "As 7d Ad 6s 6d", "Ac 5h", "8d Td", FullHouse, Flush, +1},
		{Short, "Kc Qh Jc Td 8d", "Ac 5h", "Ah 6c", Straight, Straight, 0},
		{Short, "Kc Qh Jc Td 8d", "Ah 6c", "Ac 5h", Straight, Straight, 0},
		{Short, "9c 7d 8d As Qs", "Ac 6s", "Tc Ts", Straight, Pair, -1},
		{Short, "9c 7d 8d As Qs", "Tc Ts", "Ac 6s", Pair, Straight, +1},
		{Short, "9s 7s 8s Ac Qs", "As 6s", "Tc Ts", StraightFlush, Flush, -1},
		{Short, "9s 7s 8s Ac Qs", "Tc Ts", "As 6s", Flush, StraightFlush, +1},
		{Omaha, "Td 2c Jd 4c 5c", "As Ah Qh 3s", "Ad Ac 7d 4d", Straight, Pair, -1},
		{Omaha, "Td 2c Jd 4c 5c", "Ad Ac 7d 4d", "As Ah Qh 3s", Pair, Straight, +1},
		{Omaha, "Kc Qh Jc 8d 4s", "Ac Td 3h 6c", "Ah Tc 2c 3c", Straight, Straight, 0},
		{Omaha, "Kc Qh Jc 8d 4s", "Ah Tc 2c 3c", "Ac Td 3h 6c", Straight, Straight, 0},
		{Omaha, "2d 3h 8s 8h 2s", "Kd Ts Td 4h", "Jd 7d 7c 4c", TwoPair, TwoPair, -1},
		{Omaha, "2d 3h 8s 8h 2s", "Jd 7d 7c 4c", "Kd Ts Td 4h", TwoPair, TwoPair, +1},
		{Omaha, "Tc 6c 2s 3s As", "Kd Qs Js 8h", "9h 9d 4h 4d", Flush, Pair, -1},
		{Omaha, "Tc 6c 2s 3s As", "9h 9d 4h 4d", "Kd Qs Js 8h", Pair, Flush, +1},
		{Omaha, "4s 3h 6c 2d Kd", "Kh Qs 5h 2c", "7s 7c 4h 2s", Straight, TwoPair, -1},
		{Omaha, "4s 3h 6c 2d Kd", "7s 7c 4h 2s", "Kh Qs 5h 2c", TwoPair, Straight, +1},
		/*
			{Lowball, "7h 5h 4h 3h 2c", "7h 6h 4h 3h 2c", ""},
			{Lowball, "7h 6h 4h 3h 2c", "", ""},
			{Lowball, "", "", ""},
			{Lowball, "", "", ""},
		*/
	}
	for i, test := range tests {
		board := Must(test.board)
		a := test.typ.RankHand(Must(test.a), board)
		b := test.typ.RankHand(Must(test.b), board)
		af := a.Fixed()
		if af != test.j {
			t.Errorf("test %d %s expected rank %s, got: %s", i, test.typ, test.j, af)
		}
		bf := b.Fixed()
		if bf != test.k {
			t.Errorf("test %d %s expected rank %s, got: %s", i, test.typ, test.k, bf)
		}
		if n := a.HiComp(b); n != test.exp {
			t.Errorf("test %d %s compare expected %d, got: %d", i, test.typ, test.exp, n)
		}
	}
}

func TestNumberedStreets(t *testing.T) {
	exp := []string{"Ante", "1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th", "10th", "11th", "101st", "102nd", "River"}
	streets := NumberedStreets(0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 90, 1, 1)
	v := make([]string, len(streets))
	for i := 0; i < len(streets); i++ {
		v[i] = streets[i].Name
	}
	if !reflect.DeepEqual(v, exp) {
		t.Errorf("expected items to be equal:\n%v\n%v", exp, v)
	}
}

func TestTypeUnmarshal(t *testing.T) {
	tests := []struct {
		s   string
		exp Type
	}{
		{"HOLDEM", Holdem},
		{"omaha", Omaha},
		{"studHiLo", StudHiLo},
		{"razz", Razz},
		{"BaDUGI", Badugi},
		{"fusIon", Fusion},
	}
	for i, test := range tests {
		typ := Type(^uint16(0))
		if err := typ.UnmarshalText([]byte(test.s)); err != nil {
			t.Fatalf("test %d expected no error, got: %v", i, err)
		}
		if typ != test.exp {
			t.Errorf("test %d expected %d, got: %d", i, test.exp, typ)
		}
	}
}
