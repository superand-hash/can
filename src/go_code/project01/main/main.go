package main //表示该文件归属于main包
import (     //引入包
	"fmt" //引入fmt
)

// 可以从终端输入两个数并显示和
func main() { //main函数有要求,要求在main包中,main函数是我们整个狗程序的入口函数
	//定义两个变量
	var n1 int
	var n2 int
	//从终端接收两个输入值
	fmt.Println("请输入n1的值")
	fmt.Scanln(&n1)
	fmt.Println("请输入n2的值")
	fmt.Scanln(&n2)
	var sum int
	sum = n1 + n2
	fmt.Println("sum=", sum)
}
