package gtc

import (
	"fmt"
	"os"

	"github.com/danawoodman/gtc/internal"
)

func main() {
	if err := internal.NewCmd(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
