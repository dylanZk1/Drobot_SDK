package business

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/db/db_interface"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/utils"
)

type Business struct {
	listener drobot_sdk_callback.OnCustomBusinessListener
	db       db_interface.DataBase
}

func NewBusiness(db db_interface.DataBase) *Business {
	return &Business{
		db: db,
	}
}

func (b *Business) DoNotification(jsonDetailStr string, operationID string) {
	if b.listener == nil {
		log.NewDebug(operationID, "WorkMoments listener is null", jsonDetailStr)
		return
	}
	log.NewInfo(operationID, utils.GetSelfFuncName(), "json_detail: ", jsonDetailStr)
	b.listener.OnRecvCustomBusinessMessage(jsonDetailStr)
}
