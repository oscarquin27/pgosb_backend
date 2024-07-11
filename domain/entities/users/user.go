package user_entity

type User struct {
	UserIdentification
	UserProfile
}

type UserDto struct {
	UserIdentificationDto
	UserProfileDto
}
