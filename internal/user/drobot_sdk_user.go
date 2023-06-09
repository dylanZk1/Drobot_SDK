package user

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/common"
	"drobot_sdk/pkg/constant"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/utils"
)

//func (u *User) GetUsersInfo(callback drobot_sdk_callback.Base, userIDList string, operationID string) {
//	fName := utils.GetSelfFuncName()
//	go func() {
//		log.NewInfo(operationID, fName, "args: ", userIDList)
//		var unmarshalParam sdk_params_callback.GetUsersInfoParam
//		common.JsonUnmarshalAndArgsValidate(userIDList, &unmarshalParam, callback, operationID)
//		result := u.GetUsersInfoFromSvr(callback, unmarshalParam, operationID)
//		callback.OnSuccess(utils.StructToJsonStringDefault(result))
//		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonStringDefault(result))
//	}()
//}

func (u *User) GetSelfUserInfo(callback drobot_sdk_callback.Base, operationID string) {
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ")
		result := u.getSelfUserInfo(callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(result))
	}()
}

func (u *User) SetSelfInfo(callback drobot_sdk_callback.Base, userInfo string, operationID string) {
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", userInfo)
		var unmarshalParam sdk_params_callback.SetSelfUserInfoParam
		common.JsonUnmarshalAndArgsValidate(userInfo, &unmarshalParam, callback, operationID)
		u.updateSelfUserInfo(callback, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.SetSelfUserInfoCallback))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonString(sdk_params_callback.SetSelfUserInfoCallback))
	}()
}

func (u *User) updateMsgSenderInfo(nickname, faceURL string, operationID string) {
	if nickname != "" {
		err := u.DataBase.UpdateMsgSenderNickname(u.loginUserID, nickname, constant.SingleChatType)
		if err != nil {
			log.Error(operationID, "UpdateMsgSenderNickname failed ", err.Error(), u.loginUserID, nickname, constant.SingleChatType)
		}
	}
	if faceURL != "" {
		err := u.DataBase.UpdateMsgSenderFaceURL(u.loginUserID, faceURL, constant.SingleChatType)
		if err != nil {
			log.Error(operationID, "UpdateMsgSenderFaceURL failed ", err.Error(), u.loginUserID, faceURL, constant.SingleChatType)
		}
	}
}
