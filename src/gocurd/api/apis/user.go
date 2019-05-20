package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model "gocurd/api/models"
	"net/http"
	"strconv"
)

//列表数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	result, err := user.Users()

	fmt.Println(result)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data":   result,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	id, err := user.Insert()


	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
}

func GetUsers(c *gin.Context) {
	//
	username := c.Request.FormValue("username")

	c.JSON(http.StatusOK, gin.H{
		"message" : username,
	})
}
