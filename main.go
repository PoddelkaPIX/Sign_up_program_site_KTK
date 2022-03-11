package main

import (
	"advanced_training_site_KTK/ServerSettings"
	"advanced_training_site_KTK/db"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Programs []db.Program
	Setting  db.Passport_optins
}

var router *gin.Engine
var cfg = ServerSettings.Load("settings.cfg")

func main() {
	router = gin.Default()
	router.GET("/", Handler_home_page)
	router.GET("/add_token", Add_token)
	router.GET("/delete_token", Delete_token)
	router.GET("/passport/:Id", Handler_program_passport)
	router.GET("/registration/:Id_program", Handler_registration_form)
	router.GET("/control_panel", Handler_control_panel)
	router.POST("/login", Login)
	router.POST("/successfulRegistration", Successful_registration)
	router.POST("/previewing_program", Previewing_program)
	router.POST("/add_program", Add_program)
	router.POST("/confirm/listener", Confirm_listener)
	router.Static("/assets/", "public/")
	router.LoadHTMLGlob("templates/*.html")
	_ = router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func Handler_home_page(c *gin.Context) {
	var data Data
	e := db.Connect(cfg)
	defer db.Link.Close()
	if e != nil {
		fmt.Println(e)
		return
	}

	res := db.Select_all_programs()
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 { // Переворот массива чтобы сначало были неначатые программы
		res[i], res[j] = res[j], res[i]
	}
	settings := db.Select_program_settings()
	data.Programs = res
	data.Setting = settings
	c.HTML(200, "homePage.html", data)
}

func Handler_program_passport(c *gin.Context) {
	err := db.Connect(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Link.Close()
	id := c.Params.ByName("Id")
	c.HTML(200, "programPassport.html", db.Select_program_passport(id))
}

func Handler_registration_form(c *gin.Context) {
	err := db.Connect(cfg)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Link.Close()

	id := c.Params.ByName("Id_program")

	c.HTML(200, "registrationForm.html", db.Select_program_passport(id))
}

func Handler_control_panel(c *gin.Context) {
	var data Data
	e := db.Connect(cfg)
	defer db.Link.Close()
	if e != nil {
		fmt.Println(e)
		return
	}

	res := db.Select_all_programs()
	settings := db.Select_program_settings()
	data.Programs = res
	data.Setting = settings
	c.HTML(200, "controlPanel.html", data)

}

func Login(c *gin.Context) {
	type Input_data struct {
		Login    string `json:"Login"`
		Password string `json:"Password"`
	}
	var admin db.Admin
	var login_data Input_data

	c.BindJSON(&login_data)

	db.Connect(cfg)
	defer db.Link.Close()
	rows, e := db.Link.Query(`SELECT "Login", "Password" FROM "Admin"`)
	if e != nil {
		fmt.Println(e.Error())
	}
	for rows.Next() {
		e = rows.Scan(&admin.Login, &admin.Password)
		if e != nil {
			fmt.Println(e.Error())
		}
		fmt.Println(admin.Login, admin.Password)
		if (admin.Login == login_data.Login) && (admin.Password == login_data.Password) {
			Add_token(c)
			c.Redirect(http.StatusFound, "/")
		}
	}
}

func Successful_registration(c *gin.Context) {
	surname := c.PostForm("surname")
	name := c.PostForm("name")
	patronymic := c.PostForm("patronymic")
	telephone := c.PostForm("telephone")
	education := c.PostForm("education")
	program_id := c.PostForm("program_id")
	program_id_int, _ := strconv.ParseInt(program_id, 10, 64)
	email := c.PostForm("email")
	dt := time.Now()
	registration_date := dt
	birth_date := c.PostForm("birth_date")

	db.Connect(cfg)
	defer db.Link.Close()

	_, err := db.Link.Exec(`insert into "Listeners" 
	("Surname", 
	"Name", 
	"Patronymic",
	"Telephone", 
	"Education", 
	"Program_id", 
	"E-mail", 
	"Registration_date", 
	"Birth_date",
	"Listener_status") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		surname, name, patronymic, telephone, education, program_id_int, email, registration_date, birth_date, "Подал заявление")
	if err != nil {
		panic(err)
	}
	fmt.Println(registration_date)
	fmt.Println(birth_date)
	c.Redirect(http.StatusFound, "/")
}

func Previewing_program(c *gin.Context) {
	var program db.Program
	Type_program := strings.Split(c.PostForm("Type_program"), ";@")
	Title := c.PostForm("Title")
	Training_form := strings.Split(c.PostForm("Training_form"), ";@")
	Size_program := c.PostForm("Size_program")
	Length := c.PostForm("Length")
	Place := strings.Split(c.PostForm("Place"), ";@")
	Program_level := strings.Split(c.PostForm("Program_level"), ";@")
	Direction_study := strings.Split(c.PostForm("Direction_study"), ";@")
	Price := c.PostForm("Price")
	Minimum_group_size := c.PostForm("Minimum_group_size")
	Start_date := c.PostForm("Start_date")
	if Start_date == "" {
		Start_date = "По мере набора группы"
	}
	Issued_document := strings.Split(c.PostForm("Issued_document"), ";@")
	Requirements_listeners := strings.Split(c.PostForm("Requirements_listeners"), ";@")

	program.Program_level_id = Program_level[0]
	program.Type_program_id = Type_program[0]
	program.Training_form_id = Training_form[0]
	program.Place_id = Place[0]
	program.Direction_study_id = Direction_study[0]
	program.Issued_document_id = Issued_document[0]
	program.Requirements_listeners_id = Requirements_listeners[0]

	program.Title = Title
	program.Program_level = Program_level[1]
	program.Type_program = Type_program[1]
	program.Training_form = Training_form[1]
	program.Size_program = Size_program
	program.Length = Length
	program.Place = Place[1]
	program.Direction_study = Direction_study[1]
	program.Price = Price
	program.Minimum_group_size = Minimum_group_size
	program.Start_date = Start_date
	program.Issued_document = Issued_document[1]
	program.Requirements_listeners = Requirements_listeners[1]
	c.HTML(200, "previewingProgram.html", program)
}

func Add_program(c *gin.Context) {
	Type_program := c.PostForm("Type_program")
	Title := c.PostForm("Title")
	Training_form := c.PostForm("Training_form")
	Size_program := c.PostForm("Size_program")
	Length := c.PostForm("Length")
	Place := c.PostForm("Place")
	Program_level := c.PostForm("Program_level")
	Direction_study := c.PostForm("Direction_study")
	Price := c.PostForm("Price")
	Minimum_group_size := c.PostForm("Minimum_group_size")
	Start_date := c.PostForm("Start_date")
	Issued_document := c.PostForm("Issued_document")
	Requirements_listeners := c.PostForm("Requirements_listeners")
	db.Connect(cfg)
	defer db.Link.Close()

	_, e := db.Link.Exec(`INSERT INTO "Program_passport" 
					("Type_program_id", 
					"Title", 
					"Training_form_id", 
					"Size_program",
					"Length",
					"Place_id",
					"Program_level_id",
					"Direction_study_id",
					"Price",
					"Minimum_group_size",
					"Start_date", 
					"End_date",
					"Issued_document_id",
					"Requirements_listeners_id", 
					"Program_status") values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
		Type_program, Title, Training_form, Size_program, Length, Place,
		Program_level, Direction_study, Price, Minimum_group_size, Start_date,
		"-", Issued_document, Requirements_listeners, "Ожидание")
	if e != nil {
		fmt.Println(e)
		return
	}
	c.Redirect(http.StatusFound, "/control_panel")
}
func Add_token(c *gin.Context) {
	c.SetCookie("session_token", "Кук создан", 3600, "/", "localhost", false, false)
	cookie, _ := c.Cookie("session_token")
	fmt.Println(cookie)
	fmt.Println("Токен выдан")
	c.Redirect(http.StatusFound, "/")
}

func Delete_token(c *gin.Context) {
	c.SetCookie("session_token", "Кук создан", -1, "/", "localhost", false, false)
	fmt.Println("Токен удалён")
	c.Redirect(http.StatusFound, "/")
}

func Confirm_listener(c *gin.Context) {
	type input struct {
		Lidtener_id string `json:"Listener_id"`
	}
	i := input{}
	e := c.BindJSON(&i.Lidtener_id)
	if e != nil {
		c.JSON(400, nil)
		return
	}

	fmt.Println(i.Lidtener_id)
	c.JSON(200, i.Lidtener_id)
}
