package controller

import (
	"database/sql"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	// "db/model"
)

// 関係なくすべてのtweetを取得する
func GetTweetController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	uid := r.URL.Query().Get("uid")
	tweets, err := usecase.GetTweet(db, uid)
	if err != nil {
		log.Printf("fail: usecase.GetTweet, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(tweets)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}