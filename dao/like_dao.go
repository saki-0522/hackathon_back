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

	_, err = tx.Exec("INSERT INTO likes (post_id, id) VALUES (?, ?)", model.Like.Post_id, model.Like.Id)
	if err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return "", err
	}

	return model.Like.Id, nil
}