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

// func GetReplyById(db *sql.DB, parent_id string, posted_by string) ([]model.ReplyResGet, error) {
// func GetReplyById(db *sql.DB, parent_id string) ([]model.ReplyResGet, error) {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Printf("fail: db.Begin(), %v\n", err)
// 		return nil, err
// 	}
// 	defer HandleTransaction(tx, err)
// 	rows, err := tx.Query("SELECT * FROM tweet WHERE parent_id = ?", parent_id)
// 	if err != nil {
// 		log.Printf("fail: tx.Query, %v\n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()
	
// 	replies := make([]model.ReplyResGet, 0)
// 	for rows.Next() {
// 		var u model.ReplyResGet
// 		if err := rows.Scan(&u.Id, &u.Name, &u.Time, &u.Content, &u.Likes, &u.Parent_Id, &u.Display_name); err != nil {
// 			return nil, err
// 		}
// 		replies = append(replies, u)
// 	}
// 	return replies, nil
// }

func GetStatusById(db *sql.DB, parent_id string, uid string) ([]string, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin(), %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)

	rows, err := tx.Query("SELECT post_id FROM likes WHERE id = ? AND parent_id = ?", uid, parent_id)
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()

	status := make([]string, 0)
	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			return nil, err
		}
		status = append(status, tmp)
	}
	return status ,nil
}

func GetAllReplyById(db *sql.DB, parent_id string) ([]model.ReplyResGet, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin(), %v\n", err)
		return nil, err
	}
	defer HandleTransaction(tx, err)
	rows, err := tx.Query("SELECT * FROM tweet WHERE parent_id = ?", parent_id)
	if err != nil {
		log.Printf("fail: tx.Query, %v\n", err)
		return nil, err
	}
	defer rows.Close()
	
	replies := make([]model.ReplyResGet, 0)
	// Idはtweet_id、Nameは投稿者のID、
	for rows.Next() {
		var u model.ReplyResGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Time, &u.Content, &u.Likes, &u.Parent_Id, &u.Display_name); err != nil {
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
	_, err = tx.Exec("INSERT INTO tweet (tweet_id, posted_by, posted_at, content, display_name, parent_id) VALUES (?, ?, ?, ?, ?, ?)", reply_id, model.ReplyPost.Name, t, model.ReplyPost.Content, model.ReplyPost.Display_name, model.ReplyPost.Parent_Id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}
// 多分ここで返されている値が違う
	log.Printf("CreateReply")
	log.Println(model.ReplyPost)
	return model.ReplyPost.Content, nil
}