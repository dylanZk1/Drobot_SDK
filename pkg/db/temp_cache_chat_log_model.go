package db

import (
	"drobot_sdk/pkg/db/model_struct"
	"drobot_sdk/pkg/utils"
)

func (d *DataBase) BatchInsertTempCacheMessageList(MessageList []*model_struct.TempCacheLocalChatLog) error {
	if MessageList == nil {
		return nil
	}
	return utils.Wrap(d.conn.Create(MessageList).Error, "BatchInsertTempCacheMessageList failed")
}
func (d *DataBase) InsertTempCacheMessage(Message *model_struct.TempCacheLocalChatLog) error {

	return utils.Wrap(d.conn.Create(Message).Error, "InsertTempCacheMessage failed")

}
