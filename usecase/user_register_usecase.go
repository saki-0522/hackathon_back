package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	"log"
)

func Validation() error {
	if model.RegisterData.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if len(model.RegisterData.Name) > 50 {
		return fmt.Errorf("name is too long")
	}

	if len(model.RegisterData.Email) > 50 {
		return fmt.Errorf("email is too long")
	}

	return nil
}

func RegisterUser(db *sql.DB) (string, error) {
	if err := Validation(); err != nil {
		log.Printf("fail: Validation failed, %v\n", err)
		return "", err
	}
	id, err := dao.CreateUser(db)
	if err != nil {
		log.Printf("fail: CreateUser %v\n", err)
		return "", err
	}
	return id, nil
}