package test

import (
	"encoding/json"
	"fmt"
	"drobot_sdk/drobot_im_sdk"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/utils"
)

type orgBaseCallback struct {
	OperationID string
	CallName    string
}

func (t orgBaseCallback) OnSuccess(data string) {
	fmt.Println("=======================\n", t.OperationID, t.CallName, utils.GetSelfFuncName(), data)
}

func (t orgBaseCallback) OnError(errCode int32, errMsg string) {
	log.Info(t.OperationID, t.CallName, utils.GetSelfFuncName(), errCode, errMsg)
}

type testOrganization struct {
	orgBaseCallback
}

func DoTestGetSubDepartment() {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	log.Info(test.OperationID, utils.GetSelfFuncName(), "input: ")
	drobot_im_sdk.GetSubDepartment(test, test.OperationID, "001", 0, 1)
}

func DoTestGetDepartmentMember() {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	log.Info(test.OperationID, utils.GetSelfFuncName(), "input: ")
	drobot_im_sdk.GetDepartmentMember(test, test.OperationID, "001", 0, 1)
}

func DoTestGetUserInDepartment() {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	log.Info(test.OperationID, utils.GetSelfFuncName(), "input: ")
	drobot_im_sdk.GetUserInDepartment(test, test.OperationID, "13900000000")
}

func DoTestGetDepartmentMemberAndSubDepartment() {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	log.Info(test.OperationID, utils.GetSelfFuncName(), "input: ")
	drobot_im_sdk.GetDepartmentMemberAndSubDepartment(test, test.OperationID, "002")
}

func DoTestGetParentDepartmentList(departmentID string) {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	log.Info(test.OperationID, utils.GetSelfFuncName(), "input: ")
	drobot_im_sdk.GetParentDepartmentList(test, test.OperationID, departmentID)
}

func DoTestGetDepartmentInfo(departmentID string) {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	drobot_im_sdk.GetDepartmentInfo(test, test.OperationID, departmentID)
}

func DoTestSearchOrganization(input string, offset, count int) {
	var test testOrganization
	test.OperationID = utils.OperationIDGenerator()
	test.CallName = utils.GetSelfFuncName()
	params := sdk_params_callback.SearchOrganizationParams{
		KeyWord:                 input,
		IsSearchUserName:        true,
		IsSearchUserEnglishName: true,
		IsSearchPosition:        true,
		IsSearchUserID:          true,
		IsSearchMobile:          true,
		IsSearchEmail:           true,
		IsSearchTelephone:       true,
	}
	bytes, _ := json.Marshal(params)
	fmt.Println("input params", string(bytes))
	drobot_im_sdk.SearchOrganization(test, test.OperationID, string(bytes), offset, count)
}
