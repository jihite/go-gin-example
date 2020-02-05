package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jihite/go-gin-example/models"
	"github.com/jihite/go-gin-example/pkg/e"
	"github.com/jihite/go-gin-example/pkg/setting"
	"github.com/jihite/go-gin-example/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}

func ExportTag(c *gin.Context) {

}

func ImportTag(c *gin.Context) {

}
