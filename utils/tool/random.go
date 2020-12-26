package tool

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"math"
	mRand "math/rand"
	"time"
)

// GenRandStr 生成随机字符串
func GenRandStr(length int) (string, error) {
	b := make([]byte, length/2)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", errors.New("could not successfully read from the system CSPRNG")
	}
	return hex.EncodeToString(b), nil
}

func init() {
	mRand.Seed(time.Now().UnixNano())
}

// GetRandom 在指定范围内生成随机数, min ≤ rnd < max
func GetRandom(min, max int64) int64 {
	return min + mRand.Int63n(max-min)
}

// GetRandomN 在指定范围内生成N个不重复的随机数 min ≤ rnd ≤ max(如果数据超过10000, 不建议使用)
func GetRandomN(min, max, n int64) []int64 {
	var length int64
	if diff := max - min; diff < n {
		length = diff
	} else {
		length = n
	}

	list := SliceRandList(int(min), int(max))

	var result = make([]int64, 0, length)
	for _, n := range list[:length] {
		result = append(result, int64(n))
	}

	return result
}

// RandomHit 随机命中 (判断百分比概率)
// hit 是指命中值, 值范围是 0 ~ 1
// double 是倍率, 即数值放大多少倍后再判断
func RandomHit(hit float64, double ...int64) bool {
	var dbe float64
	if len(double) > 0 {
		dbe = float64(double[0])
	} else {
		dbe = 100
	}

	// 计算概率
	rd := GetRandom(0, int64(dbe))

	if rd <= int64(math.Floor(hit*dbe)) {
		return true
	}

	return false
}
