package icrypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"log"
	"math"

	"github.com/flower/script-go/iutils"
	"github.com/google/uuid"
)

func TestICryptoBase64Funcs(src []byte) {
	// base64 & url encode
	// https://www.cnblogs.com/zhoading/p/15666892.html

	log.Println(
		base64.StdEncoding.EncodeToString(
			src,
		),
	)

}

func TestICryptoHex() {
	const (
		MD5  = "MD5"
		SHA1 = "SHA1"
	)

	type HexInp struct {
		Inp  string
		Type string
		Salt string
		Len  int
	}

	var Hex func(inp *HexInp) string = func(inp *HexInp) string {
		var h hash.Hash

		switch (*inp).Type {
		case MD5:
			h = md5.New()
		case SHA1:
			h = sha1.New()

		default:
			return ""
		}

		temp := (*inp).Inp + (*inp).Salt
		h.Write([]byte(temp))
		oup := hex.EncodeToString(h.Sum(nil))

		if (*inp).Len <= 0 {
			return oup
		} else {
			return oup[:(*inp).Len]
		}
	}

	log.Println(
		Hex(&HexInp{
			Inp:  "",
			Type: MD5,
			Salt: "",
			Len:  0,
		}))

}

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
// type UUID [16]byte
func TestICryptoUuid() {
	var (
		flag bool
	)

	// uuidBs, err := uuid.Parse("dddf999b-193e-4015-9463-da3dab1f1a6a")
	// uuid from mysql
	uuidBs, err := uuid.Parse("051AF9FDE542439482F0494D34C75600")
	flag = nil == err

	uuidStr := uuidBs.String()

	/*
		uuidBs  len 16
		uuidStr len 36 (32 + 4) (8-4-4-4-12)

		16进制，16进1，10 -> 16 * 1 + 16 * 0 -> 16
		0, 1, 2, ... 9, a, b, c, d, e -> 2进制 -> xxxx -> 4 bit -> 2*2*2*2 -> 2^4 -> 16
		4比特 可以表示 1个16进制数

		32个16进制数 <-> 32 * 4 bit <-> 128 bit <-> 128 / 8 byte <-> 16 byte

		8进制
			0o10 <-> 8
			0o7  <-> 7
	*/

	iutils.Log(
		flag,
		uuidBs,
		len(uuidBs),
		uuidStr,
		len(uuidStr),
	)

	/*
		1 byte <-> 8 bit <-> 2^8 <-> 256 <-> [0, 255], [-128, 127], ...

		位运算符
			左移 << xxxxyyyy << 4 -> yyyy0000
			右移 >> xxxxyyyy >> 4 -> 0000xxxx

			按位与 & xxxxyyyy&0x0f -> 0000yyyy

	*/
	// 1111 1111 -> 16进制 ff -> 0xff
	// 1111 0000 -> 16进制 f0 -> 0xf0

	iutils.Log(
		byte(255),
		byte(0xff),
		byte(0xf0)>>4,
		byte(0xf0)&0x0f,
		math.MaxInt8,
		math.MinInt8,
	)

	var (
		// []byte & [2]byte 不同类型
		v [2]byte
		// v [2]byte = [2]byte{}
		// v [2]byte = [...]byte{byte(0), byte(0)} // [2]byte{byte(0), byte(0)}

		// v []byte
		// v []byte = []byte{}
		// v []byte = []byte{byte(0), byte(0)}

		bs []byte = []byte{byte(0xff)}
	)

	log.Printf("%T, %v\n", v, v)
	log.Printf("%T, %v\n", v[:], v[:])

	log.Printf("%T, %v\n", []byte{}, len([]byte{}))
	log.Printf("%T, %v\n", []byte{}[:], len([]byte{}[:]))

	log.Printf("%T, %v\n", []byte{byte(0), byte(0)}, []byte{byte(0), byte(0)})
	log.Printf("%T, %v\n", []byte{byte(0), byte(0)}[:], []byte{byte(0), byte(0)}[:])

	// ? [:] -> invalid operation: cannot slice ([2]byte literal) (value of type [2]byte) (value not addressable)
	log.Printf("%T, %v\n", [2]byte{}, [2]byte{})
	log.Printf("%T, %v\n", [...]byte{byte(0), byte(0)}, [...]byte{byte(0), byte(0)})

	// [16]byte <-> string
	// cannot convert uuidBs (variable of type uuid.UUID) to string
	// log.Println(string(uuidBs))

	// []byte <-> string
	// [16]byte <-> []byte <-> string
	log.Println(uuidBs[:])
	log.Println(string(uuidBs[:]))
	log.Println(len(string(uuidBs[:])))

	log.Println([]byte("汉字"))
	log.Println(string([]byte("汉字")))
	// len -> 字节数 而不是 字符数
	log.Println(len("汉字"))
	log.Println(len("english"))

	hex.Encode(v[:], bs)
	// {1111 1111} -> {0000 1111, 0000 1111} -> {15, 15} -> {'f', 'f'}
	iutils.Log(
		[]byte("ff"),
		v,
		string(v[:]),
		hex.EncodeToString(bs),
	)

}
