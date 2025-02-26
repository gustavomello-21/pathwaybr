package hasher

func HashPassword(password string) (string, error) {
	return password, nil
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	return hashedPassword == password
}
