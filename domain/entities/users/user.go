package user_entity

type User struct {
	UserIdentification
	UserProfile
	UserStatus
}

type UserCreateDto struct {
	UserIdentification
	UserProfile
}

type UserUpdateDto struct {
	UserIdentification
	UserProfile
}