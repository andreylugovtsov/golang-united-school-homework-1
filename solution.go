package solution

import (
	// https://pkg.go.dev/github.com/kyokomi/emoji/v2
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
