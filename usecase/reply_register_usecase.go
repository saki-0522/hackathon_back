package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	"log"
)

func ReplyValidation() error {
	if model.ReplyPost.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if len(model.ReplyPost.Name) > 50 {
		return fmt.Errorf("name is too long")
	}

	if len(model.ReplyPost.Content) > 200 {
		return fmt.Errorf("reply_content is too long")
	}

	return nil
}

func RegisterReply(db *sql.DB) (string, error) {
	if err := ReplyValidation(); err != nil {
		log.Printf("fail: ReplyValidation failed, %v\n", err)
		return "", err
	}
	reply_content, err := dao.CreateReply(db)
	if err != nil {
		log.Printf("fail: CreateReply %v\n", err)
		return "", err
	}
	return reply_content, nil
}