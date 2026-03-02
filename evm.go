package sodium

import "fmt"

func GetEvmSalt(platform Platform, username string) string {
	norm := NormalizeUsername(username)
	return fmt.Sprintf("%s:%s", platform, norm)
}
