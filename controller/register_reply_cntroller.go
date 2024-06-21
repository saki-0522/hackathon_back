package controller

import (
	"database/sql"
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterReplyController(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model.ReplyPost); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := usecase.RegisterReply(db)
	log.Printf("RegisterReplyController")
	if err != nil {
		log.Printf("fail: RegisterReply %n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 成功した場合のレスポンス
	w.WriteHeader(http.StatusOK)
	// ここが違う
	response := map[string]string{"id": id}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}