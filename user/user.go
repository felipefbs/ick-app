package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        int8   `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
}

func NewUser(username, name, gender, birthdate, password string) (*User, error) {
	byteArray, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:  username,
		Name:      name,
		Gender:    gender,
		Birthdate: birthdate,
		Password:  string(byteArray),
	}, nil
}
