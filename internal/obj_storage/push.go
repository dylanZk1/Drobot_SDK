package obj_storage

import (
	ws "drobot_sdk/internal/interaction"
	"drobot_sdk/drobot_sdk_callback"

	"drobot_sdk/pkg/constant"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/server_api_params"
)

type Push struct {
	p           *ws.PostApi
	platformID  int32
	loginUserID string
}

func NewPush(p *ws.PostApi, platformID int32, loginUserID string) *Push {
	return &Push{p: p, platformID: platformID, loginUserID: loginUserID}
}

func (c *Push) UpdateFcmToken(callback drobot_sdk_callback.Base, fcmToken, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "UpdateFcmToken args: ", fcmToken)
		c.fmcUpdateToken(callback, fcmToken, operationID)
		callback.OnSuccess(sdk_params_callback.UpdateFcmTokenCallback)
		log.NewInfo(operationID, "UpdateFcmToken callback: ", sdk_params_callback.UpdateFcmTokenCallback)
	}()

}

func (c *Push) fmcUpdateToken(callback drobot_sdk_callback.Base, fcmToken, operationID string) {
	apiReq := server_api_params.FcmUpdateTokenReq{}
	apiReq.OperationID = operationID
	apiReq.Platform = int(c.platformID)
	apiReq.FcmToken = fcmToken
	c.p.PostFatalCallback(callback, constant.FcmUpdateTokenRouter, apiReq, nil, apiReq.OperationID)
}
func (c *Push) SetAppBadge(callback drobot_sdk_callback.Base, appUnreadCount int32, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, "SetAppBadge args: ", appUnreadCount)
		c.setAppBadge(callback, appUnreadCount, operationID)
		callback.OnSuccess(sdk_params_callback.SetAppBadgeCallback)
		log.NewInfo(operationID, "SetAppBadge callback: ", sdk_params_callback.SetAppBadgeCallback)
	}()
}
func (c *Push) setAppBadge(callback drobot_sdk_callback.Base, appUnreadCount int32, operationID string) {
	apiReq := server_api_params.SetAppBadgeReq{}
	apiReq.OperationID = operationID
	apiReq.FromUserID = c.loginUserID
	apiReq.AppUnreadCount = appUnreadCount
	c.p.PostFatalCallback(callback, constant.SetAppBadgeRouter, apiReq, nil, apiReq.OperationID)
}
