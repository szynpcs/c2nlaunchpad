package main

import "fmt"

func main() {
	fmt.Println("空接口 + swith + 断言")
	//switchany(1)
	//switchany("fda")
	//switchany(true)
	//switchinterface(1)
	var cuserror error = &cuserror{"这啥啊"}
	switchinterface(cuserror)
}

func switchany(a any) {
	fmt.Println("自由之翼")
	if val, ok := a.(int); ok {
		fmt.Println(a, "是第一个整形")
		fmt.Println(val)
	}

	if val, ok := a.(string); ok {
		fmt.Println(val, "这是一个字符串")
	}

	if val, ok := a.(bool); ok {
		fmt.Println(val, "这是一个布尔值")
	}
}

func switchinterface(val interface{}) {
	fmt.Println("INTERFACE")
	switch v := val.(type) {
	case int:
		fmt.Println(v, "这是一个整数类型")
	case float64:
		fmt.Println(v, "这是一个浮点类型")

	case bool:
		fmt.Println(v, "这是一个布尔类型")
	case *cuserror:
		fmt.Println(v, "我自定义了一个错误类型很刺激")
	default:
		fmt.Println("不知道是个什么鸡毛")

	}

	a := 'x'
	switch a {
	case 'x':
		fmt.Println("加墨")

	}
}

type cuserror struct {
	a string
}

func (e *cuserror) Error() string {
	return "我尼玛自定义了一个异常报错" + e.a

}
