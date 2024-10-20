package systemservice

import (
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
)

func RegisterEquipment(req *request.RegisterEquipmentRequest) (int64, error) {
	equipment := models.Equipment{
		Vender:  req.Vender,
		Type:    req.Type,
		Sn:      req.Sn,
		Imei1:   req.Imei1,
		Imei0:   req.Imei0,
		Os:      req.Os,
		UserIDs: "0",
	}
	tx := query.Q.Begin()
	err := query.Equipment.Save(&equipment)
	err = mysql.DeferTx(tx, err)
	return equipment.ID, err
}
