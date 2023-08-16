package responses

import "morris/im/models"

type ContactResponse struct {
}

func (cc *ContactResponse) FriendsResponse(userList []models.UserModel) []models.UserModel {

	for index, _ := range userList {
		userList[index].Passwd = ""
		userList[index].Salt = ""
		userList[index].Id = 0
	}
	return userList
}
