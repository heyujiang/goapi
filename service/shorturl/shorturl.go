package shorturl

import (
	"context"
	"github.com/spf13/viper"
	"goapi/cache/redis"
	"goapi/model/shorturl"
	"goapi/util"
	"log"
	"time"
)

func GenerateShortUrl(longUrl string) (string, error) {
	strs := util.LongToShortUrl(longUrl)

	shortUrl := strs[0]

	model := shorturl.ShorturlModel{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}

	if err := model.Create(); err != nil {
		return "", err
	}

	baseDemain := viper.GetString("base_domain")

	//保存缓存
	c := context.Background()
	_, _ = redis.RH.Self.Set(c, shortUrl, longUrl, 86400*time.Second).Result()

	return baseDemain + "/" + shortUrl, nil
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
