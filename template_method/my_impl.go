package template_method

import "fmt"

type MyImpl struct {
	Template
}

func (MyImpl) WorkFlowB() {
	fmt.Printf("MyImpl:: WorkB")
}
