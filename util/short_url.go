package util

import (
	"log"
	"strconv"
)

var chars []byte = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5',
	'6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

func LongToShortUrl(longUrl string) []string {
	has := Md5(longUrl)
	log.Println(has)
	strs := make([]string, 4)
	for i := 0; i < 4; i++ {
		sTempString := has[i*8 : (i+1)*8]
		s, _ := strconv.ParseUint(sTempString, 16, 32)
		ii := 0x3FFFFFFF & s
		strBytes := make([]byte, 6)
		for j := 0; j < 6; j++ {
			index := 0x0000003D & ii
			strBytes[j] = chars[index]
			ii = ii >> 5
		}
		strs[i] = string(strBytes)
	}

	return strs
}
