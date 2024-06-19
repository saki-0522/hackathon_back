package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	"log"
)

func LikeValidation() error {
	if model.Like.Post_id == "" {
		return fmt.Errorf("post_id is empty")
	}
	if model.Like.Id == "" {
		return fmt.Errorf("id is empty")
	}
	if len(model.Like.Id) > 50 {
		return fmt.Errorf("id is too long")
	}
	if len(model.Like.Post_id) > 50 {
		return fmt.Errorf("post_id is too long")
	}
	return nil
}

func RegisterLike(db *sql.DB) (string, error) {
	if err := LikeValidation(); err != nil {
		log.Printf("fail: LikeValidation failed, %v\n", err)
		return "", err
	}
	id, err := dao.CreateLike(db)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return "", err
	}
	return id, nil
}

func RegisterLike(db *sql.DB) (string, error) {
	id, err := dao.DeleteLike(db)
	if err != nil {
		log.Printf("fail: %v\n", err)
		return "", err
	}
	return id, nil
}