package main

import (
	"fmt"
	"log"
	"net/http"

	controller "./controller/data_master_controller"
	controller "github.com/jeffri/golang-test/GO_DX_SERVICES/controller/data_master_controller"

	// controller "github.com/jeffri/golang-test/GO_DX_SERVICES/controller/list_input_information"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	// start data master
	// start crud store information
	router.HandleFunc("/storeinformation", controller.ReturnAllStoreInformation).Methods("GET")
	router.HandleFunc("/storeinformation/{page}/{perPage}", controller.ReturnAllStoreInformationPagination).Methods("GET")
	router.HandleFunc("/storeinformation/{id_code_store}", controller.GetStoreInformation).Methods("GET")
	router.HandleFunc("/storeinformation", controller.CreateStoreInformation).Methods("POST")
	router.HandleFunc("/storeinformation", controller.UpdateStoreInformation).Methods("PUT")
	router.HandleFunc("/storeinformation/{id_code_store}", controller.DeleteStoreInformation).Methods("DELETE")
	// end crud store information

	// start crud departement information
	router.HandleFunc("/departement-information", controller.ReturnAllDepartementInformation).Methods("GET")
	router.HandleFunc("/departement-information/{page}/{perPage}", controller.ReturnAllDepartementInformationPagination).Methods("GET")
	router.HandleFunc("/departement-information/{id_department}", controller.GetDepartementInformation).Methods("GET")
	router.HandleFunc("/departement-information", controller.CreateDepartementInformation).Methods("POST")
	router.HandleFunc("/departement-information", controller.UpdateDepartementInformation).Methods("PUT")
	router.HandleFunc("/departement-information/{id_department}", controller.DeleteDepartementInformation).Methods("DELETE")
	// end crud deaprtemen information

	// start crud srtore section information
	router.HandleFunc("/store-section-information", controller.ReturnAllStroreSectionInformation).Methods("GET")
	router.HandleFunc("/store-section-information/{page}/{perPage}", controller.ReturnAllStroreSectionInformationPagination).Methods("GET")
	router.HandleFunc("/store-section-information/{id_store_section}", controller.GetStoreSectionInformation).Methods("GET")
	router.HandleFunc("/store-section-information", controller.CreateStoreSectionInformation).Methods("POST")
	router.HandleFunc("/store-section-information", controller.UpdateStoreSectionInformation).Methods("PUT")
	router.HandleFunc("/store-section-information/{id_store_section}", controller.DeleteStoreSectionInformation).Methods("DELETE")
	//end crud store section infomration

	// start crud unit information
	router.HandleFunc("/unit-information", controller.ReturnAllUnitInformation).Methods("GET")
	router.HandleFunc("/unit-information/{page}/{perPage}", controller.ReturnAllUnitInformationPagination).Methods("GET")
	router.HandleFunc("/unit-information/{id_unit}", controller.GetUnitInformation).Methods("GET")
	router.HandleFunc("/unit-information", controller.CreateUnitInformation).Methods("POST")
	router.HandleFunc("/unit-information", controller.UpdateUnitInformation).Methods("PUT")
	router.HandleFunc("/unit-information/{id_unit}", controller.DeleteUnitInformation).Methods("DELETE")
	// end crud unit information

	// start crud prefecture
	router.HandleFunc("/prefecture", controller.ReturnAllPrefect).Methods("GET")
	router.HandleFunc("/prefecture/{page}/{perPage}", controller.ReturnAllPrefectPagination).Methods("GET")
	router.HandleFunc("/prefecture/{id_prefecture}", controller.GetPrefect).Methods("GET")
	router.HandleFunc("/prefecture", controller.CreatePrefect).Methods("POST")
	router.HandleFunc("/prefecture", controller.UpdatePrefect).Methods("PUT")
	router.HandleFunc("/prefecture/{id_prefecture}", controller.DeletePrefect).Methods("DELETE")
	// end crud prefecture

	// start crud bank
	router.HandleFunc("/bank", controller.ReturnAllBank).Methods("GET")
	router.HandleFunc("/bank/{page}/{perPage}", controller.ReturnAllBankPagination).Methods("GET")
	router.HandleFunc("/bank/{id_bank}", controller.GetBank).Methods("GET")
	router.HandleFunc("/bank", controller.CreateBank).Methods("POST")
	router.HandleFunc("/bank", controller.UpdateBank).Methods("PUT")
	router.HandleFunc("/bank/{id_bank}", controller.DeleteBank).Methods("DELETE")
	// end crud bank

	// start crud full time salary
	router.HandleFunc("/full-time-salary", controller.ReturnAllFullTimeSalary).Methods("GET")
	router.HandleFunc("/full-time-salary/{page}/{perPage}", controller.ReturnAllFullTimeSalaryPagination).Methods("GET")
	router.HandleFunc("/full-time-salary/{id_full_time_salary}", controller.GetFullTimeSalary).Methods("GET")
	router.HandleFunc("/full-time-salary", controller.CreateFullTimeSalary).Methods("POST")
	router.HandleFunc("/full-time-salary", controller.UpdateFullTimeSalary).Methods("PUT")
	router.HandleFunc("/full-time-salary/{id_full_time_salary}", controller.DeleteFullTimeSalary).Methods("DELETE")
	// end crud full time salary

	// start crud part time salary
	router.HandleFunc("/part-time-above-18-salary", controller.ReturnAllPartTimeAbove18Salary).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/{page}/{perPage}", controller.ReturnAllPartTimeAbove18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/{id_part_time_above_18_salary}", controller.GetPartTimeAbove18Salary).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary", controller.CreatePartTimeAbove18Salary).Methods("POST")
	router.HandleFunc("/part-time-above-18-salary", controller.UpdatePartTimeAbove18Salary).Methods("PUT")
	router.HandleFunc("/part-time-above-18-salary/{id_part_time_above_18_salary}", controller.DeletePartTimeAbove18Salary).Methods("DELETE")
	// end crud part time salary

	// start crud under 18 salary
	router.HandleFunc("/part-time-under-18-salary", controller.ReturnAllPartTimeUnder18Salary).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/{page}/{perPage}", controller.ReturnAllPartTimeUnder18SalaryPagination).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/{id_part_time_under_18_salary}", controller.GetPartTimeUnder18Salary).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary", controller.CreatePartTimeUnder18Salary).Methods("POST")
	router.HandleFunc("/part-time-under-18-salary", controller.UpdatePartTimeUnder18Salary).Methods("PUT")
	router.HandleFunc("/part-time-under-18-salary/{id_part_time_under_18_salary}", controller.DeletePartTimeUnder18Salary).Methods("DELETE")
	// end crud under 18 salary

	// start crud user
	router.HandleFunc("/user", controller.ReturnAllUser).Methods("GET")
	router.HandleFunc("/user/{page}/{perPage}", controller.ReturnAllUserPagination).Methods("GET")
	router.HandleFunc("/user/{id_user}", controller.GetUser).Methods("GET")
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id_user}", controller.DeleteUser).Methods("DELETE")
	// end crud user

	// start crud exp category
	router.HandleFunc("/exp-category", controller.ReturnAllExpCategory).Methods("GET")
	router.HandleFunc("/exp-category/{page}/{perPage}", controller.ReturnAllExpCategoryPagination).Methods("GET")
	router.HandleFunc("/exp-category/{id_exp}", controller.GetExpCategory).Methods("GET")
	router.HandleFunc("/exp-category", controller.CreateExpCategory).Methods("POST")
	router.HandleFunc("/exp-category", controller.UpdateExpCategory).Methods("PUT")
	router.HandleFunc("/exp-category/{id_exp}", controller.DeleteExpCategory).Methods("DELETE")
	// end crud exp category

	// Start permission to drive
	// router.HandleFunc("/test", controller.permissionToDrive).Methods("GET")

	// end permission to drive

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

	http.Handle("/", router)
	fmt.Println("Connected to port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))

}
