package shorturl

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/cache/redis"
	"goapi/service/shorturl"
	"time"
)

type GenerateShortUrlRequest struct {
	LongUrl string `json:"long_url"`
	AddUser string `json:"add_user"`
	Type    int    `json:"type"`
}

type GenerateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

type GetLongUrlResponse struct {
	LongUrl string `json:"long_url"`
}

func GenerateShortUrl(ctx *gin.Context) {

}

func GetLongUrl(ctx *gin.Context) {

}

func RedirectToLongUrl(ctx *gin.Context) {

	c := context.Background()

	redis.RH.Self.Set(c, "hello", "hello redis gin", 10*time.Second).Result()

	name, err := redis.RH.Self.Get(c, "hello").Result()
	if err != nil {

	}

	fmt.Println(name)

	shortStr := ctx.Param("shortStr")

	longUrl := shorturl.GetLongUrl(shortStr)

	ctx.Redirect(302, longUrl)
}
