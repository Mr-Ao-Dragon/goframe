package randomseed

import (
	"crypto/rand"
	"github.com/spf13/cast"
	"math"
	"math/big"
)

func Small() *big.Int {
	var SmallSeed, _ = rand.Int(rand.Reader, big.NewInt(64)) // 真随机出一个数值并返回。
	return SmallSeed
}
func Big() int64 {
	var (
		SmallSeed, _   = rand.Int(rand.Reader, big.NewInt(64))
		Salt, _        = rand.Int(rand.Reader, big.NewInt(64))
		BigSeed        int64
		SmallSeedFloat float64
		SaltFloat      float64
		BigSeedPow     float64
	)
	SmallSeedFloat = cast.ToFloat64(SmallSeed)
	SaltFloat = cast.ToFloat64(Salt)
	BigSeedPow = math.Pow(SaltFloat, SmallSeedFloat)
	BigSeed = cast.ToInt64(BigSeedPow)
	return BigSeed
}
