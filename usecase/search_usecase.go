package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

func GetSearchTweet(db *sql.DB) ([]model.TweetResGet, error) {
	tweets, err := dao.GetSearchTweet(db)
	if err != nil {
		log.Printf("fail: GetAllTweet, %v\n", err)
		return nil, err
	}
	return tweets, nil
}