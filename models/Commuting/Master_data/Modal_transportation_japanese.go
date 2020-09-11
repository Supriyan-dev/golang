package Master_data

import (
	initMasterData "../../../initialize/Commuting"
	"../../../models"
	"log"
)

type Init_DB_CommutingMaster_Data models.DB_init

func (model Init_DB_CommutingMaster_Data) GetDataTransportation() (sh []initMasterData.Init_masterDataTransportationJapanese, err error) {

	GetDataTransportation, errDataTransportation := model.DB.Query(`SELECT id_master_transportation, code_transportation, name_transportation_english, name_transportation_japanese FROM master_transportation `)

	if errDataTransportation != nil {
		log.Println(errDataTransportation.Error())
	}
	defer GetDataTransportation.Close()
	var ms initMasterData.Init_masterDataTransportationJapanese

	for GetDataTransportation.Next() {

		errScanData := GetDataTransportation.Scan(&ms.IdMasterTransportation, &ms.CodeTransportation, &ms.NameTransportationEnglish, &ms.NameTransportationJapanese)

		if errScanData != nil {
			log.Println(errScanData)
		} else {

			sh = append(sh, ms)
		}
	}
	return sh, nil
}
