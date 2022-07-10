/**
  @author:panliang
  @data:2022/6/8
  @note
**/
package dao

import (
	"github.com/golang-module/carbon"
	"im-services/app/models/offline_message"
	"im-services/pkg/date"
	"im-services/pkg/model"
)

// 拉取离线私聊消息
func (offline *OfflineMessageDao) PullPrivateOfflineMessage(id string) []offline_message.ImOfflineMessages {

	var list []offline_message.ImOfflineMessages

	// 拉去最近半个月内的消息记录
	timeStamp := carbon.Parse(date.NewDate()).SubDays(15).Timestamp()

	model.DB.Table("im_offline_messages").
		Where("status=0 and receive_id=? and send_time>?", id, timeStamp).
		Find(&list)

	return list
}

// 更新消息状态
func (offline *OfflineMessageDao) UpdatePrivateOfflineMessageStatus(id string) {
	model.DB.Table("im_offline_messages").
		Where("status=0 and receive_id=?", id).
		Updates(map[string]interface{}{"status": 1})
}