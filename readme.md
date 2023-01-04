# mask

stdin mask package written in go

### example code

```
package main

import (
	"fmt"
	"github.com/anxtyanxty/mask"
)

func main() {
	var username string
	fmt.Print("[!] Username: ")
	fmt.Scanln(&username)

	fmt.Print("[!] Password: ")
	password := mask.Password('*') // masks every character on screen with "*" and puts input into variable

	if login(username, password) {
		fmt.Println("We have logged in :) ")
	} else {
		fmt.Println("Wrong password. Please try again later.")
	}
}

func login(username string, password string) bool {
	return username == "me" && password == "notpassword"
}
```
