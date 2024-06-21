package dao

import (
	"database/sql"
	"db/model"
	"log"
)

func GetSearchTweet(db *sql.DB) ([]model.TweetResGet, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin(), %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)
	rows, err := tx.Query("SELECT * FROM tweet")
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()
	
	tweets := make([]model.TweetResGet, 0)
	// Idはtweet_id、Nameは投稿者のID、
	for rows.Next() {
		var u model.TweetResGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Time, &u.Content, &u.Likes, &u.Parent_Id, &u.Display_name); err != nil {
			return nil, err
		}
		tweets = append(tweets, u)
	}
	return tweets, nil
}

