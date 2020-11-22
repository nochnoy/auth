package auth

// Авторизация
func Auth(login string, password string) bool {
	return (login == "roma" && password == "horoshiy")
}
