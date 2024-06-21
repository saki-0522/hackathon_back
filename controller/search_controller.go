package controller

import (
	"database/sql"
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
	"db/model"
	"strings"
)

func containsSearchTerm(source string, searchTerm string) bool {
	// 大文字小文字を区別せずに部分一致をチェックする
	source = strings.ToLower(source)
	searchTerm = strings.ToLower(searchTerm)
	return strings.Contains(source, searchTerm)
}

// 関係なくすべてのtweetを取得する
func SearchController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	searchTerm := r.URL.Query().Get("searchTerm")
	tweets, err := usecase.GetSearchTweet(db)
	if err != nil {
		log.Printf("fail: usecase.GetTweet, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	var results []model.TweetResGet

	for _, tweet := range tweets {
		// ツイートの内容またはユーザー名に検索キーワードが含まれているかチェック
		if containsSearchTerm(tweet.Content, searchTerm) || containsSearchTerm(tweet.Display_name, searchTerm) {
			results = append(results, tweet)
		}
	}

	// 検索結果をJSON形式でレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}