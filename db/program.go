package db

import (
	"fmt"
)

type Passport_optins struct {
	Type_program_list           []Type_program
	Training_form_list          []Training_form
	Place_list                  []Place
	Program_level_list          []Program_level
	Direction_study_list        []Direction_study
	Issued_document_list        []Issued_document
	Requirements_listeners_list []Requirements_listeners
}

type Program struct {
	Id                        int
	Type_program              string
	Type_program_id           string
	Title                     string
	Training_form             string
	Training_form_id          string
	Size_program              string
	Length                    string
	Place                     string
	Place_id                  string
	Program_level             string
	Program_level_id          string
	Direction_study           string
	Direction_study_id        string
	Price                     string
	Minimum_group_size        string
	Start_date                string
	End_date                  string
	Issued_document           string
	Issued_document_id        string
	Requirements_listeners    string
	Requirements_listeners_id string
	Program_status            string
	Listeners                 []Listener
}

type Listener struct {
	Id                int
	Surname           string
	Name              string
	Patronymic        string
	Telephone         string
	Education         string
	Program_id        string
	Email             string
	Registration_date string
	Birth_date        string
	Listener_status   string
}

type Type_program struct {
	Type_program    string
	Type_program_id string
}

type Training_form struct {
	Training_form    string
	Training_form_id string
}
type Place struct {
	Place    string
	Place_id string
}
type Program_level struct {
	Program_level    string
	Program_level_id string
}
type Direction_study struct {
	Direction_study    string
	Direction_study_id string
}
type Issued_document struct {
	Issued_document    string
	Issued_document_id string
}
type Requirements_listeners struct {
	Requirements_listeners    string
	Requirements_listeners_id string
}

func Select_all_programs() []Program {
	var program Program
	listeners_list := make([]Listener, 0)
	ProgramsList := make([]Program, 0)

	rows, e := Link.Query(`SELECT 
						"Program_passport"."Id", 
						"Program_passport"."Title", 
						"Training_form_list"."Training_form", 
						"Program_passport"."Training_form_id", 
						"Program_passport"."Size_program", 
						"Program_passport"."Length",
						"Place_list"."Place", 
						"Program_passport"."Place_id", 
						"Type_program_list"."Type", 
						"Program_passport"."Type_program_id",
						"Program_level_list"."Level", 
						"Program_passport"."Program_level_id",
						"Direction_study_list"."Direction", 
						"Program_passport"."Direction_study_id", 
						"Program_passport"."Price", 
						"Program_passport"."Minimum_group_size", 
						"Program_passport"."Start_date",
						"Program_passport"."End_date",
						"Issued_document_list"."Document", 
						"Program_passport"."Issued_document_id",
						"Requirements_listeners_list"."Requirements", 
						"Program_passport"."Requirements_listeners_id", 
						"Program_passport"."Program_status" 
						FROM "Program_passport"
						JOIN "Place_list" ON "Place_list"."Id" = "Program_passport"."Place_id"
						JOIN "Training_form_list" ON "Training_form_list"."Id" = "Program_passport"."Training_form_id"
						JOIN "Type_program_list" ON "Type_program_list"."Id" = "Program_passport"."Type_program_id"
						JOIN "Program_level_list" ON "Program_level_list"."Id" = "Program_passport"."Program_level_id"
						JOIN "Direction_study_list" ON "Direction_study_list"."Id" = "Program_passport"."Direction_study_id"
						JOIN "Issued_document_list" ON "Issued_document_list"."Id" = "Program_passport"."Issued_document_id"
						JOIN "Requirements_listeners_list" ON "Requirements_listeners_list"."Id" = "Program_passport"."Requirements_listeners_id"
						order by "Program_status" ASC, "Title" ASC`)

	if e != nil {
		fmt.Println("Ошибка в запросе в базу данных с выбором всех программ")
		fmt.Println(e)
		return nil
	}

	for rows.Next() {
		var listener Listener
		e = rows.Scan(&program.Id, &program.Title, &program.Training_form, &program.Training_form_id, &program.Size_program,
			&program.Length, &program.Place, &program.Place_id, &program.Type_program, &program.Type_program_id,
			&program.Program_level, &program.Program_level_id, &program.Direction_study, &program.Direction_study_id,
			&program.Price, &program.Minimum_group_size,
			&program.Start_date, &program.End_date, &program.Issued_document, &program.Issued_document_id,
			&program.Requirements_listeners, &program.Requirements_listeners_id, &program.Program_status)
		if e != nil {
			fmt.Println(e)
			return nil
		}
		rows_listeners, e := Link.Query(`SELECT 
						"Listeners"."Id",
						"Listeners"."Surname",
						"Listeners"."Name",
						"Listeners"."Patronymic",
						"Listeners"."Telephone",
						"Listeners"."Education",
						"Listeners"."Program_id",
						"Listeners"."E-mail",
						"Listeners"."Registration_date",
						"Listeners"."Birth_date",
						"Listeners"."Listener_status"
						FROM "Listeners"
						WHERE "Program_id" = $1
						ORDER BY "Name"`, &program.Id)

		if e != nil {
			fmt.Println("Ошибка в запросе в базу данных с выбором слушателей во всех слушателей")
			fmt.Println(e)
		}

		for rows_listeners.Next() {
			e := rows_listeners.Scan(&listener.Id, &listener.Surname, &listener.Name, &listener.Patronymic,
				&listener.Telephone, &listener.Education, &listener.Program_id,
				&listener.Email, &listener.Registration_date, &listener.Birth_date, &listener.Listener_status)
			if e != nil {
				fmt.Println(e)
			}

			listeners_list = append(listeners_list, Listener{
				Id:                listener.Id,
				Surname:           listener.Surname,
				Name:              listener.Name,
				Patronymic:        listener.Patronymic,
				Telephone:         listener.Telephone,
				Education:         listener.Education,
				Program_id:        listener.Program_id,
				Email:             listener.Email,
				Registration_date: listener.Registration_date,
				Birth_date:        listener.Birth_date,
				Listener_status:   listener.Listener_status,
			})
		}
		ProgramsList = append(ProgramsList, Program{
			Id:                        program.Id,
			Type_program:              program.Type_program,
			Type_program_id:           program.Type_program_id,
			Title:                     program.Title,
			Training_form:             program.Training_form,
			Training_form_id:          program.Training_form_id,
			Size_program:              program.Size_program,
			Length:                    program.Length,
			Place:                     program.Place,
			Place_id:                  program.Place_id,
			Program_level:             program.Program_level,
			Program_level_id:          program.Program_level_id,
			Direction_study:           program.Direction_study,
			Direction_study_id:        program.Direction_study_id,
			Price:                     program.Price,
			Minimum_group_size:        program.Minimum_group_size,
			Start_date:                program.Start_date,
			End_date:                  program.End_date,
			Issued_document:           program.Issued_document,
			Issued_document_id:        program.Issued_document_id,
			Requirements_listeners:    program.Requirements_listeners,
			Requirements_listeners_id: program.Requirements_listeners_id,
			Program_status:            program.Program_status,
			Listeners:                 listeners_list,
		})
		listeners_list = []Listener{}
	}
	return ProgramsList
}

