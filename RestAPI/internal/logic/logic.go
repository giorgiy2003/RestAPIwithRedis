package Logic

import (
	"encoding/json"
	"fmt"
	"log"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
)

func Create(p Model.Person) error {
	personBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = Repository.Rdb.Set(fmt.Sprint(p.Id), personBytes, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func Read() ([]Model.Person, error) {
	var personInfo = []Model.Person{}

	keys, err := Repository.Rdb.Keys("*").Result()
	if err != nil {
		return personInfo, err
	}
	for _, key := range keys {
		person, err := ReadOne(fmt.Sprint(key))
		if err != nil {
			log.Println(err)
		}
		personInfo = append(personInfo, person)
	}

	return personInfo, nil
}

func ReadOne(id string) (Model.Person, error) {
	var person Model.Person

	value, err := Repository.Rdb.Get(id).Result()
	if err != nil {
		return person, err
	}
	err = json.Unmarshal([]byte(value), &person)
	if err != nil {
		return person, err
	}
	return person, nil
}

func Delete(id string) error {
	value , err := Repository.Rdb.Del(id).Result()
	if err != nil {
		return err
	}
	if value == 0 {
		return fmt.Errorf("Записи с id = %s не существует", id)
	}
	return nil
}
