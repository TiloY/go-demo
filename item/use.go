package item

type IUse interface {
	DoUse()
	//DoUse(name string) // 实现参数的条件就是 函数名一样  参数也是一样
}
