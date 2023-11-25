package common

import (
	"context"
	"net/http"

	"gen-templates/middleware/biz_ctx"
	"gen-templates/middleware/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ResponseOK(c *gin.Context, ctx context.Context, data interface{}) {
	ctx = biz_ctx.AppendFieldsToContext(ctx, "uri", c.Request.RequestURI)
	log.Get().InfoCtx(ctx, "[API Response] Succeed.")
	c.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "succeed",
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context, ctx context.Context, data interface{}) {
	ctx = biz_ctx.AppendFieldsToContext(ctx, "uri", c.Request.RequestURI)
	log.Get().InfoCtx(ctx, "[API Response] Created.")
	c.JSON(http.StatusCreated, gin.H{
		"code":    20100,
		"message": "create succeed",
		"data":    data,
	})
}

func ResponseError(c *gin.Context, ctx context.Context, err error) {
	ctx = biz_ctx.AppendFieldsToContext(ctx, "uri", c.Request.RequestURI)
	log.Get().ErrorCtx(ctx, "[API Response] Something error",
		zap.Error(err),
	)
	c.JSON(http.StatusOK, gin.H{
		"code":    500,
		"message": "internal error",
		"error":   err,
	})
}
