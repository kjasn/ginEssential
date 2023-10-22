package dto

import "Kjasn/ginEssential/model"

// UserDto 封装 http返回    只返回用户名和手机号
type UserDto struct {
	Username  string `json:"username"`
	Telephone string `json:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username:  user.Name,
		Telephone: user.Telephone,
	}
}
