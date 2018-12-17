package ternary

// a Binary Coded Trit encodes a Trit in 2 bits with -1 => 00, 0 => 01 and 1 => 10
type BCTrit struct {
    Lo uint
    Hi uint
}

// a Binary Coded Trinary consists out of many Binary Coded Trits
type BCTrinary struct {
    Lo []uint
    Hi []uint
}
