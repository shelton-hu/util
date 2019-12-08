package idutil

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
	"time"

	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"

	"github.com/shelton-hu/util/stringutil"
	"github.com/shelton-hu/util/timeutil"
)

var sf *sonyflake.Sonyflake
var randsrc = mathrand.NewSource(time.Now().Unix())

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
}

func GetIntId() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}

// format likes: B6BZVN3mOPvx
func GetUuid(prefix string) string {
	id := GetIntId()
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	i, err := h.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(i)
}

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

// format likes: 300m50zn91nwz5
func GetUuid36(prefix string) string {
	id := GetIntId()
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	i, err := h.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(i)
}

// 生成订单号(订单号格式: 首字母 + 年(后两位) + 月(两位) + 日(两位) + 时(两位) + 分(两位) + 模块号(两位) + 随机数(四位)))
func GetOrderNo(firstLetter string, moduelNo string) string {
	// 1.取年月日时分
	datetimeString := timeutil.Time2String(time.Now(), "0601021504")
	// 2.取流水号
	serialString := fmt.Sprintf("%0*d", 4, mathrand.New(randsrc).Intn(9999))
	return stringutil.StringJoin(firstLetter, datetimeString, moduelNo, serialString)
}

func randString(letters string, n int) string {
	output := make([]byte, n)

	// We will take n bytes, one byte for each character of output.
	randomness := make([]byte, n)

	// read all random
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(letters)
	// fill output
	for pos := range output {
		// get random item
		random := uint8(randomness[pos])

		// random % 64
		randomPos := random % uint8(l)

		// put into output
		output[pos] = letters[randomPos]
	}

	return string(output)
}

func GetSecret() string {
	return randString(Alphabet62, 50)
}

func GetRefreshToken() string {
	return randString(Alphabet62, 50)
}

func GetAttachmentPrefix() string {
	return randString(Alphabet62, 30)
}
