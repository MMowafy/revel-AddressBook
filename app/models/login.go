package models

type UserLogin struct {
	Email string
	Password string
}

func VerifyUser(user UserLogin)  bool  {
	if user.Email!="" && user.Password!="" {
		return true

	} else {
		return false
	}

}
