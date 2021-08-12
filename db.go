package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Database struct {
	file *os.File
	data map[string]interface{}
}

func NewDatabase(fileName string) (*Database, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	file.Seek(0, 0)
	if len(byteValue) == 0 {
		_, err = file.WriteString(`{}`)
		if err != nil {
			return nil, err
		}

		file.Seek(0, 0)
		byteValue, err = ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return &Database{file: file, data: data}, nil
}

func (db *Database) Close() error {
	return db.file.Close()
}

func (db *Database) Get(id string) (interface{}, bool) {
	value, ok := db.data[id]
	return value, ok
}

func (db *Database) Put(id string, value interface{}) error {
	if value == nil {
		_, ok := db.data[id]
		if ok {
			delete(db.data, id)
		}
		return nil
	}

	db.data[id] = value

	byteValue, err := json.Marshal(db.data)
	if err != nil {
		return err
	}

	err = db.file.Truncate(0)
	if err != nil {
		return err
	}

	db.file.Seek(0, 0)
	_, err = db.file.Write(byteValue)
	return err
}

func (db *Database) Print() {
	fmt.Println("Printing database:")

	sb := strings.Builder{}
	sb.WriteString("|")
	for key := range db.data {
		sb.WriteString(key)
		sb.WriteString("|")
	}

	fmt.Println(sb.String())
}
