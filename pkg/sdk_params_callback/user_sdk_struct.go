package sdk_params_callback

import (
	"drobot_sdk/pkg/constant"
	"drobot_sdk/pkg/db/model_struct"
	"drobot_sdk/pkg/server_api_params"
)

//other user
type GetUsersInfoParam []string
type GetUsersInfoCallback []server_api_params.FullUserInfo

//type GetSelfUserInfoParam string
type GetSelfUserInfoCallback *model_struct.LocalUser

type SetSelfUserInfoParam server_api_params.ApiUserInfo

const SetSelfUserInfoCallback = constant.SuccessCallbackDefault
