package isyntax

// func TestISyntaxFuncsInit() struct{} {
// 	/*
// 		init函数的主要特点:
// 			init函数先于main函数自动执行，不能被其他函数调用
// 			init函数没有输入参数、返回值
// 			每个包可以有多个init函数
// 			包的每个源文件也可以有多个init函数，这点比较特殊
// 			同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序
// 			不同包的init函数按照包导入的依赖关系决定执行顺序

// 		golang初始化顺序:
// 			初始化顺序: 变量初始化 -> init() -> main()
// 	*/

// 	log.Println("pre init")

// 	// 类比 var x interface{} & slice{} ...
// 	var x struct{}
// 	return x
// }

// func init() {
// 	log.Println("init")
// }

// func init() {
// 	log.Println("init again")
// }

// func TestISyntaxFuncsInit2() struct{} {
// 	log.Println("pre init again")

// 	var x struct{}
// 	return x
// }

// var v struct{} = TestISyntaxFuncsInit()
