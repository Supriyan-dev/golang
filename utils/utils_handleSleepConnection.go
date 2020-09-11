package utils

import (
	"database/sql"
)

func KillSleepConnection(con *sql.DB, ConnectionId string){
	con.Exec(`KILL QUERY ?`,ConnectionId)
}