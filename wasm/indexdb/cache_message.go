package indexdb

import (
	"drobot_sdk/pkg/db/model_struct"
	"drobot_sdk/pkg/utils"
)

type LocalCacheMessage struct {
}

func (i *LocalCacheMessage) BatchInsertTempCacheMessageList(MessageList []*model_struct.TempCacheLocalChatLog) error {
	_, err := Exec(utils.StructToJsonString(MessageList))
	return err
}

func (i *LocalCacheMessage) InsertTempCacheMessage(Message *model_struct.TempCacheLocalChatLog) error {
	_, err := Exec(utils.StructToJsonString(Message))
	return err
}
