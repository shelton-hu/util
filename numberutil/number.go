package numberutil

import (
	"errors"
	"math/big"
)

var (
	ErrMinIsGreaterThanMax = errors.New("最小值大于最大值")
)

func Uint32Slice2Int32Slice(ss []uint32) []int32 {
	rs := make([]int32, len(ss))
	for i, s := range ss {
		rs[i] = int32(s)
	}
	return rs
}

func ValidateMinAndMax(min int, max int) error {
	if min != 0 && max != 0 {
		if min > max {
			return ErrMinIsGreaterThanMax
		}
	}
	return nil
}

func Uint642String(u uint64) string {
	return new(big.Int).SetUint64(u).String()
}

func BigIntAdd1(b *big.Int) *big.Int {
	return new(big.Int).Add(b, big.NewInt(1))
}

func BigIntSub1(b *big.Int) *big.Int {
	return new(big.Int).Sub(b, big.NewInt(1))
}
