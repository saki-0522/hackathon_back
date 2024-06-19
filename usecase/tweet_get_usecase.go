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
func GetTweet(db *sql.DB, uid string) ([]model.TweetReturn, error) {
	tweets, err := dao.GetAllTweet(db)
	if err != nil {
		log.Printf("fail: GetAllTweet, %v\n", err)
		return nil, err
	}
	likes, err := dao.GetStatusById(db, uid)
	if err != nil {
		log.Printf("fail: GetStatusById, %v\n", err)
		return nil, err
	}

	likesMap := make(map[string]bool)
	for _, id := range likes {
		likesMap[id] = true
	}

	var tweet_ret []model.TweetReturn
	for _, tweet := range tweets {
		var status int
		if likesMap[tweet.Id] {
			status = 1
		} else {
			status = 0
		}
	
		// TweetReturn を作成して新しい配列に追加する
		tweetReturn := model.TweetReturn{
			Id:      tweet.Id,
			Name:    tweet.Name,
			Time:    tweet.Time,
			Content: tweet.Content,
			Likes:   tweet.Likes,
			Status:  status,
			Parent_Id: tweet.Parent_Id,
		}
	
		tweet_ret = append(tweet_ret, tweetReturn)
	}
	return tweet_ret, nil
}