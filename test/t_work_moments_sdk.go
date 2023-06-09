package test

import (
	"fmt"
	"drobot_sdk/drobot_im_sdk"
	"drobot_sdk/pkg/utils"
)

type WBase struct {
}

func (WBase) OnError(errCode int32, errMsg string) {
	fmt.Println("get workmoments OnError", errCode, errMsg)
}
func (WBase) OnSuccess(data string) {
	fmt.Println("get workmoments OnSuccess, ", data)
}

func (WBase) OnProgress(progress int) {
	fmt.Println("OnProgress, ", progress)
}

func TestGetWorkMomentsUnReadCount() {
	operationID := utils.OperationIDGenerator()
	var cb WBase
	drobot_im_sdk.GetWorkMomentsUnReadCount(cb, operationID)
}

func TestGetWorkMomentsNotification() {
	operationID := utils.OperationIDGenerator()
	var cb WBase
	offset := 0
	count := 10
	drobot_im_sdk.GetWorkMomentsNotification(cb, operationID, offset, count)
}

func TestClearWorkMomentsNotification() {
	operationID := utils.OperationIDGenerator()
	var cb WBase
	drobot_im_sdk.ClearWorkMomentsNotification(cb, operationID)
}
