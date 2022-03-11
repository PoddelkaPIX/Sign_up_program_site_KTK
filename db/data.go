package db

import (
	"encoding/csv"
	"fmt"
	"os"

	"advanced_training_site_KTK/ServerSettings"
)

var cfg = ServerSettings.Load("settings.cfg")

func Open_csv() {
	e := Connect(cfg)
	if e != nil {
		fmt.Println(e)
	}
	defer Link.Close()

	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println("Ошибка", e)
			break
		}
		fmt.Println(record)
		var type_id int
		var form_id int
		var place_id int
		Link.QueryRow(`SELECT "Id" FROM "Type_program_list" WHERE "Type" = $1 `, record[0]).Scan(&type_id)
		Link.QueryRow(`SELECT "Id" FROM "Training_form_list" WHERE "Training_form" = $1 `, record[2]).Scan(&form_id)
		Link.QueryRow(`SELECT "Id" FROM "Place_list" WHERE "Place" = $1 `, record[6]).Scan(&place_id)

		fmt.Println(record)

		// _, err := db.Link.Exec(`insert into "Program_passport" ("Type_program_id", "Title", "Training_form_id", "Size_program", "Length", "Price", "Place_id", "Program_level_id", "Direction_study_id",
		// "Minimum_group_size", "Beginning_classes", "Issued_document_id", "Requirements_listeners_id") values ($1, $2, $3, $4, $5, $6, $7, "-", "-", "-", "-", "-", "-")`,
		// 	type_id, record[1], 1, record[3], string(record[4]), string(record[5]), 1)
		// if err != nil {
		// 	panic(err)
		// }
	}
}
