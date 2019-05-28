package database

import (
	"chat-golang/server/model"
	"database/sql"
	"time"
)

func WriteToDb(db *sql.DB, msg *model.Message) {

	time := getTime()

	stmt, err := db.Prepare("insert into chat values(?,?,?,?)")
	checkErr(err)

	stmt.Exec(time, msg.Name, msg.Message, "basegroup")

}

func ReadMessageFromDb(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM chat")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func getTime() string {
	t := time.Now()
	tf := t.Format("2006-01-02 15:04:05")
	return tf
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
