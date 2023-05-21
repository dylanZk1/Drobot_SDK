package wasm_wrapper

import (
	"drobot_sdk/drobot_im_sdk"
	"drobot_sdk/pkg/utils"
	"drobot_sdk/wasm/event_listener"
	"syscall/js"
)

//------------------------------------group---------------------------
type WrapperUser struct {
	*WrapperCommon
}

func NewWrapperUser(wrapperCommon *WrapperCommon) *WrapperUser {
	return &WrapperUser{WrapperCommon: wrapperCommon}
}

func (w *WrapperUser) GetSelfUserInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(drobot_im_sdk.GetSelfUserInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) SetSelfInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(drobot_im_sdk.SetSelfInfo, callback, &args).AsyncCallWithCallback()
}

func (w *WrapperUser) GetUsersInfo(_ js.Value, args []js.Value) interface{} {
	callback := event_listener.NewBaseCallback(utils.FirstLower(utils.GetSelfFuncName()), w.commonFunc)
	return event_listener.NewCaller(drobot_im_sdk.GetUsersInfo, callback, &args).AsyncCallWithCallback()

}
