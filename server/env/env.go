package env

import "os"

var (
	JWT_SECRET_KEY  string    = os.Getenv("JWT_SIGNING_KEY")
)