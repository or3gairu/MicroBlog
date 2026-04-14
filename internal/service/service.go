package service

import (
	"MicroBlog/internal/models"
	"sort"
	"strings"
)

type Service struct {
	usersQuantity int
	postsQuantity int

	userStorage map[int]*models.User
	postStorage map[int]*models.Post
}

func New() *Service {
	return &Service{
		usersQuantity: 0,
		postsQuantity: 0,

		userStorage: make(map[int]*models.User),
		postStorage: make(map[int]*models.Post),
	}
}

func (s *Service) Register(username string) (*models.User, error) {
	var illegalCharSymbols = "!@#$%^&*(_+`~){}[]|/<>,:; "
	var currentID = s.usersQuantity

	if username == "" {
		return nil, ErrEmptyUsername
	}
	if strings.ContainsAny(username, illegalCharSymbols) {
		return nil, ErrIllegalCharSymbol
	}

	user := &models.User{
		ID:         currentID,
		Username:   username,
		LikedPosts: make(map[int]bool),
	}

	s.usersQuantity++

	s.userStorage[user.ID] = user

	return user, nil
}

func (s *Service) CreatePost(userID int, message string) (*models.Post, error) {
	var msg = strings.TrimSpace(message)
	var currentID = s.postsQuantity

	if msg == "" {
		return nil, ErrMessageEmpty
	}

	if _, err := s.FindUserByID(userID); err != nil {
		return nil, ErrUserNotFound
	}

	post := &models.Post{
		ID:       currentID,
		AuthorID: userID,
		Message:  msg,
		Likes:    make(map[int]int),
	}

	s.postsQuantity++

	s.postStorage[post.ID] = post

	return post, nil
}

func (s *Service) LikePost(postID, userID int) error {
	currentPost, pErr := s.FindPostByID(postID)
	currentUser, uErr := s.FindUserByID(userID)

	if pErr != nil {
		return ErrPostNotFound
	}
	if uErr != nil {
		return ErrUserNotFound
	}

	if s.IsPostLikedByUser(userID, postID) {
		currentPost.Likes[postID]--
		currentUser.LikedPosts[postID] = false
		return nil
	}

	currentPost.Likes[postID]++
	currentUser.LikedPosts[postID] = true

	return nil
}

func (s *Service) IsPostLikedByUser(userID, postID int) bool {
	return s.userStorage[userID].LikedPosts[postID]
}

func (s *Service) FindUserByID(userID int) (*models.User, error) {
	if _, ok := s.userStorage[userID]; !ok {
		return nil, ErrUserNotFound
	}
	return s.userStorage[userID], nil
}

func (s *Service) FindUserByName(username string) ([]*models.User, error) {
	var userListID = make([]*models.User, 0, s.usersQuantity)

	for _, v := range s.userStorage {
		if v.Username == username {
			userListID = append(userListID, v)
		}
	}

	if len(userListID) == 0 {
		return nil, ErrUserNotFound
	}

	return userListID, nil
}

func (s *Service) FindPostByID(postID int) (*models.Post, error) {
	if _, ok := s.postStorage[postID]; !ok {
		return nil, ErrPostNotFound
	}
	return s.postStorage[postID], nil
}

func (s *Service) FindPostByAuthor(authorID int) ([]*models.Post, error) {
	var postList = make([]*models.Post, 0, s.postsQuantity)
	var postIndexes = make([]int, 0, s.postsQuantity)

	for _, p := range s.postStorage {
		if p.AuthorID == authorID {
			postIndexes = append(postIndexes, p.ID)
		}
	}

	if len(postIndexes) == 0 {
		return nil, ErrPostNotFound
	}

	sort.Ints(postIndexes)

	for i := range postIndexes {
		postList = append(postList, s.postStorage[postIndexes[i]])
	}
	return postList, nil
}
