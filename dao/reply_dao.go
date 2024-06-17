package dao

import (
	"database/sql"
	"db/model"
	"github.com/oklog/ulid/v2"
	"log"
	"math/rand"
	"time"
)

type ReplyDAO struct {
	DB *sql.DB
}

func NewReplyDAO(db *sql.DB) *ReplyDAO {
	return &ReplyDAO{DB: db}
}

// func GetReplyById(db *sql.DB, ini_tweet_id string, posted_by string) ([]model.ReplyResGet, error) {
func GetReplyById(db *sql.DB, ini_tweet_id string) ([]model.ReplyResGet, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin(), %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)
	rows, err := tx.Query("SELECT * FROM reply WHERE ini_tweet_id = ?", ini_tweet_id)
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()
	
	replies := make([]model.ReplyResGet, 0)
	// Idはtweet_id、Nameは投稿者のID、
	for rows.Next() {
		var u model.ReplyResGet
		// ここの生身も状態によって変更する
		if err := rows.Scan(&u.Display_name, &u.Time, &u.Content); err != nil {
			return nil, err
		}
		replies = append(replies, u)
	}
	return replies, nil
}

// Postをもらった時の処理,tweetをデータベースに挿入
func CreateReply(db *sql.DB) (string, error) {
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
	reply_id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	// ここ何やっているのか、ここに入れるのかえる
	_, err = tx.Exec("INSERT INTO reply (ini_tweet_id, ini_posted_by, reply_content, display_name, posted_by, posted_at, reply_id) VALUES (?, ?, ?, ?, ?, ?, ?)", model.ReplyPost.Ini_tweet_id, model.ReplyPost.Name, model.ReplyPost.Content, model.ReplyPost.Display_name, t, reply_id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}
// 多分ここで返されている値が違う
	return model.ReplyPost.Content, nil
}