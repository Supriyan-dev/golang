package list_general_recruitment

import (
	db2 "../../db"
	"strconv"
)

func GetDataWorkingTime(idemployeeType int ,employeeType string, id_basic_information int ) (start_working string,end_working string) {
	db := db2.Connect()

	query := ``
	if employeeType == `34` {
		query = `SELECT COALESCE(full_time_employee.start_working,''), COALESCE(full_time_employee.end_working,'') 
            FROM employment_type JOIN full_time_employee 
            ON employment_type.id_employment_type = full_time_employee.id_employment_type 
            JOIN full_time_salary ON full_time_salary.id_full_time_salary = full_time_salary.id_full_time_salary 
            WHERE employment_type.id_basic_information =` +strconv.Itoa(id_basic_information) +` and employment_type.id_employment_type = ` +strconv.Itoa(idemployeeType)
	}else if employeeType == `41` || employeeType == `43`{
		query = `SELECT COALESCE(part_time_above_18_employee.start_working,''), COALESCE(part_time_above_18_employee.end_working,'')  
            FROM employment_type JOIN part_time_above_18_employee 
            ON employment_type.id_employment_type = part_time_above_18_employee.id_employment_type
            JOIN part_time_above_18_salary ON part_time_above_18_employee.id_part_time_above_18_salary =
			part_time_above_18_salary.id_part_time_above_18_salary
            WHERE employment_type.id_basic_information =`+strconv.Itoa(id_basic_information) +` and employment_type.id_employment_type = ` +strconv.Itoa(idemployeeType)
	}else if employeeType == `42` {
		query = `SELECT COALESCE(part_time_under_18_employee.start_working,''), COALESCE(part_time_under_18_employee.end_working,'') 
            FROM employment_type JOIN part_time_under_18_employee 
            ON employment_type.id_employment_type = part_time_under_18_employee.id_employment_type
            JOIN part_time_under_18_salary ON part_time_under_18_employee.id_part_time_under_18_salary = part_time_under_18_salary.id_part_time_under_18_salary
            WHERE employment_type.id_basic_information = `+strconv.Itoa(id_basic_information) +` and employment_type.id_employment_type = ` +strconv.Itoa(idemployeeType)
	}
	errQR := db.QueryRow(query).Scan(&start_working,&end_working)
	if errQR != nil {
		//log.Println(errQR.Error())
	}
	return start_working, end_working
}

func GetDataDetailWorkingTime(employeeType string, id_basic_information int ) (salary int, EmployentInsuranceNo string,PensionNo string, start_working string,end_working string) {
	db := db2.Connect()

	query := ``
	if employeeType == `34` {
		query = `SELECT COALESCE(full_time_salary.salary,0), COALESCE(full_time_employee.employee_insurance_no,''),
			COALESCE(full_time_employee.pension_no,''),
			COALESCE(full_time_employee.start_working,''), COALESCE(full_time_employee.end_working,'') 
            FROM employment_type JOIN full_time_employee 
            ON employment_type.id_employment_type = full_time_employee.id_employment_type 
            JOIN full_time_salary ON full_time_salary.id_full_time_salary = full_time_salary.id_full_time_salary 
            WHERE employment_type.id_basic_information =` +strconv.Itoa(id_basic_information)
	}else if employeeType == `41` || employeeType == `43`{
		query = `SELECT  COALESCE(part_time_above_18_salary.salary,0), COALESCE(part_time_above_18_employee.employee_insurance_no,''),
			COALESCE(part_time_above_18_employee.pension_no,''),
			COALESCE(part_time_above_18_employee.start_working,''), COALESCE(part_time_above_18_employee.end_working,'')  
            FROM employment_type JOIN part_time_above_18_employee 
            ON employment_type.id_employment_type = part_time_above_18_employee.id_employment_type
            JOIN part_time_above_18_salary ON part_time_above_18_employee.id_part_time_above_18_salary =
			part_time_above_18_salary.id_part_time_above_18_salary
            WHERE employment_type.id_basic_information =`+strconv.Itoa(id_basic_information)
	}else if employeeType == `42` {
		query = `SELECT  COALESCE(part_time_under_18_salary.salary,0), COALESCE(part_time_under_18_employee.employee_insurance_no,''),
			COALESCE(part_time_under_18_employee.pension_no,''),
			COALESCE(part_time_under_18_employee.start_working,''), COALESCE(part_time_under_18_employee.end_working,'') 
            FROM employment_type JOIN part_time_under_18_employee 
            ON employment_type.id_employment_type = part_time_under_18_employee.id_employment_type
            JOIN part_time_under_18_salary ON part_time_under_18_employee.id_part_time_under_18_salary = part_time_under_18_salary.id_part_time_under_18_salary
            WHERE employment_type.id_basic_information = `+strconv.Itoa(id_basic_information)
	}
	errQR := db.QueryRow(query).Scan(&salary,&EmployentInsuranceNo,&PensionNo,&start_working,&end_working)
	if errQR != nil {
		//log.Println(errQR.Error())
	}
	return salary, EmployentInsuranceNo, PensionNo , start_working , end_working
}