package shorturl

import (
	"context"
	"github.com/spf13/viper"
	"goapi/cache/redis"
	"goapi/model/shorturl"
	"goapi/util"
	"log"
	"strconv"
	"time"
)

var chars []byte = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5',
	'6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

func GenerateShortUrl(longUrl string) []string {
	has := util.Md5(longUrl)

	log.Println(has)

	strs := make([]string, 4)
	for i := 0; i < 4; i++ {
		sTempString := has[i*8 : (i+1)*8]
		log.Println(sTempString)

		s, _ := strconv.ParseUint(sTempString, 16, 32)
		log.Println(s)

		ii := 0x3FFFFFFF & s

		strBytes := make([]byte, 6)
		for j := 0; j < 6; j++ {
			index := 0x0000003D & ii

			strBytes[j] = chars[index]

			log.Println(index)
			ii = ii >> 5
		}

		log.Println("==========")
		strs[i] = string(strBytes)
	}

	return strs
}

//获得长连接
func GetLongUrl(shortUrl string) (string, error) {
	if shortUrl == "" {
		return viper.GetString("default_long_url"), nil
	}

	//读取缓存的长连接
	c := context.Background()

	longUrl, _ := redis.RH.Self.Get(c, shortUrl).Result()
	log.Printf("redis cache shortUrl : %s , longUrl : %s ;", shortUrl, longUrl)

	//从数据库读取
	if longUrl == "" {
		s, err := shorturl.GetInfoByShortUrl(shortUrl)

		log.Println(s)

		if err != nil {
			return viper.GetString("default_long_url"), err
		}

		if s.LongUrl == "" {
			return viper.GetString("default_long_url"), nil
		}

		longUrl = s.LongUrl

		//保存缓存
		_, _ = redis.RH.Self.Set(c, shortUrl, longUrl, 86400*time.Second).Result()
	}

	return longUrl, nil
}
