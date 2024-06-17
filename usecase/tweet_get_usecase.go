package usecase

import (
	"database/sql"
	"db/dao"
	"db/model"
	"log"
)

// func SearchUser(db *sql.DB, name string) ([]model.UserResForHTTPGet, error) {
// 	users, err := dao.GetUserByName(db, name)
// func SearchUser(db *sql.DB, uid string) ([]model.UserResForHTTPGet, error) {
func GetTweet(db *sql.DB, uid string) ([]model.TweetResGet, error) {
	tweets, err := dao.GetAllTweet(db, uid)
	if err != nil {
		log.Printf("fail: GetAllTweet, %v\n", err)
		return nil, err
	}
	return tweets, nil
}