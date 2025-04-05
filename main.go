package main

import (
	"smart-deals-api/user/signin"
	"smart-deals-api/user/signup"
)

func main() {
	signin.SignIn()
	signup.SignUp()
}
