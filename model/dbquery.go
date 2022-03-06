package model

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

func AddStudent(name, address, class, phone string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "student")
	data.Set("dbtype", "sqlite3")
	data.Set("name", name)
	data.Set("address", address)
	data.Set("class", class)
	data.Set("phone", phone)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

func ShowAll() []map[string]interface{} {
	qs := "SELECT * FROM student"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
