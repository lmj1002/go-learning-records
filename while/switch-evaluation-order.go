package main

//switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today!")
	case today + 1:
		fmt.Println("tomorrow!")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")

	}
	//*注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义留给读者去发现。
	// go's birthday
}
