package full

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/common"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/utils"
)

func (u *Full) GetUsersInfo(callback drobot_sdk_callback.Base, userIDList string, operationID string) {
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", userIDList)
		var unmarshalParam sdk_params_callback.GetUsersInfoParam
		common.JsonUnmarshalAndArgsValidate(userIDList, &unmarshalParam, callback, operationID)
		result := u.getUsersInfo(callback, unmarshalParam, operationID)
		callback.OnSuccess(utils.StructToJsonStringDefault(result))
		log.NewInfo(operationID, fName, "callback: ", utils.StructToJsonStringDefault(result))
	}()
}
