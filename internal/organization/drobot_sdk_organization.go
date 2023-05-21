package organization

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/common"
	"drobot_sdk/pkg/log"
	"drobot_sdk/pkg/sdk_params_callback"
	"drobot_sdk/pkg/utils"
)

func (o *Organization) SetListener(callback drobot_sdk_callback.OnOrganizationListener) {
	if callback == nil {
		return
	}
	o.listener = callback
}

func (o *Organization) GetSubDepartment(callback drobot_sdk_callback.Base, departmentID string, offset, count int, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", departmentID, offset, count)
		result := o.getSubDepartment(callback, departmentID, offset, count, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) GetDepartmentMember(callback drobot_sdk_callback.Base, departmentID string, offset, count int, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", departmentID, offset, count)
		result := o.getDepartmentMember(callback, departmentID, offset, count, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) GetUserInDepartment(callback drobot_sdk_callback.Base, userID string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", userID)
		result := o.getUserInDepartment(callback, userID, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) GetDepartmentMemberAndSubDepartment(callback drobot_sdk_callback.Base, departmentID string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", departmentID)
		result := o.getDepartmentMemberAndSubDepartment(callback, departmentID, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) GetParentDepartmentList(callback drobot_sdk_callback.Base, departmentID string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", departmentID)
		result := o.getParentDepartmentList(callback, departmentID, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) GetDepartmentInfo(callback drobot_sdk_callback.Base, departmentID string, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", departmentID)
		result := o.getDepartmentInfo(callback, departmentID, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}

func (o *Organization) SearchOrganization(callback drobot_sdk_callback.Base, input string, offset, count int, operationID string) {
	if callback == nil {
		return
	}
	fName := utils.GetSelfFuncName()
	go func() {
		log.NewInfo(operationID, fName, "args: ", input)
		var searchParam sdk_params_callback.SearchOrganizationParams
		common.JsonUnmarshalAndArgsValidate(input, &searchParam, callback, operationID)
		result := o.searchOrganization(callback, searchParam, offset, count, operationID)
		resp := utils.StructToJsonStringDefault(result)
		callback.OnSuccess(resp)
		log.NewInfo(operationID, fName, " callback: ", resp)
	}()
}
