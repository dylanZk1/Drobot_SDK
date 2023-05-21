package business

import (
	"drobot_sdk/drobot_sdk_callback"
	"drobot_sdk/pkg/log"
)

func (w *Business) SetListener(callback drobot_sdk_callback.OnCustomBusinessListener) {
	if callback == nil {
		log.NewError("", "callback is null")
		return
	}
	log.NewDebug("", "callback set success")
	w.listener = callback
}
