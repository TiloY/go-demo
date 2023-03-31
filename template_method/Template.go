package template_method

import "fmt"

// netty 就是个模板方法模式  父类决定调用顺序   子类自己实现

type Template struct {
}

// 这个据相当于模板方法模式  父类决定调用顺序  子类决定具体实现
func (t *Template) WorkFlow() {
	t.WorkFlowA()
	t.WorkFlowB()
}

func (t *Template) WorkFlowA() {
	fmt.Println("Template::workA")
}

func (t *Template) WorkFlowB() {
	fmt.Println("Template::workB")
}
