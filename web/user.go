package web

import (
	"awesomeProject/conf"
	"awesomeProject/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	db     = database.GetGlobalDB()
	logger = conf.GetGlobalLogger()
)

func LoggingMiddleware(c *gin.Context) { // 日志记录中间件，创建一个中间件来统一处理日志记录，这样就不用在每个处理函数里重复编写日志记录代码了
	c.Next() //会使得当前中间件暂停执行，等待后续处理函数完成请求处理，之后再继续执行当前中间件剩余的代码，确保能记录到最终的响应状态码
	logger.Debugf("Method:%s URL:%s Code:%d", c.Request.Method, c.Request.URL.String(), c.Writer.Status())
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!!!!!!!!!",
	})
}

func QueryAllUsers(c *gin.Context) { //指定字段查询所有用户
	users := database.SelectFieldsQueryAllUser(db)
	c.JSON(http.StatusOK, users)
}
