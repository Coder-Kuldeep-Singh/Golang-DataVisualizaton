package models

import (
	"Dashboard/conn"
	"log"
)

func UserIPAddress(ip string) {
	db, err := conn.InitDB()
	if err != nil {
		log.Println("Database Connection Failed :", err.Error())
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := `CREATE TABLE IF NOT EXISTS userip(
		id int NOT NULL AUTO_INCREMENT,
		ip varchar(255) NOT NULL,
		PRIMARY KEY (id));`
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println("Statement have problem :", err.Error())
		panic(err.Error())
		return
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error to create table :", err.Error())
		return
	}
	log.Println("userip table has created!")

	InsertStatement := `INSERT INTO userip(
		ip) VALUES(?)`
	result, err := db.Exec(InsertStatement, ip)
	if err != nil {
		log.Println("Error to Execute Insert Statement :", err.Error())
		return
	}
	log.Println(result.LastInsertId())

}
