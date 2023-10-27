package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wb := &bodyInterceptorWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = wb
		ctx.Next()
		data := wb.body.Bytes()
		code := wb.Status()
		if code == 200 {
			code = 0
		}
		wb.ResponseWriter.WriteHeader(200)
		_, _ = wb.ResponseWriter.Write([]byte(fmt.Sprintf(`{"code":%d,"data":%s}`, code, data)))
		wb.body.Reset()
	}
}

type bodyInterceptorWriter struct {
	body *bytes.Buffer
	gin.ResponseWriter
}

func (w bodyInterceptorWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (w bodyInterceptorWriter) WriteString(s string) (int, error) {
	return w.body.WriteString(s)
}

type LoginReq struct {
	WalletAddress string `json:"wallet_address"`
	SignData      string `json:"sign_data"`
}
