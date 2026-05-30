// 定义了go中的一个包pkg,包是Go语言的基本组成单元，通常使用单个的小写单词命名，一个Go程序本质上就是一组包的集合。所有Go代码都有自己隶属的包，在这里我们的“hello，world”示例的所有代码都在一个名为main的包中。main包在Go中是一个特殊的包，整个Go程序中仅允许存在一个名为main的包。
package main

//导入fmt包的包路径
import "fmt"

//定义一个main函数，当运行一个可执行的 Go 程序的时候，所有的代码都会从这个入口函数开始运行。
func main() {
	//打印输出hello，world
	fmt.Println("hello,world")
}
