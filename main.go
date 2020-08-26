package main

import (
	entertheinformation "./controller/Enter_the_information"
	controllerDataMaster "./controller/data_master_controller"
	controllerPermissionToDrive "./controller/list_input_information"
	"fmt"
	"github.com/rs/cors"
	"net/http"

	// controllerDataMaster "github.com/jeffri/golang-test/GO_DX_SERVICES/controller/data_master_controller"
	// controllerPermissionToDrive "github.com/jeffri/golang-test/GO_DX_SERVICES/controller/list_input_information"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// Start permission to drive
	router.HandleFunc("/permission_to_drive", controllerPermissionToDrive.PermissionToDrive).Methods("GET")
	router.HandleFunc("/permission_to_drive", controllerPermissionToDrive.PermissionToDriveUpdate).Methods("PUT")
	router.HandleFunc("/permission_to_drive/{page}/{perPage}", controllerPermissionToDrive.PermissionToDrivePagination).Methods("GET")
	// end permission to drive

	// start data master
	// start crud store information
	router.HandleFunc("/storeinformation", controllerDataMaster.ReturnAllStoreInformation).Methods("GET")
	router.HandleFunc("/storeinformation/{page}/{perPage}", controllerDataMaster.ReturnAllStoreInformationPagination).Methods("GET")
	router.HandleFunc("/storeinformation/{id_code_store}", controllerDataMaster.GetStoreInformation).Methods("GET")
	router.HandleFunc("/storeinformation", controllerDataMaster.CreateStoreInformation).Methods("POST")
	router.HandleFunc("/storeinformation", controllerDataMaster.UpdateStoreInformation).Methods("PUT")
	router.HandleFunc("/storeinformation/{id_code_store}", controllerDataMaster.DeleteStoreInformation).Methods("DELETE")
	// end crud store information

	// start crud departement information
	router.HandleFunc("/departement-information", controllerDataMaster.ReturnAllDepartementInformation).Methods("GET")
	router.HandleFunc("/departement-information/{page}/{perPage}", controllerDataMaster.ReturnAllDepartementInformationPagination).Methods("GET")
	router.HandleFunc("/departement-information/{id_department}", controllerDataMaster.GetDepartementInformation).Methods("GET")
	router.HandleFunc("/departement-information", controllerDataMaster.CreateDepartementInformation).Methods("POST")
	router.HandleFunc("/departement-information", controllerDataMaster.UpdateDepartementInformation).Methods("PUT")
	router.HandleFunc("/departement-information/{id_department}", controllerDataMaster.DeleteDepartementInformation).Methods("DELETE")
	// end crud deaprtemen information

	// start crud srtore section information
	router.HandleFunc("/store-section-information", controllerDataMaster.ReturnAllStroreSectionInformation).Methods("GET")
	router.HandleFunc("/store-section-information/{page}/{perPage}", controllerDataMaster.ReturnAllStroreSectionInformationPagination).Methods("GET")
	router.HandleFunc("/store-section-information/{id_store_section}", controllerDataMaster.GetStoreSectionInformation).Methods("GET")
	router.HandleFunc("/store-section-information", controllerDataMaster.CreateStoreSectionInformation).Methods("POST")
	router.HandleFunc("/store-section-information", controllerDataMaster.UpdateStoreSectionInformation).Methods("PUT")
	router.HandleFunc("/store-section-information/{id_store_section}", controllerDataMaster.DeleteStoreSectionInformation).Methods("DELETE")
	//end crud store section infomration

	// start crud unit information
	router.HandleFunc("/unit-information", controllerDataMaster.ReturnAllUnitInformation).Methods("GET")
	router.HandleFunc("/unit-information/{page}/{perPage}", controllerDataMaster.ReturnAllUnitInformationPagination).Methods("GET")
	router.HandleFunc("/unit-information/{id_unit}", controllerDataMaster.GetUnitInformation).Methods("GET")
	router.HandleFunc("/unit-information", controllerDataMaster.CreateUnitInformation).Methods("POST")
	router.HandleFunc("/unit-information", controllerDataMaster.UpdateUnitInformation).Methods("PUT")
	router.HandleFunc("/unit-information/{id_unit}", controllerDataMaster.DeleteUnitInformation).Methods("DELETE")
	// end crud unit information

	// start crud prefecture
	router.HandleFunc("/prefecture", controllerDataMaster.ReturnAllPrefect).Methods("GET")
	router.HandleFunc("/prefecture/{page}/{perPage}", controllerDataMaster.ReturnAllPrefectPagination).Methods("GET")
	router.HandleFunc("/prefecture/{id_prefecture}", controllerDataMaster.GetPrefect).Methods("GET")
	router.HandleFunc("/prefecture", controllerDataMaster.CreatePrefect).Methods("POST")
	router.HandleFunc("/prefecture", controllerDataMaster.UpdatePrefect).Methods("PUT")
	router.HandleFunc("/prefecture/{id_prefecture}", controllerDataMaster.DeletePrefect).Methods("DELETE")
	// end crud prefecture

	// start crud bank
	router.HandleFunc("/bank", controllerDataMaster.ReturnAllBank).Methods("GET")
	router.HandleFunc("/bank/{page}/{perPage}", controllerDataMaster.ReturnAllBankPagination).Methods("GET")
	router.HandleFunc("/bank/{id_bank}", controllerDataMaster.GetBank).Methods("GET")
	router.HandleFunc("/bank", controllerDataMaster.CreateBank).Methods("POST")
	router.HandleFunc("/bank", controllerDataMaster.UpdateBank).Methods("PUT")
	router.HandleFunc("/bank/{id_bank}", controllerDataMaster.DeleteBank).Methods("DELETE")
	// end crud bank

	// start crud full time salary
	router.HandleFunc("/full-time-salary", controllerDataMaster.ReturnAllFullTimeSalary).Methods("GET")
	router.HandleFunc("/full-time-salary/{page}/{perPage}", controllerDataMaster.ReturnAllFullTimeSalaryPagination).Methods("GET")
	router.HandleFunc("/full-time-salary/{id_full_time_salary}", controllerDataMaster.GetFullTimeSalary).Methods("GET")
	router.HandleFunc("/full-time-salary", controllerDataMaster.CreateFullTimeSalary).Methods("POST")
	router.HandleFunc("/full-time-salary", controllerDataMaster.UpdateFullTimeSalary).Methods("PUT")
	router.HandleFunc("/full-time-salary/{id_full_time_salary}", controllerDataMaster.DeleteFullTimeSalary).Methods("DELETE")
	// end crud full time salary

	// start crud part time salary
	router.HandleFunc("/part-time-above-18-salary", controllerDataMaster.ReturnAllPartTimeAbove18Salary).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeAbove18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/{id_part_time_above_18_salary}", controllerDataMaster.GetPartTimeAbove18Salary).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary", controllerDataMaster.CreatePartTimeAbove18Salary).Methods("POST")
	router.HandleFunc("/part-time-above-18-salary", controllerDataMaster.UpdatePartTimeAbove18Salary).Methods("PUT")
	router.HandleFunc("/part-time-above-18-salary/{id_part_time_above_18_salary}", controllerDataMaster.DeletePartTimeAbove18Salary).Methods("DELETE")
	// end crud part time salary

	// start crud under 18 salary
	router.HandleFunc("/part-time-under-18-salary", controllerDataMaster.ReturnAllPartTimeUnder18Salary).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeUnder18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/{id_part_time_under_18_salary}", controllerDataMaster.GetPartTimeUnder18Salary).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary", controllerDataMaster.CreatePartTimeUnder18Salary).Methods("POST")
	router.HandleFunc("/part-time-under-18-salary", controllerDataMaster.UpdatePartTimeUnder18Salary).Methods("PUT")
	router.HandleFunc("/part-time-under-18-salary/{id_part_time_under_18_salary}", controllerDataMaster.DeletePartTimeUnder18Salary).Methods("DELETE")
	// end crud under 18 salary

	// start crud user
	router.HandleFunc("/user", controllerDataMaster.ReturnAllUser).Methods("GET")
	router.HandleFunc("/user/{page}/{perPage}", controllerDataMaster.ReturnAllUserPagination).Methods("GET")
	router.HandleFunc("/user/{id_user}", controllerDataMaster.GetUser).Methods("GET")
	router.HandleFunc("/user", controllerDataMaster.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllerDataMaster.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id_user}", controllerDataMaster.DeleteUser).Methods("DELETE")
	// end crud user

	// start crud exp category
	router.HandleFunc("/exp-category", controllerDataMaster.ReturnAllExpCategory).Methods("GET")
	router.HandleFunc("/exp-category/{page}/{perPage}", controllerDataMaster.ReturnAllExpCategoryPagination).Methods("GET")
	router.HandleFunc("/exp-category/{id_exp}", controllerDataMaster.GetExpCategory).Methods("GET")
	router.HandleFunc("/exp-category", controllerDataMaster.CreateExpCategory).Methods("POST")
	router.HandleFunc("/exp-category", controllerDataMaster.UpdateExpCategory).Methods("PUT")
	router.HandleFunc("/exp-category/{id_exp}", controllerDataMaster.DeleteExpCategory).Methods("DELETE")
	// end crud exp category
	//end data master

	// router.HandleFunc("/basic-information", controller.ReturnAllBasicInformation).Methods("GET")
	// router.HandleFunc("/cash-claim", controller.ReturnAllCashClaim).Methods("GET")
	// router.HandleFunc("/cash-claim-code", controller.ReturnAllCashClaimCode).Methods("GET")
	// router.HandleFunc("/cash-claim-join", controller.ReturnAllCashClaimJoin).Methods("GET")
	// router.HandleFunc("/category-134", controller.ReturnAllCategory_134).Methods("GET")
	// router.HandleFunc("/category-136", controller.ReturnAllCategory_136).Methods("GET")
	// router.HandleFunc("/category-137", controller.ReturnAllCategory_137).Methods("GET")
	// router.HandleFunc("/category-138", controller.ReturnAllCategory_138).Methods("GET")
	// router.HandleFunc("/code-commuting", controller.ReturnAllCodeCommuting).Methods("GET")
	// router.HandleFunc("/commuting-basic-information", controller.ReturnAllCommutingBasicInformation).Methods("GET")
	// router.HandleFunc("/commuting-trip", controller.ReturnAllCommutingTrip).Methods("GET")

	router.HandleFunc("/commuting-basic-information", entertheinformation.ReturnCreateCommutingBasicInformation).Methods("POST")
	router.HandleFunc("/commuting-basic-information-CheckData", entertheinformation.ReturnGetByCommutingBasicInformation).Methods("POST")
	router.HandleFunc("/commuting-UsageRecord-CheckData", entertheinformation.ReturnGetByCommutingUsageRecord).Methods("POST")
	router.HandleFunc("/commuting-UsageRecord-Apply", entertheinformation.ReturnInsertUsageRecordApplyForTravelExpenses).Methods("POST")
	// end Commuting Transportation Application

	fmt.Println("Connected to port 9000")
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":9000", handler)

}