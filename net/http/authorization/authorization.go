package authorization

import "fmt"

func Format(authType, credentials string) string {
	return fmt.Sprintf("%s %s", authType, credentials)
}
