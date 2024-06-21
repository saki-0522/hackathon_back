package controller

import (
	"database/sql"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func SearchUserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	uid := r.URL.Query().Get("uid")
	if uid == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// log.Println(uid)
	user, err := usecase.SearchUser(db, uid)
	if err != nil {
		log.Printf("fail: usecase.SearchUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}