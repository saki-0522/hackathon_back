package model

import (
	"time"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Email  string    `json:"email"`
}

type RequestDataType struct {
	Name string `json:"name"`
	Email  string    `json:"email"`
}

var RequestData struct {
	Name string `json:"name"`
	Email  string    `json:"email"`
}

var RegisterData struct {
	Name string `json:"displayName"`
	Email string `json:"email"`
	Id string `json:"uid"`
}

type TweetResGet struct {
	Id   string `json:"tweet_id"`
	Name string `json:"name"`
	Time  string    `json:"posted_at"`
	Content  string    `json:"content"`
	Likes int `json:"likes"`
	Heart int `json:"flag"`
}

type TweetRes struct {
	Id   string `json:"tweet_id"`
	Name string `json:"posted_by"`
	Time  string    `json:"posted_at"`
	Content  string    `json:"content"`
}

var TweetPost struct {
	Name string `json:"posted_by"`
	Content  string    `json:"content"`
	Display_name string `json:"displayName"`
}

// このデータの中身を考える
type ReplyResGet struct {
	Display_name   string `json:"display_name"`
	Time  time.Time    `json:"posted_at"`
	Content  string    `json:"reply_content"`
}

// Replyを登録するときに、フロントから受け取る値たち
var ReplyPost struct {
	Name string `json:"posted_by"`
	Content  string    `json:"reply_content"`
	Ini_tweet_id string `json:"ini_tweet_id"`
	Display_name string `json:"display_name"`
}

// likes tableに入力する値
var Like struct {
	Post_id string `json:"post_id"`
	Id string `json:"id"`
}
