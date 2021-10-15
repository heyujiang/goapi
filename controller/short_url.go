package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapi/entity/dto"
	"goapi/entity/vo"
	shorturl2 "goapi/model/shorturl"
	"goapi/service/shorturl"
)

type GetLongUrlResponse struct {
	LongUrl string `json:"long_url"`
}

//生成短链接并保存
func GenerateShortUrl(ctx *gin.Context) {
	var g dto.GenerateShortUrlDto
	if err := ctx.Bind(&g); err != nil {
		SendError(ctx, err, nil)
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
		SendError(ctx, err, nil)
		return
	}

	baseDemain := viper.GetString("base_domain")
	SendSuccess(ctx, vo.GenerateShortUrlVo{ShortUrl: baseDemain + "/" + shortUrl})
	return
}

//短链接重定向值原始长连接
func RedirectToLongUrl(ctx *gin.Context) {
	shortStr := ctx.Param("shortStr")

	longUrl, _ := shorturl.GetLongUrl(shortStr)

	ctx.Redirect(302, longUrl)
}
