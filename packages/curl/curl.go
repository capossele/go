package curl

import (
    "github.com/iotadevelopment/go/packages/ternary"
    "math"
)

const (
    HASH_LENGTH         = 243
    STATE_LENGTH        = ternary.NUMBER_OF_TRITS_IN_A_TRYTE * HASH_LENGTH
)

var (
    TRUTH_TABLE = ternary.Trits{1, 0, -1, 2, 1, -1, 0, 2, -1, 1, 0}
)


type Hash interface {
    Initialize()
    InitializeCurl(trits *[]int8, length int, rounds int)
    Reset()
    Absorb(trits *[]int8, offset int, length int)
    Squeeze(resp []int8, offset int, length int) []int
}


type Curl struct {
    Hash
    state  ternary.Trits
    hashLength int
    rounds int
}

func NewCurl(hashLength int, rounds int) *Curl {
    this := &Curl{
        hashLength: hashLength,
        rounds: rounds,
    }

    this.Reset()

    return this
}

func (curl *Curl) Initialize() {
    curl.InitializeCurl(nil, 0, curl.rounds)
}

func (curl *Curl) InitializeCurl(trinary ternary.Trits, length int, rounds int) {
    curl.rounds = rounds
    if trinary != nil {
        curl.state = trinary
    } else {
        curl.state = make(ternary.Trits, STATE_LENGTH)
    }
}

func (curl *Curl) Reset() {
    curl.InitializeCurl(nil, 0, curl.rounds)
}

func (curl *Curl) Absorb(trinary ternary.Trits, offset int, length int) {
    for {
        limit := int(math.Min(HASH_LENGTH, float64(length)))
        copy(curl.state, trinary[offset:offset+limit])
        curl.Transform()
        offset += HASH_LENGTH
        length -= HASH_LENGTH
        if length <= 0 {
            break
        }
    }
}

func (curl *Curl) Squeeze(resp ternary.Trits, offset int, length int) ternary.Trits {
    for {
        limit := int(math.Min(HASH_LENGTH, float64(length)))
        copy(resp[offset:offset+limit], curl.state)
        curl.Transform()
        offset += HASH_LENGTH
        length -= HASH_LENGTH
        if length <= 0 {
            break
        }
    }
    return resp
}

func (curl *Curl) Transform() {
    var index = 0
    for round := 0; round < curl.rounds; round++ {
        stateCopy := make(ternary.Trits, STATE_LENGTH)
        copy(stateCopy, curl.state)
        for i := 0; i < STATE_LENGTH; i++ {
            incr := 364
            if index >= 365 {
                incr = -365
            }
            index2 := index + incr
            curl.state[i] = TRUTH_TABLE[stateCopy[index]+(stateCopy[index2]<<2)+5]
            index = index2
        }
    }
}