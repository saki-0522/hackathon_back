package model

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
	Display_name string `json:"displayName`
}

// このデータの中身を考える
type ReplyResGet struct {
	Ini_id   string `json:"tweet_id"`
	Name string `json:"name"`
	Time  string    `json:"posted_at"`
	Content  string    `json:"content"`
}

