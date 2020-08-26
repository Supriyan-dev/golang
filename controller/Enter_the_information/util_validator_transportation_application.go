package enter_the_information

import (
	db2 "../../db"
	"log"
)

func CheckDataById(sql string, id string) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}

func ValidatorInsertBasicInformation(Request *InsertBasicInformation) (valid bool, message string) {

	if Request.CarInsuranceDocumentExpiryDate == "" {
		return false, "Missing required field in body request → CarInsuranceDocumentExpiryDate = <empty string>"
	}

	if Request.PropertyDamage == "" {
		return false, "Missing required field in body request → PropertyDamage = <empty string>"
	}

	if Request.PersonalInjury == "" {
		return false, "Missing required field in body request → PersonInjury = <empty string>"
	}

	if Request.DriverLicenseExpiryDate == "" {
		return false, "Missing required field in body request → DriverLicenseExpiryDate = <empty string>"
	}

	if Request.InsuranceCompany == "" {
		return false, "Missing required field in body request → InsureCompany = <empty string>"
	}

	if Request.IdGeneralInformation == "" {
		return false, "Missing required field in body request → IdGeneralInformation = <empty string>"
	}

	//Data := "select count(*) from commuting_basic_information where id_general_information = ?"
	//
	//CheckData := CheckDataById(Data, Request.IdGeneralInformation)
	//log.Println(CheckData)
	//if CheckData >= 1 {
	//	return false, "IdGeneralInformation Must 1"
	//}
	return true, "done"
}

func ValidatorInsertUsageRecordApplyForTravelExpenses(Request *InsertTransportationApplication) (valid bool, message string) {

	if Request.RouteProfileName == "" {
		return false, "Missing required field in body request → RouteProfileName  = <empty string>"
	}
	if Request.Date == "" {
		return false, "Missing required field in body request → Date  = <empty string>"
	}
	if Request.Attendance == "" {
		return false, "Missing required field in body request → Attendance  = <empty string>"
	}
	if Request.CodeCommuting == "" {
		return false, "Missing required field in body request → CodeCommuting  = <empty string>"
	}
	if Request.IdGeneralInformation == "" {
		return false, "Missing required field in body request → IdGeneralInformation  = <empty string>"
	}
	return true, "done"
}

func ValidatorDetailInsertUsageRecordApplyForTravelExpenses(Request *InsertDetailTransportationApplication) (valid bool, message string) {

	if Request.IdCommutingTrip == 0 {
		return false, "Missing required field in body request → TypeOfTransport = 0"
	}
	if Request.TypeOfTransport == "" {
		return false, "Missing required field in body request → TypeOfTransport = <empty string>"
	}
	if Request.Purpose == "" {
		return false, "Missing required field in body request → Purpose = <empty string>"
	}
	if Request.DetailFrom == "" {
		return false, "Missing required field in body request → DetailFrom = <empty string>"
	}
	if Request.DetailTo == "" {
		return false, "Missing required field in body request → DetailTo = <empty string>"
	}
	if Request.TransitPoint == "" {
		return false, "Missing required field in body request → TransitPoint = <empty string>"
	}

	return true, "done"

}
