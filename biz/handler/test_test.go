package handler

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

// hertz test
func TestPerformRequest(t *testing.T) {
	h := server.Default()
	h.GET("/ping", PingHandler)
	w := ut.PerformRequest(h.Engine, "GET", "/ping", &ut.Body{Body: bytes.NewBufferString("1"), Len: 1},
		ut.Header{Key: "Connection", Value: "close"})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "{\"message\":\"pong\"}", string(resp.Body()))
}
