package main

// "运算符优先级"在英文中通常被称为 "Operator Precedence"。

//从低到高 #
//||
//&&
//<- (通道操作符，后面会讲到，暂时先忽略)
//== != < <= > >=
//+ - | ^
//* / % << >> & &^
//^ !

func main() {
	println(((1+3)*2 + 4) * 6 / 3)
}