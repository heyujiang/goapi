package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapi/handler"
	shorturl2 "goapi/model/shorturl"
	"goapi/pkg/errno"
	"goapi/service/shorturl"
	"log"
)

type GenerateShortUrlRequest struct {
	LongUrl string `json:"long_url"`
}

type GenerateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

type GetLongUrlResponse struct {
	LongUrl string `json:"long_url"`
}

func GenerateShortUrl(ctx *gin.Context) {
	var g GenerateShortUrlRequest
	log.Println(g)
	if err := ctx.Bind(&g); err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	//生成短链接
	strs := shorturl.GenerateShortUrl(g.LongUrl)

	shortUrl := strs[0]

	model := shorturl2.ShorturlModel{
		LongUrl:  g.LongUrl,
		ShortUrl: shortUrl,
	}

	if err := model.Create(); err != nil {
		handler.SendResponse(ctx, err, nil)
	}

	baseDemain := viper.GetString("base_domain")
	handler.SendResponse(ctx, errno.OK, GenerateShortUrlResponse{ShortUrl: baseDemain + "/" + shortUrl})
}

func GetLongUrl(ctx *gin.Context) {

}

func RedirectToLongUrl(ctx *gin.Context) {
	shortStr := ctx.Param("shortStr")

	longUrl, _ := shorturl.GetLongUrl(shortStr)

	ctx.Redirect(302, longUrl)
}