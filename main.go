package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"db/controller"
	"db/dao"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// ① GoプログラムからMySQLへ接続
var db *sql.DB
var userDao *dao.UserDAO

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// // ①-1
	mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PWD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	fmt.Println(mysqlDatabase)

    connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
    _db, err := sql.Open("mysql", connStr)

	// ①-2
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db

	userDao = dao.NewUserDAO(db)
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodOptions:
		return 
	// Login画面からGETがくる、displayNameを送り返す
	case http.MethodGet:
		controller.SearchUserController(w, r, db)
	// Singup画面からPOSTがくる、データベースに保存する
	case http.MethodPost:
		controller.RegisterUserController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// HomeとPost画面での処理
func handlerTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodOptions:
		return 
	// Home画面からGetリクエストがくる、時間、投稿内容、投稿者を返す
	case http.MethodGet:
		controller.GetTweetController(w, r, db)
	// Post画面からPOSTがくる、データベースに保存する
	case http.MethodPost:
		controller.RegisterTweetController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handlerHeart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodOptions:
		return 
	// Getに来るものは今のところない
	case http.MethodGet:
		// controller.GetHeartController(w, r, db)
		return
	// Heartが押されたら
	case http.MethodPost:
		controller.RegisterLikeController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handlerReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodOptions:
		return 
	// Getに来たものに対して、条件の合うリプライを全部返す
	case http.MethodGet:
		controller.GetReplyController(w, r, db)
	// Postできたものに対してハートの状態を保存する
	case http.MethodPost:
		return 
		// log.Printf("post")
		// controller.RegisterReplyController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handlerSearchTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodOptions:
		return 
	// Getに来たものに対して、条件の合うリプライを全部返す
	case http.MethodGet:
		controller.SearchController(w, r, db)
		log.Printf("register")
	// Postできたものに対してハートの状態を保存する
	case http.MethodPost:
		return
		// log.Printf("post")
		// controller.RegisterReplyController(w, r, db)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)
	http.HandleFunc("/tweet", handlerTweet)
	http.HandleFunc("/heart", handlerHeart);
	http.HandleFunc("/reply", handlerReply);
	http.HandleFunc("/tweet/search", handlerSearchTweet);

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}