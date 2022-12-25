package utils

import "os"

var Constants = struct {
	JWT_SECRET_KEY string
}{
	os.Getenv("JWT_SECRET_KEY"),
}
