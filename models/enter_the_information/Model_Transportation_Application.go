package enter_the_information

import (
	"Go_DX_Services/models"
	"log"
)

type models_init models.DB_init

func (model models_init) GetIdByCodeCommuting(store_number string, employee_number string) (sh *[]ShowDetailTransportationApplication, err error) {

	rows, err := model.DB.Query(`select a.route_profile_name, a.date, b.type_of_transport, a.attendance_code, b.purpose
									   from commuting_trip a,detail_commuting_trip b where a.id_commuting_trip = b.id_commuting_trip 
									   and a.code_commuting =? `, store_number)
	var init_container ShowDetailTransportationApplication
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		rows.Scan(&init_container.RouteProfileName, &init_container.Date, &init_container)
	}

	return sh, nil

}

func (model models_init) InsertTransportationApplication(insertD *InsertTransportationApplication)(it []InsertTransportationApplication,  err error){

	rows, err := model.DB.Prepare(`insert into commuting_trip(route_profile_name,date,attendance_code,code_commuting) VALUES(?,?,?,?)`)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	rows.Exec(insertD.RouteProfileName,insertD.Date,insertD.Attendance,insertD.CodeCommuting)

	datainsert := InsertTransportationApplication{
		RouteProfileName: insertD.RouteProfileName,
		Date:             insertD.Date,
		Attendance:       insertD.Attendance,
		CodeCommuting:    insertD.CodeCommuting,
	}

	it = append(it,datainsert)

	return it, nil

}

