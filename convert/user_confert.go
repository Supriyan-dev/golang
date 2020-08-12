package convert

import (
	"log"

	"github.com/jeffri/golang-test/db"
)

func userConfert(sql string, id_user string) int {
	var dataInt int

	eksekusiInteger, errEksekusiInt := db.Connect()
	eksekusiInteger.QueryRow(&sql, &id_user)
	if eksekusiInteger != nil {
		log.Println(eksekusiInteger)
	}

	return dataInt, nil

}
