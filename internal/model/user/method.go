package user

func (u User) GetUserPermissionString() string {
	return string(u.Permission)
}

func (u User) SetUserPermissionString(p string) User {
	switch p {
	case string(Write):
		u.Permission = Write
	default:
		u.Permission = Read
	}
	return u
}
