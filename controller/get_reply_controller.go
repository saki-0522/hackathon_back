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
	ini_tweet_id := r.URL.Query().Get("ini_tweet_id")
	if ini_tweet_id == "" {
		log.Println("fail: ini_tweet_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tweets, err := usecase.GetReply(db, ini_tweet_id)
	if err != nil {
		log.Printf("fail: usecase.GetReply, %v\n", err)
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