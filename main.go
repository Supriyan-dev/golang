package main

import (
	"fmt"
	"net/http"

	entertheinformation "./controller/Enter_the_information"
	controllerDataMaster "./controller/data_master_controller"
	controllerPermissionToDrive "./controller/list_input_information"
	login "./login_controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	// start login user
	router.HandleFunc("/generate-hash/{password}", login.GenerateHashPassword).Methods("GET")
	// router.HandleFunc("/login", login.CheckLogin).Methods("POST")
	//end login user

	// Start permission to drive
	router.HandleFunc("/permission_to_drive", controllerPermissionToDrive.PermissionToDrive).Methods("GET")
	router.HandleFunc("/permission_to_drive", controllerPermissionToDrive.PermissionToDriveUpdate).Methods("PUT")
	router.HandleFunc("/permission_to_drive/{page}/{perPage}", controllerPermissionToDrive.PermissionToDrivePagination).Methods("GET")
	// end permission to drive

	// start data master
	// start crud store information
	router.HandleFunc("/storeinformation", controllerDataMaster.ReturnAllStoreInformation)
	router.HandleFunc("/storeinformation/{page}/{perPage}", controllerDataMaster.ReturnAllStoreInformationPagination)
	router.HandleFunc("/storeinformation/get", controllerDataMaster.GetStoreInformation)
	router.HandleFunc("/storeinformation/create", controllerDataMaster.CreateStoreInformation)
	router.HandleFunc("/storeinformation/update", controllerDataMaster.UpdateStoreInformation)
	router.HandleFunc("/storeinformation/delete", controllerDataMaster.DeleteStoreInformation)
	// // end crud store information

	// start crud departement information
	router.HandleFunc("/departement-information", controllerDataMaster.ReturnAllDepartementInformation)
	router.HandleFunc("/departement-information/{page}/{perPage}", controllerDataMaster.ReturnAllDepartementInformationPagination).Methods("GET")
	router.HandleFunc("/departement-information/get", controllerDataMaster.GetDepartementInformation)
	router.HandleFunc("/departement-information/create", controllerDataMaster.CreateDepartementInformation)
	router.HandleFunc("/departement-information/update", controllerDataMaster.UpdateDepartementInformation)
	router.HandleFunc("/departement-information/delete", controllerDataMaster.DeleteDepartementInformation)
	// end crud deaprtemen information

	// start crud srtore section information
	router.HandleFunc("/store-section-information", controllerDataMaster.ReturnAllStroreSectionInformation)
	router.HandleFunc("/store-section-information/{page}/{perPage}", controllerDataMaster.ReturnAllStroreSectionInformationPagination).Methods("GET")
	router.HandleFunc("/store-section-information/get", controllerDataMaster.GetStoreSectionInformation)
	router.HandleFunc("/store-section-information/create", controllerDataMaster.CreateStoreSectionInformation)
	router.HandleFunc("/store-section-information/update", controllerDataMaster.UpdateStoreSectionInformation)
	router.HandleFunc("/store-section-information/delete", controllerDataMaster.DeleteStoreSectionInformation)
	//end crud store section infomration

	// start crud unit information
	router.HandleFunc("/unit-information", controllerDataMaster.ReturnAllUnitInformation)
	router.HandleFunc("/unit-information/{page}/{perPage}", controllerDataMaster.ReturnAllUnitInformationPagination).Methods("GET")
	router.HandleFunc("/unit-information/get", controllerDataMaster.GetUnitInformation)
	router.HandleFunc("/unit-information/create", controllerDataMaster.CreateUnitInformation)
	router.HandleFunc("/unit-information/update", controllerDataMaster.UpdateUnitInformation)
	router.HandleFunc("/unit-information/delete", controllerDataMaster.DeleteUnitInformation)
	// end crud unit information

	// start crud prefecture
	router.HandleFunc("/prefecture", controllerDataMaster.ReturnAllPrefect)
	router.HandleFunc("/prefecture/{page}/{perPage}", controllerDataMaster.ReturnAllPrefectPagination).Methods("GET")
	router.HandleFunc("/prefecture/get", controllerDataMaster.GetPrefect)
	router.HandleFunc("/prefecture/create", controllerDataMaster.CreatePrefect)
	router.HandleFunc("/prefecture/update", controllerDataMaster.UpdatePrefect)
	router.HandleFunc("/prefecture/delete", controllerDataMaster.DeletePrefect)
	// end crud prefecture

	// // start crud bank
	router.HandleFunc("/bank", controllerDataMaster.ReturnAllBank)
	router.HandleFunc("/bank/{page}/{perPage}", controllerDataMaster.ReturnAllBankPagination).Methods("GET")
	router.HandleFunc("/bank/get", controllerDataMaster.GetBank)
	router.HandleFunc("/bank/create", controllerDataMaster.CreateBank)
	router.HandleFunc("/bank/update", controllerDataMaster.UpdateBank)
	router.HandleFunc("/bank/delete", controllerDataMaster.DeleteBank)

	// start crud exp category

	// end crud bank

	// start crud full time salary
	router.HandleFunc("/full-time-salary", controllerDataMaster.ReturnAllFullTimeSalary)
	router.HandleFunc("/full-time-salary/{page}/{perPage}", controllerDataMaster.ReturnAllFullTimeSalaryPagination).Methods("GET")
	router.HandleFunc("/full-time-salary/get", controllerDataMaster.GetFullTimeSalary)
	router.HandleFunc("/full-time-salary/create", controllerDataMaster.CreateFullTimeSalary)
	router.HandleFunc("/full-time-salary/update", controllerDataMaster.UpdateFullTimeSalary)
	router.HandleFunc("/full-time-salary/delete", controllerDataMaster.DeleteFullTimeSalary)
	// end crud full time salary

	// start crud part time salary
	router.HandleFunc("/part-time-above-18-salary", controllerDataMaster.ReturnAllPartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeAbove18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/get", controllerDataMaster.GetPartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/create", controllerDataMaster.CreatePartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/update", controllerDataMaster.UpdatePartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/delete", controllerDataMaster.DeletePartTimeAbove18Salary)
	// end crud part time salary

	// start crud under 18 salary
	router.HandleFunc("/part-time-under-18-salary", controllerDataMaster.ReturnAllPartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeUnder18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/get", controllerDataMaster.GetPartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/create", controllerDataMaster.CreatePartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/update", controllerDataMaster.UpdatePartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/delete", controllerDataMaster.DeletePartTimeUnder18Salary)
	// end crud under 18 salary

	// start crud user
	router.HandleFunc("/user", login.CheckLogin(controllerDataMaster.ReturnAllUser))
	router.HandleFunc("/user/{page}/{perPage}", controllerDataMaster.ReturnAllUserPagination).Methods("GET")
	router.HandleFunc("/user/get", controllerDataMaster.GetUser)
	router.HandleFunc("/user/create", controllerDataMaster.CreateUser)
	router.HandleFunc("/user/update", controllerDataMaster.UpdateUser)
	router.HandleFunc("/user/delete", controllerDataMaster.DeleteUser)
	// end crud user

	// // start crud exp category

	router.HandleFunc("/exp-category", controllerDataMaster.ReturnAllExpCategory)
	router.HandleFunc("/exp-category/{page}/{perPage}", controllerDataMaster.ReturnAllExpCategoryPagination).Methods("GET")
	router.HandleFunc("/exp-category/get", controllerDataMaster.GetExpCategory)
	router.HandleFunc("/exp-category/create", controllerDataMaster.CreateExpCategory)
	router.HandleFunc("/exp-category/update", controllerDataMaster.UpdateExpCategory)
	router.HandleFunc("/exp-category/delete", controllerDataMaster.DeleteExpCategory)
	// // end crud exp category
	// //end data master

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
	router.HandleFunc("/commuting-transportation-applicationCheckData", entertheinformation.ReturnGetByCommutingUsageRecord).Methods("POST")
	router.HandleFunc("/commuting-UsageRecord-Apply", entertheinformation.ReturnInsertUsageRecordApplyForTravelExpenses).Methods("POST")
	router.HandleFunc("/commuting-UsageRecord-DetailApply", entertheinformation.ReturnDetailInsertUsageRecordApplyForTravelExpenses).Methods("POST")
	// end Commuting Transportation Application

	fmt.Println("Connected to port 9000")
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":9000", handler)

}
