package ServerSettings

import (
	"encoding/json"
	"fmt"
	"os"
)

type Setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgName     string
}

func Load(filename string) *Setting {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Файл не открылся!")
		return nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	setting := new(Setting)
	e := decoder.Decode(&setting)
	if e != nil {
		return nil
	}
	return setting
}
