package controller

import (
	"database/sql"
	"db/model"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// func RegisterLikeController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&model.Like); err != nil {
// 		log.Printf("fail: json.Decode, %v\n", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	log.Println(model.Like)

// 	id, err := usecase.RegisterLike(db)
// 	if err != nil {
// 		log.Printf("fail: %n\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	// 成功した場合のレスポンス
// 	w.WriteHeader(http.StatusOK)
// 	response := map[string]string{"id": id}
// 	bytes, err := json.Marshal(response)
// 	if err != nil {
// 		log.Printf("fail: json.Marshal, %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(bytes)
// }

func RegisterLikeController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	statusStr := r.URL.Query().Get("status")
	status, err1 := strconv.Atoi(statusStr)
	uid := r.URL.Query().Get("uid")
	if err1 != nil {
		log.Println("Invalid status value")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(uid)
	if uid == "" {
		log.Println("fail: uid is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("test");
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model.Like); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	var id string
	var err error
	if (status == 0){
		id, err = usecase.RegisterLike(db)
		if err != nil {
			log.Printf("fail: %n\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else{
		id, err = usecase.DeleteLike(db)
		if err != nil {
			log.Printf("fail: %n\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	
	// 成功した場合のレスポンス
	w.WriteHeader(http.StatusOK)
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