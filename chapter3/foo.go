package foo

import "fmt"

const (
	OUTER_MAX    = 999
	internal_max = 99
)

var (
	OuterVar    = 1
	internalVar = 2
)

func OuterFunc() {
}
func internalFunc() {
}

func defineFireYokai() {
	const word = "メラメーラ！"
	const name = "メラメライオン"
}
func someFunc2(a int) (b string) {
	//fmt.Println(word + name) // スコープ外でエラー
	//var a int // 引数の再定義はエラー
	//var b string // 戻り値の再定義はエラー
	return
}

func etc() {
	for {
		fmt.Println("無限ループだニャン")
	}
	// 伝統的なループ
	for i := 0; i < 10; i++ {
		if i == 9 {
			break // breakも他言語と同じ。continueはネストのひとつ外側へ
		}
	}
	// fruits := [2]{"banana", "cherry"}
	// for i, name := range fruits {
	// 	// インデックスがi、nameにi番目の文字列が入る
	// }
	x := false
	if x == 1 {
		// 真の時の処理
	}

	if true {
		// 実行される。上が変数だったら()なしのはず
	}
	if x, y := 1, 2; x < y {
		// 実行される。変数x,yのスコープはこのifブロックの中だけ
	}
	if _, err := doSomething(); err != nil {
		// 関数の2番目の戻り値がエラーを返した時の処理。Goでよくやる書き方
	}

	n := 3
	switch n {
	case 1, 2:
	// 条件を満たすときの処理。Goはここにbreakを書かなくても
	// caseの処理が終わる。(=フォールスルーでない) 素晴らしい！
	case 3, 4:
		// 逆にここに予約語 fallthrough を書くと、
		// 次のcase節とも比較してくれる。
		n = 5
		fallthrough
	case 5:
	// 5のときの処理
	default:
	}
	switch {
	case n >= 0:
	// case節に条件でbool型を返す式も書ける。この時、switchの後の式は省略できるし省略すべき
	case n < 0:
		//case 1:
		// ひとつのswitch文で定数と式の両方があるとエラー。
	}

	// なんでも入るinterface{}型なので、xはint型ではない
	var x interface{} = 3.14
	i, isInt := x.(int)         //iは3、isIntはfalse
	f, isFloat64 := x.(float64) // iは3.14、isFloat64はtrue

	// switchで分岐できる。この時typeは変数名でなく予約語で参照できない
	switch x.(type) {
	case bool:
		fmt.Println("xは論理型", x)
	case int, uint:
		// 続く
	}
}

func deferSample() {
	defer fmt.Println("1")
	defer fmt.Println("2") // 複数定義すると最後から順に実行
	defer func() {
		// 複数処理があったら無名関数で書ける。実行するので最後の()も必要
	}()

	// 関数のメイン処理
}

func main2() {
	defer func() { // Javaなどのfinally的な処理
		if x := recover(); x != nil {
			// 戻り値のxがnilでなければpanicが実行されている。
			fmt.Println(x) // "パニック発生！"の文字列が取れる。
		}
	}()

	panic("パニック発生！") // 文字列以外にもinterface{}型でいろいろ渡せる
	fmt.Println("ここのメイン処理は実行されない")
}

func main3() {
	go sub() // sub関数の中でサブの無限ループ
	for {
		//メインの無限ループ
	}
}

// 実行すると両方のループが並行してGO！

func init() {
	// パッケージ変数を使ったりして初期処理
}

func mainSample() {
	// initのあとに実行
}
