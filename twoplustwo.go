//go:build forcefat || (!portable && !embedded)

package cardrank

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func init() {
	if twoplustwo01Dat != nil {
		twoPlusTwo = NewTwoPlusTwo()
	}
}

// NewTwoPlusTwo creates a new two plus hand rank eval func.
//
// The TwoPlusTwo eval is a version of the 2+2 poker forum hand rank evaluator.
// Uses the embedded twoplustwo*.dat files to provide extremely fast 7 card
// hand lookup. Uses Cactus Kev values.
//
// The lookup table is contained in the 'twoplustwo*.dat' files, and were
// broken up from a single file to get around GitHub's size limitations. Files
// were generated with 'gen.go', which is a pure-Go implementation of the code
// generator available at: https://github.com/tangentforks/TwoPlusTwoHandEvaluator
//
// When recombined, the lookup table has the same hash as the original table
// generated using the C code.
func NewTwoPlusTwo() HandRankFunc {
	const total, chunk, last = 32487834, 2621440, 1030554
	tbl, pos := make([]uint32, total), 0
	for i, buf := range [][]byte{
		twoplustwo00Dat,
		twoplustwo01Dat,
		twoplustwo02Dat,
		twoplustwo03Dat,
		twoplustwo04Dat,
		twoplustwo05Dat,
		twoplustwo06Dat,
		twoplustwo07Dat,
		twoplustwo08Dat,
		twoplustwo09Dat,
		twoplustwo10Dat,
		twoplustwo11Dat,
		twoplustwo12Dat,
	} {
		n, exp := len(buf), chunk
		if i == 12 {
			exp = last
		}
		if n%4 != 0 || n/4 != exp {
			panic(fmt.Sprintf("twoplustwo%02d.dat is bad: expected %d uint32, has: %d", i, exp, n/4))
		}
		if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, tbl[pos:pos+n/4]); err != nil {
			panic(fmt.Sprintf("twoplustwo%02d.dat is bad: %v", i, err))
		}
		pos += n / 4
	}
	if pos != total {
		panic("short read twoplustwo*.dat")
	}
	// build card map
	m := make(map[Card]uint32, 52)
	for i, r := uint32(0), Two; r <= Ace; r++ {
		for _, s := range []Suit{Spade, Heart, Club, Diamond} {
			m[New(r, s)] = i + 1
			i++
		}
	}
	ranks := [10]uint32{
		uint32(Invalid),
		uint32(HighCard),
		uint32(Pair),
		uint32(TwoPair),
		uint32(ThreeOfAKind),
		uint32(Straight),
		uint32(Flush),
		uint32(FullHouse),
		uint32(FourOfAKind),
		uint32(StraightFlush),
	}
	return func(v []Card) HandRank {
		i := uint32(53)
		for _, c := range v {
			i = tbl[i+m[c]]
		}
		if len(v) < 7 {
			i = tbl[i]
		}
		return HandRank(ranks[i>>12] - i&0xfff + 1)
	}
}
