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

func GetAllTweet(db *sql.DB, uid string) ([]model.TweetResGet, error) {
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
		if err := rows.Scan(&u.Id, &u.Name, &u.Time, &u.Content); err != nil {
			return nil, err
		}
		// db, err := tx.Query("SELECT COUNT(*) FROM likes WHERE post_id = 'id'")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer db.Close()

		// これは一回しか呼ばれない
		// for db.Next() {
			// if err := db.Scan(&u.Likes); err != nil {
			// 	return nil, err
			// }
		// }

		// var exists bool
		// err2 := tx.QueryRow("SELECT EXISTS (SELECT 1 FROM likes WHERE id = ? AND post_id = ?)", uid, u.Name).Scan(&exists)
		// if err2 != nil {
		// 	log.Printf("fail: tx.QueryRow, %v\n", err2)
		// 	return nil, err2
		// }
		// if exists {
		// 	u.Heart = 1
		// } else {
		// 	u.Heart = 0
		// }
		

		tweets = append(tweets, u)
	}
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
	_, err = tx.Exec("INSERT INTO tweet (tweet_id, posted_by, posted_at, content, display_name) VALUES (?, ?, ?, ?, ?)", tweet_id, model.TweetPost.Name, t, model.TweetPost.Content, model.TweetPost.Display_name)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}
// 多分ここで返されている値が違う
	return model.TweetPost.Content, nil
}