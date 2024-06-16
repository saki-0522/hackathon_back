package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"fmt"
	"log"
)

func TweetValidation() error {
	if model.TweetPost.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if len(model.TweetPost.Name) > 50 {
		return fmt.Errorf("name is too long")
	}

	if len(model.TweetPost.Content) > 200 {
		return fmt.Errorf("content is too long")
	}

	return nil
}

func RegisterTweet(db *sql.DB) (string, error) {
	if err := TweetValidation(); err != nil {
		log.Printf("fail: TweetValidation failed, %v\n", err)
		return "", err
	}
	content, err := dao.CreateTweet(db)
	if err != nil {
		log.Printf("fail: CreateTweet %v\n", err)
		return "", err
	}
	return content, nil
}