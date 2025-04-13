package utils

import (
	"strings"
	"fmt"
)


func Replace() {
	broken := "G# r#ck!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}

