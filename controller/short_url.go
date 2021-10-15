package controller

import (
	"github.com/gin-gonic/gin"
	"goapi/entity/dto"
	"goapi/entity/vo"
	"goapi/service/shorturl"
)

//生成短链接并保存
func GenerateShortUrl(ctx *gin.Context) {
	var g dto.GenerateShortUrlDto
	if err := ctx.Bind(&g); err != nil {
		SendError(ctx, err, nil)
		return
	}

	//生成短链接
	shortUrl, err := shorturl.GenerateShortUrl(g.LongUrl)
	if err != nil {
		SendError(ctx, err, nil)
		return
	}

	SendSuccess(ctx, vo.GenerateShortUrlVo{ShortUrl: shortUrl})
	return
}

//短链接重定向值原始长连接
func RedirectToLongUrl(ctx *gin.Context) {
	shortStr := ctx.Param("shortStr")

	longUrl, _ := shorturl.GetLongUrl(shortStr)

	ctx.Redirect(302, longUrl)
}
