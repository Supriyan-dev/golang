package models

// func StoreInfor(r *http.Request) (arrStoreInformation []initialize.StoreInformation, err error) {
// 	var storeInformation initialize.StoreInformation
// 	var response initialize.Response

// 	code := mux.Vars(r)
// 	page := code["page"]
// 	newpage, err := strconv.Atoi(page)
// 	var dataTotal int = 100
// 	dataPage := 10

// 	totalPage := math.Ceil(float64(dataTotal) / float64(dataPage))

// 	firstPage := (dataPage * newpage) - dataPage

// 	fmt.Println(totalPage)

// 	db, err := db.Connect()
// 	rows, err := db.Query("SELECT * FROM store_information ORDER BY id_code_store LIMIT", firstPage+dataPage)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer db.Close()

// 	for rows.Next() {
// 		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
// 			log.Fatal(err.Error())

// 		} else {
// 			arrStoreInformation = append(arrStoreInformation, storeInformation)
// 		}
// 	}

// 	response.Status = 200
// 	response.Message = "Success"
// 	response.Data = arrStoreInformation

// 	return arrStoreInformation, nil
// }

// func Getinfo() (test []initialize.StoreInformation, err error) {
// 	var storeInformation initialize.StoreInformation

// 	db, err := db.Connect()
// 	rows, err := db.Query("SELECT * FROM store_information")

// 	// var test []initialize.StoreInformation

// 	for rows.Next() {
// 		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
// 			log.Fatal(err.Error())

// 		} else {
// 			test = append(test, storeInformation)
// 		}
// 	}

// 	return nil, test
// }

// func ReturnAllStoreInformation() {
// 	var storeInformation initialize.StoreInformation
// 	var arrStoreInformation []initialize.StoreInformation
// 	// var response initialize.Response

// 	db, err := db.Connect()

// 	rows, err := db.Query("SELECT * FROM store_information")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer db.Close()

// 	for rows.Next() {
// 		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
// 			log.Fatal(err.Error())

// 		} else {
// 			arrStoreInformation = append(arrStoreInformation, storeInformation)
// 		}
// 	}
// 	// response.Status = 200
// 	// response.Message = "Success"
// 	// response.Data = arrStoreInformation

// 	return arrStoreInformation

// }

// func StoreInformationPagination() int {
// 	db, err := db.Connect()
// 	defer db.Close()
// 	var count int
// 	db.Connect(&store{}).Count(&count)
// 	return count
// }
