package item

import "fmt"

//武器
type Weapon struct {
	Item   //  暂时 理解为JAVA 中的 extends Item  ,注意是不是继承？？
	Damage int
}

func (weapon *Weapon) DoUse() {
	fmt.Printf("weapon  :: DoUse , ItemId = %d\n", weapon.ItemId)
}
