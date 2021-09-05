package main

import "fmt"

func main() {

	/* var studentName string = "Zeki Ünyıldız"
	var grade int = 77
	var isPassed bool = true */

	/* var (
		studentName string = "Zeki Ünyıldız"
		grade       int    = 77
		isPassed    bool   = true
	) */

	var studentName, grade, isPassed = "Zeki Ünyıldız", 77, true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}