func Select_program_passport(program_id string) Program {
	var program Program
	var listener Listener
	listeners_list := make([]Listener, 0)

	row := Link.QueryRow(`SELECT 
						"Program_passport"."Id", 
						"Program_passport"."Title", 
						"Training_form_list"."Training_form", 
						"Program_passport"."Training_form_id", 
						"Program_passport"."Size_program", 
						"Program_passport"."Length",
						"Place_list"."Place", 
						"Program_passport"."Place_id", 
						"Type_program_list"."Type", 
						"Program_passport"."Type_program_id",
						"Program_level_list"."Level", 
						"Program_passport"."Program_level_id",
						"Direction_study_list"."Direction", 
						"Program_passport"."Direction_study_id", 
						"Program_passport"."Price", 
						"Program_passport"."Minimum_group_size", 
						"Program_passport"."Start_date",
						"Program_passport"."End_date",
						"Issued_document_list"."Document", 
						"Program_passport"."Issued_document_id",
						"Requirements_listeners_list"."Requirements", 
						"Program_passport"."Requirements_listeners_id", 
						"Program_passport"."Program_status" 
						FROM "Program_passport"
						JOIN "Place_list" ON "Place_list"."Id" = "Program_passport"."Place_id"
						JOIN "Training_form_list" ON "Training_form_list"."Id" = "Program_passport"."Training_form_id"
						JOIN "Type_program_list" ON "Type_program_list"."Id" = "Program_passport"."Type_program_id"
						JOIN "Program_level_list" ON "Program_level_list"."Id" = "Program_passport"."Program_level_id"
						JOIN "Direction_study_list" ON "Direction_study_list"."Id" = "Program_passport"."Direction_study_id"
						JOIN "Issued_document_list" ON "Issued_document_list"."Id" = "Program_passport"."Issued_document_id"
						JOIN "Requirements_listeners_list" ON "Requirements_listeners_list"."Id" = "Program_passport"."Requirements_listeners_id" WHERE "Program_passport"."Id" = $1`, program_id)
	e := row.Scan(&program.Id, &program.Title, &program.Training_form, &program.Training_form_id, &program.Size_program,
		&program.Length, &program.Place, &program.Place_id, &program.Type_program, &program.Type_program_id,
		&program.Program_level, &program.Program_level_id, &program.Direction_study, &program.Direction_study_id,
		&program.Price, &program.Minimum_group_size,
		&program.Start_date, &program.End_date, &program.Issued_document, &program.Issued_document_id,
		&program.Requirements_listeners, &program.Requirements_listeners_id, &program.Program_status)
	if e != nil {
		fmt.Println("Ошибка при записи данных в экземпляр структуры Program")
		fmt.Println(e)
		return Program{}
	}
	rows_listeners, e := Link.Query(`SELECT 
						"Listeners"."Id",
						"Listeners"."Surname",
						"Listeners"."Name",
						"Listeners"."Patronymic",
						"Listeners"."Telephone",
						"Listeners"."Education",
						"Listeners"."Program_id",
						"Listeners"."E-mail",
						"Listeners"."Registration_date",
						"Listeners"."Birth_date",
						"Listeners"."Listener_status"
						FROM "Listeners"
						WHERE "Program_id" = $1`, &program.Id)

	if e != nil {
		fmt.Println("Ошибка в запросе в базу данных с выбором слушателей во всех программах")
		fmt.Println(e)
	}

	for rows_listeners.Next() {
		e := rows_listeners.Scan(&listener.Id, &listener.Surname, &listener.Name, &listener.Patronymic,
			&listener.Telephone, &listener.Education, &listener.Program_id,
			&listener.Email, &listener.Registration_date, &listener.Birth_date, &listener.Listener_status)
		if e != nil {
			fmt.Println("Ошибка при записи данных в экземпляр структуры Listener")
			fmt.Println(e)
		}

		listeners_list = append(listeners_list, Listener{
			Id:                listener.Id,
			Surname:           listener.Surname,
			Name:              listener.Name,
			Patronymic:        listener.Patronymic,
			Telephone:         listener.Telephone,
			Education:         listener.Education,
			Program_id:        listener.Program_id,
			Email:             listener.Email,
			Registration_date: listener.Registration_date,
			Birth_date:        listener.Birth_date,
			Listener_status:   listener.Listener_status,
		})
	}
	program.Listeners = listeners_list

	fmt.Println("Program", program)
	return program
}

