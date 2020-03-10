package service

import (
	"sync"

	"server/model"
)

// Vod service.
var Vod = &vodService{
	mutex: &sync.Mutex{},
}

type vodService struct {
	mutex *sync.Mutex
}

const (
	VodStatusOK = iota
)

func (srv *vodService) GetVods(typeId int, typeId1 int, page int, pageSize int) (ret []*model.VingVod) {

	offset := (page - 1) * pageSize

	logger.Info("offset --", offset)

	logger.Info("type id --  ", typeId, ", typeId1 = ", typeId1, ", pagesize = ", pageSize)
	if err := db.Model(&model.VingVod{}).
		Where("`type_id` = ? AND `type_id_1` = ?", typeId, typeId1).
		Order("`type_id` DESC").
		Offset(offset).Limit(pageSize).
		Find(&ret).Error; nil != err {
		logger.Errorf("get vods failed: " + err.Error())
	}
	return
}
