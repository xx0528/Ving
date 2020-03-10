package service

import (
	"server/model"
)

func TestGetVods(typeId int, groupId int, page int, pageSize int) (ret []*model.VingVod, count int) {

	vodList, count := Vod.GetVods(4, 0, 1, 20)
	logger.Debugf("vodList [%#v]", vodList)
	return
}
