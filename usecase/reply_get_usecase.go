package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

func GetReply(db *sql.DB) ([]model.ReplyResGet, ini_tweet_id string) {
	users, err := dao.GetReplyById(db, ini_tweet_id)
	if err != nil {
		log.Printf("fail: GetReplyById, %v\n", err)
		return nil, err
	}
	return users, nil
}