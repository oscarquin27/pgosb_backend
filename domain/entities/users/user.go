package user_entity

type User struct {
	UserIdentification
	UserProfile
	UserStatus
}

type UserDto struct {
	UserIdentificationDto
	UserProfileDto
}
