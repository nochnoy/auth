package auth

// Auth сраная функция
func Auth(login string, password string) bool {
	return (login == "roma" && password == "horoshiy")
}
