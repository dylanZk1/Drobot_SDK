package common

import (
	"github.com/golang/protobuf/proto"
	api "drobot_sdk/pkg/server_api_params"
	"drobot_sdk/pkg/utils"
)

func UnmarshalTips(msg *api.MsgData, detail proto.Message) error {
	var tips api.TipsComm
	if err := proto.Unmarshal(msg.Content, &tips); err != nil {
		return utils.Wrap(err, "")
	}
	if err := proto.Unmarshal(tips.Detail, detail); err != nil {
		return utils.Wrap(err, "")
	}
	return nil
}
