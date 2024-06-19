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
func GetReplyController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	uid := r.URL.Query().Get("uid")
	parent_id := r.URL.Query().Get("parent_id")
	if parent_id == "" {
		log.Println("fail: parent_id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	replies, err := usecase.GetReply(db, parent_id, uid)
	if err != nil {
		log.Printf("fail: usecase.GetReply, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(replies)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}