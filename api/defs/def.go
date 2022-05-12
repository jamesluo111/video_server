package defs

type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type NewVideo struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type AddComments struct {
	VideoId  string `json:"video_id"`
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"videoId"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type CommentsInfo struct {
	Comments []*Comment
}

type SimpleSession struct {
	UserName string
	TTL      int64
}
