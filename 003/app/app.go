package app

import "fmt"

var (
	AppName    = "chihuo"
	appVersion = "1.0.0"
)

func init() {
	fmt.Printf("app version is %s\n", appVersion)
}
