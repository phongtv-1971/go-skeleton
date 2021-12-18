package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func CustomLog(param gin.LogFormatterParams) string {
	reqBody, _ := ioutil.ReadAll(param.Request.Body)
	message := fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\" \n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
	if string(reqBody) != "" {
		message += fmt.Sprintf("Request body: %s\n", string(reqBody))
	}
	return message
}
