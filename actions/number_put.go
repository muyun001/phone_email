package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phone_email/services"
	"strconv"
)

// NumberPut 保存手机号等信息
// 一次接收一条数据
func NumberPut(c *gin.Context) {
	callId := c.Param("call_id")
	number := c.Param("number")
	emailType, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "传入的Request格式错误",
		})
		return
	}

	if err := services.SaveNumber(callId, number, emailType); err != nil {
		if err.Error() == "Duplicate entry" {
			c.JSON(http.StatusOK, gin.H{
				"message": "重复发送数据",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "保存数据失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "保存数据成功",
	})
}
