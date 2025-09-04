package fallthroughladder

import "fmt"

func FallThroughLadder() {
	level := 1

	fmt.Println("--------------------")
	switch level {
	case 1:
		fmt.Println("Level 1 Unlocked")
		fallthrough
	case 2:
		fmt.Println("Level 2 Unlocked")
		fallthrough
	case 3:
		fmt.Println("Level 3 Unlocked")
	}

	level += 1
	fmt.Println("--------------------")

}
