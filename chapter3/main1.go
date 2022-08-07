package main

func main() {
	a := [3]string{
		"apple",
		"banana",
		// "cinamon"のあとが文末と認識されて;が挿入されてしまう。
		//"cinamon" //コンパイルエラーだが,を入れると大丈夫
		"cinamon", //コンパイルエラーだが,を入れると大丈夫
	}

	var {
		a, b int
		fullName string
	}

 	index := 1 // 推論されてint型になる
	isTrue := true // 真偽値型
	f := 3.14 // float64型
	season := "april!" // 文字列型

	rawString := `
Goの
複数行文字列。
PHPでEOFを使って複数行文字を表すときみたい！
Javaでも後のバージョンでようやく実現されました
`

}

func samplePlus(x, y, int) (int, int) {
	var rtnCd = 0
	return x + y, rtnCd
}
q, _ := samplePlus(1, 2) // リターンコード的な2つめの戻り値は無視
result, err := samplePlus(100, 200) // 両方使うパターン

func plus (x, y int) int {
	return x + y
}
var funcToVar = plus // ここで()はつけない
funcToVar(1, 2) // 実行すると3が戻り値

func returnFunc func() {
	return func() {
	  fmt.Println("関数を返しちゃうニャン")
	}
}
f := returnFunc() // 関数を実行するので()つき、その戻り値の無名関数がfに代入
f() // 実行すると"関数を返しちゃうニャン"

func executeF(f func()) {
	f()
}
sampleF = func() { fmt.Println("関数実行でウィス") }
executeF(sampleF)  //  "関数実行でウィス"

const (
	X = 1 // int型でなく"整数値を持つ定数"になる。最大値がないのでint型変換時にオーバーフローすることも
	Y // 1が入る
	Z
	F64 float64 = 1.222 // 明確に型を指定した場合
	F62a = float64(1.222) // 型変換を使うこちらが好まれる
)

const (
	A = iota // 0が入る 1+iotaにすると1始まり
	B // B = iota の省略形。1が入る
	C // C = iota の省略形。2が入る
)