package dao

import (
	"database/sql"
	"db/model"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"time"
)

type TweetDAO struct {
	DB *sql.DB
}

func NewTweetDAO(db *sql.DB) *TweetDAO {
	return &TweetDAO{DB: db}
}

// tweet用に変更する

func GetAllTweet(db *sql.DB) ([]model.TweetResGet, error) {
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
	
	tweets := make([]model.TweetResGet, 0)
	// Idはtweet_id、Nameは投稿者のID、
	for rows.Next() {
		var u model.TweetResGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Time, &u.Content); err != nil {
			return nil, err
		}
		tweets = append(tweets, u)
	}
	defer rows.Close()
	return tweets, nil
}

// Postをもらった時の処理,tweetをデータベースに挿入
func CreateTweet(db *sql.DB) (string, error) {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return "", err
	}
	defer HandleTransaction(tx, err)

	// IDの生成
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	tweet_id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	// ここ何やっているのか
	_, err = tx.Exec("INSERT INTO tweet (tweet_id, posted_by, posted_at, content, displayName) VALUES (?, ?, ?, ?, ?)", tweet_id, model.TweetPost.Name, t, model.TweetPost.Content, model.TweetPost.DisplayName)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}
// 多分ここで返されている値が違う
	return model.TweetPost.Content, nil
}