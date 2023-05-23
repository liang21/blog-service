package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lgl/blog-service/global"
	"github.com/lgl/blog-service/pkg/app"
	"github.com/lgl/blog-service/pkg/errcode"
	"github.com/lgl/blog-service/pkg/logger"
	"time"
)

type AccessWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter
		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
