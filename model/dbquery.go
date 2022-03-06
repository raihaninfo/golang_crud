package model

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

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

func ShowById(id string) []map[string]interface{} {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	qs := fmt.Sprintf("SELECT * FROM student WHERE id=%v;", idInt)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
