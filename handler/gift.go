package handler

import (
	"lottery/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLotteryGifts(ctx *gin.Context) {
	gifts := database.GetAllGiftsV1()
	if len(gifts) == 0 {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		// 去掉敏感信息 不需要返回给前端商品库存的数量
		for _, gift := range gifts {
			gift.Count = 0
		}
		ctx.JSON(http.StatusOK, gifts)
	}
}
