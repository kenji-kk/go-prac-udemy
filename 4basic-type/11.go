package main

import "fmt"

//型
//整数型

func main() {
	//bit数を選択しない場合は環境依存で決まる
	var i int = 100

	//明示的にbit数を指定
	var i2 int64 = 200
	fmt.Println(i + 50)

	//fmt.Println(i + i2) bit数が違うのでエラーになる

	//%Tで型を表示できる
	fmt.Printf("%T\n", i2)

	//型変換
	fmt.Printf("%T\n", int32(i2))

	//違うbit同士の計算
	fmt.Println(i + int(i2))
}
