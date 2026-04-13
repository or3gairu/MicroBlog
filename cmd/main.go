package main

import (
	"MicroBlog/internal/models"
	"fmt"
)

func main() {
	//testing a new user create (3x)
	user := models.UserTemplate.CreateNewUser("Aleks")
	user1 := models.UserTemplate.CreateNewUser("Gleb")
	user2 := models.UserTemplate.CreateNewUser("Sonnet")
	//making simply two types user storages as a slice and map
	userStorage := []*models.User{user, user1, user2}
	userMapStorage := make(map[int]string)
	//fill the map storage
	for _, v := range userStorage {
		uid, uname := v.GetUserInfo()
		userMapStorage[uid] = uname
	}
	//output storage
	fmt.Println(userMapStorage)

	/*	for _, u := range userStorage {
			uid, name := u.GetUserInfo()
			fmt.Printf("UID: %d || Name: %s\n", uid, name)
		}
	*/
}
