package main // JAVA  package

import (
	"go-demo/item"
	"go-demo/template_method"
) // JAVA import class but fmt  is a  directory

//func (n *int) SayHello()  {
//
//}

//func reeseReplacement(usingItem *item.Item) {
//	usingItem.DoUse()
//}

// 为什么不要 *
func reeseReplacement(useImpl item.IUse) {
	useImpl.DoUse()
}

// 匿名接口
func reeseReplacement3(useImpl interface{ DoUse() }) {
	useImpl.DoUse()
}

func main() { // void  main()
	//fmt.Println("hello , the world")
	//fmt.Println("hello  user = " + user.GetUserName() + " , password= " + user.GetPassWord())

	//Item item = new Item()
	//newItem := &item.Item{}
	//newItem.ItemId = 10086
	//newItem.DoUse()

	newWeaPon := &item.Weapon{}
	newWeaPon.ItemId = 10011
	newWeaPon.DoUse()
	newWeaPon.ItemId = 16888
	newWeaPon.DoUse()
	// 这里不是继承  更像是组合
	// 继承的话 应该满足一个原则  李氏替换原则   适合父类的地方也适合子类 这就是李氏替换原则
	reeseReplacement(newWeaPon) //  不能方结构体 可以放一个接口

	//模板方法的 go 写起来就是可能比较麻烦

	myMpl := &template_method.MyImpl{}
	myMpl.WorkFlow()

	/**  这个结构就是因为不是继承  所以没有重写
	Template::workA
	Template::workB
	*/

	// JAVA 面试题  ： 继承  1. 符合李氏替换原则  2 .  继承就会有重写 覆盖 不是继承就没有覆盖的那个机制

	// & 表示什么？
	// * 表示什么？
}
