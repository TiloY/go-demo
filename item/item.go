package item

import "fmt"

// Item  数据 和行为分离了  /**
type Item struct { //可以暂时理解位JAVA类
	ItemId   int    //大写字母 相当于 JAVA public  int itemId
	itemName string // private String itemName
}

func (thisItem *Item) DoUse() { // public void doUser(){}  * 是
	// item --> this
	fmt.Printf("Item::DoUse,ItemId = %d", thisItem.ItemId)
}

//class Item{
//
//	public void DoUse(){
//
//   }
//
//}
