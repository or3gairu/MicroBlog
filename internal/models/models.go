package models

type User struct {
	ID         int
	Username   string
	LikedPosts map[int]bool
}

type Post struct {
	ID       int
	AuthorID int
	Message  string
	Likes    map[int]int //key - postid, value - count of likes
}
