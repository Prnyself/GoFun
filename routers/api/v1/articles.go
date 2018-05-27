package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 删除文章
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 200 {string} json "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Router /api/v1/articles/{id} [delete]
func Get(c *gin.Context) {
	c.String(200, "ddd")
}
func DeleteArticle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 12,
		"msg":  "ds",
		"data": make(map[string]string),
	})
}
