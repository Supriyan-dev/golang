package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/jeffri/golang-test/controller/"
	controller "github.com/jeffri/golang-test/controller/data_master_controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/storeinformation", controller.ReturnAllStoreInformation).Methods("GET")
	router.HandleFunc("/storeinformation?limit=10&offset=20", controller.ReturnAllStoreInformationPagination).Methods("GET")

	router.HandleFunc("/departement-information", controller.ReturnAllDepartementInformation).Methods("GET")
	router.HandleFunc("/departement-information/{page}", controller.ReturnAllDepartementInformationPagination).Methods("GET")

	router.HandleFunc("/store-section-information", controller.ReturnAllStroreSectionInformation).Methods("GET")
	router.HandleFunc("/store-section-information/{page}", controller.ReturnAllStroreSectionInformationPagination).Methods("GET")

	router.HandleFunc("/unit-information", controller.ReturnAllUnitInformation).Methods("GET")
	router.HandleFunc("/unit-information/{page}", controller.ReturnAllUnitInformationPagination).Methods("GET")

	router.HandleFunc("/prefecture", controller.ReturnAllPrefect).Methods("GET")
	router.HandleFunc("/prefecture/{page}", controller.ReturnAllPrefectPagination).Methods("GET")

	router.HandleFunc("/bank", controller.ReturnAllBank).Methods("GET")
	router.HandleFunc("/bank/{page}", controller.ReturnAllBankPagination).Methods("GET")

	router.HandleFunc("/full-time-salary", controller.ReturnAllFullTimeSalary).Methods("GET")
	router.HandleFunc("/full-time-salary/{page}", controller.ReturnAllFullTimeSalaryPagination).Methods("GET")

	router.HandleFunc("/part-time-above-18-salary", controller.ReturnAllPartTimeAbove18Salary).Methods("GET")
	router.HandleFunc("/part-time-above-18-salary/{page}", controller.ReturnAllPartTimeAbove18SalaryPagination).Methods("GET")

	router.HandleFunc("/part-time-under-18-salary", controller.ReturnAllPartTimeUnder18Salary).Methods("GET")
	router.HandleFunc("/part-time-under-18-salary/{page}", controller.ReturnAllPartTimeUnder18SalaryPagination).Methods("GET")

	router.HandleFunc("/user", controller.ReturnAllUser).Methods("GET")
	router.HandleFunc("/user/{page}", controller.ReturnAllUserPagination).Methods("GET")

	router.HandleFunc("/exp-category", controller.ReturnAllExpCategory).Methods("GET")
	router.HandleFunc("/exp-category/{page}/", controller.ReturnAllExpCategoryPagination).Methods("GET")

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
