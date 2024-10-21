package systemservice

import (
	"gin_study/factory"
	"gin_study/gen/models"
	"gin_study/gen/mysql"
	"gin_study/gen/query"
	"gin_study/gen/request"
	"strconv"
)

func RegisterEquipment(req *request.RegisterEquipmentRequest) (int64, error) {
	equipmentRawStr := req.Vender + req.Type + req.Sn + req.Imei1 + req.Imei0 + req.Os + strconv.Itoa(req.IsPhysicalDevice)
	fingerprint := factory.Sha256Hash(equipmentRawStr)
	equip, _ := query.Equipment.Where(query.Equipment.Fingerprint.Eq(fingerprint)).First()
	if equip != nil {
		return equip.ID, nil
	} 
	equipment := models.Equipment{
		Vender:           req.Vender,
		Type:             req.Type,
		Sn:               req.Sn,
		Imei1:            req.Imei1,
		Imei0:            req.Imei0,
		Os:               req.Os,
		IsPhysicalDevice: req.IsPhysicalDevice,
		UserIDs:          "0",
		Fingerprint:      fingerprint,
	}
	tx := query.Q.Begin()
	err := query.Equipment.Save(&equipment)
	err = mysql.DeferTx(tx, err)
	return equipment.ID, err
}

func LogError(userID int64, equipmentID int64, req *request.LogErrorRequest) (int64, error) {
	errorLog := models.FontLogs{
		EquipmentID: equipmentID,
		UserID:      userID,
		Version:     req.Version,
		Stack:       req.Stack,
		Error:       req.Error,
	}
	tx := query.Q.Begin()
	err := query.FontLogs.Save(&errorLog)
	err = mysql.DeferTx(tx, err)
	return errorLog.ID, err
}
