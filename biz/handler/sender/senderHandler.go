package sender

import (
	"context"
	"jx-hook/biz/handler"
	"jx-hook/biz/models"
	"jx-hook/biz/models/senderConfig"
	"jx-hook/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Save(ctx context.Context, c *app.RequestContext) {
	var saveVo senderConfig.SenderSaveVO
	err := c.BindAndValidate(&saveVo)
	if err != nil {
		handler.ReturnErr(c, models.InvalidParamCode, models.ErrInvalidParam, err)
		return
	}
	config := saveVo.ToConfig()
	utils.Cache(utils.SenderPrefix+config.ID, config, -1)
	hlog.Info("Succeed save config ", config.ID, config.Name)
	handler.ReturnSuccess(c, consts.StatusCreated, "", config)
}

func Query(ctx context.Context, c *app.RequestContext) {
	config, err := getConfig(c)
	if err != nil {
		return
	}
	handler.ReturnSuccess(c, consts.StatusOK, "", config)
}

func Del(ctx context.Context, c *app.RequestContext) {
	var idOpt models.IDOpt
	err := c.BindAndValidate(&idOpt)
	if err != nil {
		handler.ReturnErr(c, models.InvalidParamCode, models.ErrInvalidParam, err)
		return
	}
	utils.RemoveCache(utils.SenderPrefix + idOpt.ID)
	hlog.Info("Succeed remove alert ", idOpt.ID)
	handler.ReturnSuccess(c, consts.StatusOK, "", nil)
}

func Enable(ctx context.Context, c *app.RequestContext) {
	setEnable(true, c)
}

func Disable(ctx context.Context, c *app.RequestContext) {
	setEnable(false, c)
}

func setEnable(enable bool, c *app.RequestContext) {
	config, err := getConfig(c)
	if err != nil {
		hlog.Warn("Failed to set config enable : ", enable, err)
		return
	}
	config.Enable = &enable
	utils.Cache(utils.SenderPrefix+config.ID, config, -1)
	hlog.Info("Succeed set config ", config.ID, config.Name, " status : ", enable)
	handler.ReturnSuccess(c, consts.StatusOK, "", config)
}

func getConfig(c *app.RequestContext) (senderConfig.SenderConfig, error) {
	var idOpt models.IDOpt
	var config senderConfig.SenderConfig
	err := c.BindAndValidate(&idOpt)
	if err != nil {
		handler.ReturnErr(c, models.InvalidParamCode, models.ErrInvalidParam, err)
		return config, nil
	}

	err = utils.GetCache(utils.SenderPrefix+idOpt.ID, &config)
	if err != nil {
		handler.ReturnErr(c, models.DataNotExistCode, models.ErrDataNotExist, err)
		return config, nil
	}
	return config, nil
}
