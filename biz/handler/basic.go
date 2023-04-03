package handler

import (
	"context"
	"jx-hook/biz/models/vos"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Auth(ctx context.Context, c *app.RequestContext) {
	c.Status(consts.StatusOK)
}

func ReturnSuccess(c *app.RequestContext, code int, msg string, data any) {
	c.JSON(code, &vos.CommonResp{
		Code:    vos.SucceedCode,
		Succeed: true,
		Msg:     msg,
		Data:    data,
	})
}

func ReturnErr(c *app.RequestContext, code int, myErr, err error) {
	hlog.Warn("Return err resp from uri ", c.URI(), myErr, err)
	c.JSON(consts.StatusBadRequest, &vos.CommonResp{
		Code:    code,
		Succeed: false,
		Msg:     myErr.Error(),
	})
}