func Select_program_settings() Passport_optins {
	var type_program Type_program
	var training_form Training_form
	var place Place
	var program_level Program_level
	var direction_study Direction_study
	var issued_document Issued_document
	var requirements_listeners Requirements_listeners
	var type_program_list []Type_program
	var training_form_list []Training_form
	var place_list []Place
	var program_level_list []Program_level
	var direction_study_list []Direction_study
	var issued_document_list []Issued_document
	var requirements_listeners_list []Requirements_listeners
	var data Passport_optins

	rows, e := Link.Query(`SELECT "Id", "Type" FROM "Type_program_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех типов программ.")
		type_program_list = nil
	}
	for rows.Next() {
		rows.Scan(&type_program.Type_program_id, &type_program.Type_program)
		type_program_list = append(type_program_list, type_program)
	}

	rows, e = Link.Query(`SELECT "Id", "Training_form" FROM "Training_form_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех форм программ.")
		training_form_list = nil
	}
	for rows.Next() {
		rows.Scan(&training_form.Training_form_id, &training_form.Training_form)
		training_form_list = append(training_form_list, training_form)
	}

	rows, e = Link.Query(`SELECT "Id", "Place" FROM "Place_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех мест программ.")
		place_list = nil
	}
	for rows.Next() {
		rows.Scan(&place.Place_id, &place.Place)
		place_list = append(place_list, place)
	}

	rows, e = Link.Query(`SELECT "Id", "Level" FROM "Program_level_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех уровней программ.")
		program_level_list = nil
	}
	for rows.Next() {
		rows.Scan(&program_level.Program_level_id, &program_level.Program_level)
		program_level_list = append(program_level_list, program_level)
	}

	rows, e = Link.Query(`SELECT "Id", "Direction" FROM "Direction_study_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех направлений программ.")
		direction_study_list = nil
	}
	for rows.Next() {
		rows.Scan(&direction_study.Direction_study_id, &direction_study.Direction_study)
		direction_study_list = append(direction_study_list, direction_study)
	}

	rows, e = Link.Query(`SELECT "Id", "Document" FROM "Issued_document_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех типов программ.")
		issued_document_list = nil
	}
	for rows.Next() {
		rows.Scan(&issued_document.Issued_document_id, &issued_document.Issued_document)
		issued_document_list = append(issued_document_list, issued_document)
	}

	rows, e = Link.Query(`SELECT "Id", "Requirements" FROM "Requirements_listeners_list"`)
	if e != nil {
		fmt.Println("Ошибка запроса на выбор всех типов программ.")
		requirements_listeners_list = nil
	}
	for rows.Next() {
		rows.Scan(&requirements_listeners.Requirements_listeners_id, &requirements_listeners.Requirements_listeners)
		requirements_listeners_list = append(requirements_listeners_list, requirements_listeners)
	}

	data.Type_program_list = type_program_list
	data.Training_form_list = training_form_list
	data.Place_list = place_list
	data.Program_level_list = program_level_list
	data.Direction_study_list = direction_study_list
	data.Issued_document_list = issued_document_list
	data.Requirements_listeners_list = requirements_listeners_list

	return data
}
