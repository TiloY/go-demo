package user // go 语言所有的文件名都是小写

func GetUserName() string {
	return "北纬8°"
}

func enCode(orig string) string { //  首字母小写 就是相当于 private
	return orig + "HongKong.china"
}

// go 没有 class   有结构体
