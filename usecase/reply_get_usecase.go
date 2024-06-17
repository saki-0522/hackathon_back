package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

// func GetReply(db *sql.DB, ini_tweet_id string, posted_by string) ([]model.ReplyResGet, error) {
func GetReply(db *sql.DB, ini_tweet_id string) ([]model.ReplyResGet, error) {
	// users, err := dao.GetReplyById(db, ini_tweet_id, posted_by)
	users, err := dao.GetReplyById(db, ini_tweet_id)
	if err != nil {
		log.Printf("fail: GetReplyById, %v\n", err)
		return nil, err
	}
	return users, nil
}