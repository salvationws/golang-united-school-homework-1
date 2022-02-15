package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	GetMessage := emoji.Sprint("Hello :world_map:!")
	return GetMessage
}

func main() {
	Message := GetMessage()
	fmt.Println(Message)
}
