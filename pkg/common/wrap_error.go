package common

//import (
//	"github.com/mitchellh/mapstructure"
//	"drobot_im_sdk/drobot_sdk_callback"
//	"drobot_im_sdk/pkg/db"
//	"drobot_im_sdk/pkg/db/model_struct"
//)
//
//func GetGroupMemberListByGroupID(callback drobot_sdk_callback.Base, operationID string, db *db.DataBase, groupID string) []*model_struct.LocalGroupMember {
//	memberList, err := db.GetGroupMemberListByGroupID(groupID)
//	CheckDBErrCallback(callback, err, operationID)
//	return memberList
//}
//
//func MapstructureDecode(input interface{}, output interface{}, callback drobot_sdk_callback.Base, oprationID string) {
//	err := mapstructure.Decode(input, output)
//	CheckDataErrCallback(callback, err, oprationID)
//}
