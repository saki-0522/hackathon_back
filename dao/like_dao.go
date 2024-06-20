package dao

import (
	"database/sql"
	"db/model"
	"log"
)

type LikeDAO struct {
	DB *sql.DB
}

func NewLikeDAO(db *sql.DB) *LikeDAO {
	return &LikeDAO{DB: db}
}

func CreateLike(db *sql.DB) (string, error) {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return "", err
	}
	defer HandleTransaction(tx, err)

	_, err = tx.Exec("INSERT INTO likes (post_id, id, parent_id) VALUES (?, ?, ?)", model.Like.Post_id, model.Like.Id, model.Like.Parent_Id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	_, err = tx.Exec("UPDATE tweet SET like_count = like_count + 1 WHERE tweet_id = ?", model.Like.Post_id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	return model.Like.Id, nil
}

func DeleteLike(db *sql.DB) (string, error) {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return "", err
	}
	defer HandleTransaction(tx, err)

	_, err = tx.Exec("DELETE FROM likes WHERE post_id = ? AND id = ? AND parent_id = ?", model.Like.Post_id, model.Like.Id, model.Like.Parent_Id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	_, err = tx.Exec("UPDATE tweet SET like_count = like_count - 1 WHERE tweet_id = ?", model.Like.Post_id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	return model.Like.Id, nil
}