package alert

import (
	"context"
	"jx-hook/biz/handler"
	"jx-hook/biz/models"
	"jx-hook/biz/models/alertConfig"
	"jx-hook/biz/models/senderConfig"
	"jx-hook/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Doing the alert, each alert may connect multi sender
func Alert(ctx context.Context, c *app.RequestContext) {
	config, err := getConfig(c)
	if err != nil {
		return
	}
	if !*config.Enable {
		handler.ReturnErr(c, models.DateNotActiveCode, models.ErrDataNotActive, nil)
		return
	}

	body, err := c.Body()
	if err != nil {
		handler.ReturnErr(c, models.InValidBodyCode, models.ErrInvalidBody, err)
		return
	}
	for _, senderID := range config.SenderIds {
		senderConfig, err := getSenderConfig(senderID)
		if err != nil {
			hlog.Warn("Failed to send due to no sender found, sender id ", senderID)
			continue
		}
		if *senderConfig.Enable {
			msg, err := utils.ResolveTemplate(senderConfig.TemplateMsg, body)
			if err != nil {
				hlog.Warn("Failed to send due to template resolve failed, sender id ", senderID, err)
				continue
			}
			err = utils.SendMsg(senderConfig, msg)
			if err != nil {
				hlog.Warn("Failed to send due to wechat failed, sender id ", senderID, err)
			}
		}
	}
}

func Save(ctx context.Context, c *app.RequestContext) {
	var saveVo alertConfig.AlertSaveVO
	err := c.BindAndValidate(&saveVo)
	if err != nil {
		handler.ReturnErr(c, models.InvalidParamCode, models.ErrInvalidParam, err)
		return
	}
	config := saveVo.ToConfig()
	utils.Cache(utils.AlertPrefix+config.ID, config, -1)
	hlog.Info("Succeed save config ", config.ID)
	handler.ReturnSuccess(c, consts.StatusCreated, "", nil)
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
	utils.RemoveCache(utils.AlertPrefix + idOpt.ID)
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
		return
	}
	config.Enable = &enable
	utils.Cache(utils.AlertPrefix+config.ID, config, -1)
	hlog.Info("Succeed set config ", config.ID, " status : ", enable)
	handler.ReturnSuccess(c, consts.StatusOK, "", config)
}

func getConfig(c *app.RequestContext) (alertConfig.AlertConfig, error) {
	var idOpt models.IDOpt
	var config alertConfig.AlertConfig
	err := c.BindAndValidate(&idOpt)
	if err != nil {
		handler.ReturnErr(c, models.InvalidParamCode, models.ErrInvalidParam, err)
		return config, err
	}
	err = utils.GetCache(utils.AlertPrefix+idOpt.ID, &config)
	if err != nil {
		handler.ReturnErr(c, models.DataNotExistCode, models.ErrDataNotExist, err)
		return config, err
	}
	return config, nil
}

func getSenderConfig(id string) (senderConfig.SenderConfig, error) {
	var senderConfig senderConfig.SenderConfig
	err := utils.GetCache(utils.SenderPrefix+id, &senderConfig)
	if err != nil {
		return senderConfig, err
	}
	return senderConfig, nil
}
