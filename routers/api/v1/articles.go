package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaoheigou/GoFun/pkg/db"
)

// @Summary 删除文章
// @Produce  json
// @Param id param int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 200 {string} json "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Router /api/v1/articles/{id} [delete]
func Get(c *gin.Context) {
	user := new(db.Users)
	has, err := dd.Get(user) // for single primary key

	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"has":  has,
		"err":  err,
	})
}
func DeleteArticle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 12,
		"msg":  "ds",
		"data": make(map[string]string),
	})
}
