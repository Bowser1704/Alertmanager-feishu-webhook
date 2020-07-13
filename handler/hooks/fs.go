package hooks

import (
	"github.com/Bowser1704/Alertmanager-feishu-webhook/handler"
	"github.com/Bowser1704/Alertmanager-feishu-webhook/pkg/errno"
	"github.com/Bowser1704/Alertmanager-feishu-webhook/pkg/feishu"
	"github.com/Bowser1704/Alertmanager-feishu-webhook/pkg/model"
	"github.com/gin-gonic/gin"
)

// Feishu 飞书
func Feishu(c *gin.Context) {
	fstoken := c.Query("fs")
	fs := feishu.FS{
		Token: fstoken,
	}

	var msg model.WebhookMessage
	if err := c.BindJSON(&msg); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// alertmanager msg to feishu msg
	fs.Init(&msg)

	// firing
	fs.Firing()

	handler.SendResponse(c, nil, "OK")
}
