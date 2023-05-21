package workMoments

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/utils"
)

func (w *WorkMoments) SetListener(callback drobot_sdk_callback.OnWorkMomentsListener) {
	if callback == nil {
		log.NewError("", "callback is null")
		return
	}
	log.NewDebug("", "callback set success")
	w.listener = callback
}

func (w *WorkMoments) GetWorkMomentsUnReadCount(callback drobot_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName)
		result := w.getWorkMomentsNotificationUnReadCount(callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, fName)
	}()
}

func (w *WorkMoments) GetWorkMomentsNotification(callback drobot_sdk_callback.Base, offset, count int, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, offset, count)
		result := w.getWorkMomentsNotification(offset, count, callback, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, fName)
	}()
}

func (w *WorkMoments) ClearWorkMomentsNotification(callback drobot_sdk_callback.Base, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName)
		w.clearWorkMomentsNotification(callback, operationID)
		callback.OnSuccess(sdk_params_callback.ClearWorkMomentsMessageCallback)
		log.NewInfo(operationID, fName)
	}()
}
