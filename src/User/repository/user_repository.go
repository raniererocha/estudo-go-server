package user_repository

import (
	Dto "server/src/User/dto"
)

var userDb = []Dto.User{}

func IsEmailUnique(email string) bool {

	for _, value := range userDb {
		if value.Email == email {
			return false
		}
	}
	return true

}

func StoreUser(user Dto.CreateUserDTO) Dto.User {
	var newUser Dto.User

	newUser.Id = len(userDb) + 1
	newUser.Email = user.Email
	newUser.Name = user.Name

	userDb = append(userDb, newUser)
	return newUser
}
func FindUsers() []Dto.User {
	return userDb
}
