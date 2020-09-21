package main

import (
	Approve "./controller/Commuting/Approve"
	entertheinformation "./controller/Commuting/transportation_application"
	list_GR "./controller/list_general_recruitment"
	controllerDriverManagement "./controller/driver_management_controller"

	controllerDataMaster "./controller/data_master_controller"
	generalRecrutment "./controller/general_recrutment_controller"
	controllerPermissionToDrive "./controller/list_input_information"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"os"

	"fmt"
	"net/http"

	loginController "./controller/login_controller"
	login "./login_controller"
	ForgotPassword "./controller/forgot_password"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// start login user
	// router.HandleFunc("/login", login.CheckLogin).Methods("POST")
	// end login user

	// start profile data user encrypt
	router.HandleFunc("/generate_hash_work_flow/{password}", login.GenerateHashPasswordWorkFlow).Methods("GET")
	router.HandleFunc("/generate_hash_data_master/{password}", login.GenerateHashPasswordDataMaster).Methods("GET")
	router.HandleFunc("/read_work_flow", login.CheckLogin(loginController.WorkFlowLogin))
	router.HandleFunc("/read_data_master", login.CheckLoginDataMaster(loginController.DataMasterLogin))
	// end profile data user encrypt

	// start forgot password
	router.HandleFunc("/forgot-passwordWithEmail", ForgotPassword.ReturnForgotPasswordWithEmail)
	router.HandleFunc("/forgot-passwordWithPin", ForgotPassword.ReturnForgotPasswordWithPin)
	router.HandleFunc("/forgot-password-action", ForgotPassword.ReturnForgotPasswordAction)
	// end forgot password

	// start data status approve general information status approve
	router.HandleFunc("/general_recrutment/create", generalRecrutment.DataGeneralRecrutment)
	// end data status approve  general information status approve

	router.HandleFunc("/generate_hash_work_flow/{password}", login.GenerateHashPasswordWorkFlow).Methods("GET")

	// Start permission to drive
	router.HandleFunc("/permission_to_drive", controllerPermissionToDrive.PermissionToDrive)
	router.HandleFunc("/permission_to_drive/update", controllerPermissionToDrive.PermissionToDriveUpdate)
	router.HandleFunc("/permission_to_drive/{page}/{perPage}", controllerPermissionToDrive.PermissionToDrivePagination)
	// end permission to drive

	// start data master
	// start crud store information
	router.HandleFunc("/storeinformation", controllerDataMaster.ReturnAllStoreInformation)
	router.HandleFunc("/storeinformation/filter", controllerDataMaster.ReturnAllFilterInformation)
	router.HandleFunc("/storeinformation/{page}/{perPage}", controllerDataMaster.ReturnAllStoreInformationPagination)
	router.HandleFunc("/storeinformation/get", controllerDataMaster.GetStoreInformation)
	router.HandleFunc("/storeinformation/create", controllerDataMaster.CreateStoreInformation)
	router.HandleFunc("/storeinformation/update", controllerDataMaster.UpdateStoreInformation)
	router.HandleFunc("/storeinformation/{id_code_store}", controllerDataMaster.DeleteStoreInformation)
	// // end crud store information

	// start crud departement information
	router.HandleFunc("/departement-information", controllerDataMaster.ReturnAllDepartementInformation)
	router.HandleFunc("/departement-information/{page}/{perPage}", controllerDataMaster.ReturnAllDepartementInformationPagination)
	router.HandleFunc("/departement-information/get", controllerDataMaster.GetDepartementInformation)
	router.HandleFunc("/departement-information/create", controllerDataMaster.CreateDepartementInformation)
	router.HandleFunc("/departement-information/update", controllerDataMaster.UpdateDepartementInformation)
	router.HandleFunc("/departement-information/{id_department}", controllerDataMaster.DeleteDepartementInformation)
	// end crud deaprtemen information

	// start crud srtore section information
	router.HandleFunc("/store-section-information", controllerDataMaster.ReturnAllStroreSectionInformation)
	router.HandleFunc("/store-section-information/{page}/{perPage}", controllerDataMaster.ReturnAllStroreSectionInformationPagination)
	router.HandleFunc("/store-section-information/get", controllerDataMaster.GetStoreSectionInformation)
	router.HandleFunc("/store-section-information/create", controllerDataMaster.CreateStoreSectionInformation)
	router.HandleFunc("/store-section-information/update", controllerDataMaster.UpdateStoreSectionInformation)
	router.HandleFunc("/store-section-information/{id_store_section}", controllerDataMaster.DeleteStoreSectionInformation)
	//end crud store section infomration

	// start crud unit information
	router.HandleFunc("/unit-information", controllerDataMaster.ReturnAllUnitInformation)
	router.HandleFunc("/unit-information/{page}/{perPage}", controllerDataMaster.ReturnAllUnitInformationPagination)
	router.HandleFunc("/unit-information/get", controllerDataMaster.GetUnitInformation)
	router.HandleFunc("/unit-information/create", controllerDataMaster.CreateUnitInformation)
	router.HandleFunc("/unit-information/update", controllerDataMaster.UpdateUnitInformation)
	router.HandleFunc("/unit-information/{id_unit}", controllerDataMaster.DeleteUnitInformation)
	// end crud unit information

	// start crud prefecture
	router.HandleFunc("/prefecture", controllerDataMaster.ReturnAllPrefect)
	router.HandleFunc("/prefecture/{page}/{perPage}", controllerDataMaster.ReturnAllPrefectPagination)
	router.HandleFunc("/prefecture/get", controllerDataMaster.GetPrefect)
	router.HandleFunc("/prefecture/create", controllerDataMaster.CreatePrefect)
	router.HandleFunc("/prefecture/update", controllerDataMaster.UpdatePrefect)
	router.HandleFunc("/prefecture/{id_prefecture}", controllerDataMaster.DeletePrefect)
	// end crud prefecture

	// // start crud bank
	router.HandleFunc("/bank", controllerDataMaster.ReturnAllBank)
	router.HandleFunc("/bank/{page}/{perPage}", controllerDataMaster.ReturnAllBankPagination)
	router.HandleFunc("/bank/get", controllerDataMaster.GetBank)
	router.HandleFunc("/bank/create", controllerDataMaster.CreateBank)
	router.HandleFunc("/bank/update", controllerDataMaster.UpdateBank)
	router.HandleFunc("/bank/{id_bank}", controllerDataMaster.DeleteBank)

	// start crud exp category

	// end crud bank

	// start crud full time salary
	router.HandleFunc("/full-time-salary", controllerDataMaster.ReturnAllFullTimeSalary)
	router.HandleFunc("/full-time-salary/{page}/{perPage}", controllerDataMaster.ReturnAllFullTimeSalaryPagination)
	router.HandleFunc("/full-time-salary/get", controllerDataMaster.GetFullTimeSalary)
	router.HandleFunc("/full-time-salary/create", controllerDataMaster.CreateFullTimeSalary)
	router.HandleFunc("/full-time-salary/update", controllerDataMaster.UpdateFullTimeSalary)
	router.HandleFunc("/full-time-salary/{id_full_time_salary}", controllerDataMaster.DeleteFullTimeSalary)
	// end crud full time salary

	// start crud part time salary
	router.HandleFunc("/part-time-above-18-salary", controllerDataMaster.ReturnAllPartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeAbove18SalaryPagination)
	router.HandleFunc("/part-time-above-18-salary/get", controllerDataMaster.GetPartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/create", controllerDataMaster.CreatePartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/update", controllerDataMaster.UpdatePartTimeAbove18Salary)
	router.HandleFunc("/part-time-above-18-salary/{id_part_time_above_18_salary}", controllerDataMaster.DeletePartTimeAbove18Salary)
	// end crud part time salary

	// start crud under 18 salary
	router.HandleFunc("/part-time-under-18-salary", controllerDataMaster.ReturnAllPartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/{page}/{perPage}", controllerDataMaster.ReturnAllPartTimeUnder18SalaryPagination)
	router.HandleFunc("/part-time-under-18-salary/get", controllerDataMaster.GetPartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/create", controllerDataMaster.CreatePartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/update", controllerDataMaster.UpdatePartTimeUnder18Salary)
	router.HandleFunc("/part-time-under-18-salary/{id_part_time_under_18_salary}", controllerDataMaster.DeletePartTimeUnder18Salary)
	// end crud under 18 salary

	// start crud user
	router.HandleFunc("/user", controllerDataMaster.ReturnAllUser)
	router.HandleFunc("/user/{page}/{perPage}", controllerDataMaster.ReturnAllUserPagination)
	router.HandleFunc("/user/get", controllerDataMaster.GetUser)
	router.HandleFunc("/user/create", controllerDataMaster.CreateUser)
	router.HandleFunc("/user/update", controllerDataMaster.UpdateUser)
	router.HandleFunc("/user/{id_user}", controllerDataMaster.DeleteUser)
	// end crud user

	// // start crud exp category

	router.HandleFunc("/exp-category", controllerDataMaster.ReturnAllExpCategory)
	router.HandleFunc("/exp-category/{page}/{perPage}", controllerDataMaster.ReturnAllExpCategoryPagination)
	router.HandleFunc("/exp-category/get", controllerDataMaster.GetExpCategory)
	router.HandleFunc("/exp-category/create", controllerDataMaster.CreateExpCategory)
	router.HandleFunc("/exp-category/update", controllerDataMaster.UpdateExpCategory)
	router.HandleFunc("/exp-category/{id_exp}", controllerDataMaster.DeleteExpCategory)
	// // end crud exp category
	// //end data master

// ==========================================================================================================================================================
	// start data driver management

	// start prefecture driver management 
	router.HandleFunc("/prefecture-driver-management", controllerDriverManagement.ReturnAllPrefect)
	router.HandleFunc("/prefecture-driver-management/{page}/{perPage}", controllerDriverManagement.ReturnAllPrefectPagination)
	router.HandleFunc("/prefecture-driver-management/get", controllerDriverManagement.GetPrefect)
	// router.HandleFunc("/prefecture-driver-management/create", controllerDriverManagement.CreatePrefect)
	// router.HandleFunc("/prefecture-driver-management/update", controllerDriverManagement.UpdatePrefect)
	router.HandleFunc("/prefecture-driver-management/{id_prefecture}", controllerDriverManagement.DeletePrefect)
	// end prefecture driver management
	
	// start store information driver management
	router.HandleFunc("/storeinformation-driver-management", controllerDriverManagement.ReturnAllStoreInformation)
	router.HandleFunc("/storeinformation-driver-management/filter", controllerDriverManagement.ReturnAllFilterInformation)
	router.HandleFunc("/storeinformation-driver-management/{page}/{perPage}", controllerDriverManagement.ReturnAllStoreInformationPagination)
	router.HandleFunc("/storeinformation-driver-management/get", controllerDriverManagement.GetStoreInformation)
	// router.HandleFunc("/storeinformation-driver-management/search/{store_name}/{code_store}", controllerDriverManagement.SearchStoreInformation)
	// router.HandleFunc("/storeinformation-driver-management/create", controllerDriverManagement.CreateStoreInformation)
	// router.HandleFunc("/storeinformation-driver-management/update", controllerDriverManagement.UpdateStoreInformation)
	// router.HandleFunc("/storeinformation-driver-management/{id_code_store}", controllerDriverManagement.DeleteStoreInformation)
	// end store information driver management

	// start store section information driver management
	router.HandleFunc("/store-section-information-driver-management", controllerDriverManagement.ReturnAllStroreSectionInformation)
	router.HandleFunc("/store-section-information-driver-management/{page}/{perPage}", controllerDriverManagement.ReturnAllStroreSectionInformationPagination)
	router.HandleFunc("/store-section-information-driver-management/get", controllerDriverManagement.GetStoreSectionInformation)
	// router.HandleFunc("/store-section-information-driver-management/create", controllerDriverManagement.CreateStoreSectionInformation)
	// router.HandleFunc("/store-section-information-driver-management/update", controllerDriverManagement.UpdateStoreSectionInformation)
	router.HandleFunc("/store-section-information-driver-management/{id_store_section}", controllerDriverManagement.DeleteStoreSectionInformation)
	// end store section information driver management 

	// start unit information driver management
	router.HandleFunc("/unit-information-driver-management", controllerDriverManagement.ReturnAllUnitInformation)
	router.HandleFunc("/unit-information-driver-management/{page}/{perPage}", controllerDriverManagement.ReturnAllUnitInformationPagination)
	router.HandleFunc("/unit-information-driver-management/get", controllerDriverManagement.GetUnitInformation)
	// router.HandleFunc("/unit-information-driver-management/create", controllerDriverManagement.CreateUnitInformation)
	// router.HandleFunc("/unit-information-driver-management/update", controllerDriverManagement.UpdateUnitInformation)
	router.HandleFunc("/unit-information-driver-management/{id_unit}", controllerDriverManagement.DeleteUnitInformation)
	// end unit information driver management

	// start department information 
	router.HandleFunc("/department-information-driver-management", controllerDriverManagement.ReturnAllDepartementInformation)
	router.HandleFunc("/department-information-driver-management/{page}/{perPage}", controllerDriverManagement.ReturnAllDepartementInformationPagination)
	router.HandleFunc("/department-information-driver-management/get", controllerDriverManagement.GetDepartementInformation)
	// router.HandleFunc("/department-information-driver-management/create", controllerDriverManagement.CreateDepartementInformation)
	// router.HandleFunc("/department-information-driver-management/update", controllerDriverManagement.UpdateDepartementInformation)
	router.HandleFunc("/department-information-driver-management/{id_department}", controllerDriverManagement.DeleteDepartementInformation)
	// end department information
	
	// end data driver management
// ==========================================================================================================================================================

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

	// start Commuting Transportation Application
	router.HandleFunc("/commuting-basic-information/{employee_number}", entertheinformation.ReturnCreateCommutingBasicInformation)
	router.HandleFunc("/commuting-basic-information-CheckData", entertheinformation.ReturnGetByCommutingBasicInformation)
	router.HandleFunc("/commuting-UsageRecord-CheckData", entertheinformation.ReturnGetByCommutingUsageRecord)
	router.HandleFunc("/commuting-UsageRecord-CheckDataForEdit", entertheinformation.ReturnGetByCommutingUsageRecordForEdit)
	router.HandleFunc("/commuting-UsageRecord-Apply/{condition}/{store_id}/{employee_id}", entertheinformation.ReturnInsertUsageRecordApplyForTravelExpenses)
	router.HandleFunc("/commuting-UsageRecord-Apply-Update", entertheinformation.ReturnUpdateUsageRecordApplyForTravelExpenses)
	router.HandleFunc("/commuting-UsageRecord-Delete/{id_commuting_trip}", entertheinformation.ReturnDeleteUsageRecord)
	router.HandleFunc("/commuting-UsageRecord-ShowUseMyRoute", entertheinformation.ReturnGetByCommutingUsageRecordUseMyRoute)
	router.HandleFunc("/commuting-UsageRecord-ShowHistory", entertheinformation.ReturnGetByCommutingUsageRecordHistory)
	router.HandleFunc("/commuting-UsageRecord-Draft/{id_commuting_trip}", entertheinformation.ReturnUpdateUsageRecordDraft)
	router.HandleFunc("/commuting-UseUsageRecord/{id_commuting_trip}/{date}", entertheinformation.ReturnUseUsageRecord)
	router.HandleFunc("/commuting-InputConfirmation-ShowDataById", entertheinformation.ReturnGetByCommutingInputConfirmation)
	router.HandleFunc("/commuting-InputConfirmation-Submit/{id_commuting_trip}", entertheinformation.ReturnSubmitInputConfirmation)
	// end Commuting Transportation Application

	// start Commuting Approve
	router.HandleFunc("/commuting-ApproveShowData", Approve.ReturnGetDataApproveCommutingSumByAllEmployeeCode)
	router.HandleFunc("/commuting-ApproveShowDataByEmployeeCode", Approve.ReturnGetDataApproveByCommutingEmployeeCode)
	router.HandleFunc("/commuting-ApproveShowDataByEmployeeCodeDetail", Approve.ReturnDetailCommutingByEmployeeCode)
	router.HandleFunc("/commuting-Approve/{employee_number}/{id_basic_information}/{code_commuting}", Approve.ReturnCommutingApproveOrReject)
	// end Commuting Approve

	// start master Data Transportation
	router.HandleFunc("/Get-MasterDataTransportation", entertheinformation.ReturnGetDataMasterTransportation)
	// end master Data Transportation

	// start list general recruitment
	router.HandleFunc("/Get-list-GR", list_GR.ReturnGetAllDataByStatus)
	router.HandleFunc("/Get-list-GR-detail", list_GR.ReturnGetAllDataDetailByStatus)
	// end list general recruitment


	fmt.Println("Connected to port 9000")

	handler := handlers.LoggingHandler(os.Stdout,router)
		corsHandle := cors.AllowAll().Handler(handler)
	http.ListenAndServe(":9000", corsHandle)

}
// install fresh golang
// export GOPATH=$HOME/go
// export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
// go get github.com/pilu/fresh
// fresh