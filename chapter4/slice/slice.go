package main

import "fmt"

func main() {
	var s []int         // 型としては配列と違い[]内に数を書かない。
	s = make([]int, 5)  // [0 0 0 0 0]
	s[5] = 5            // 代入も同じ
	s[6] = 6            // 要素数を超えているのでパニック発生
	fmt.Println(len(s)) // len関数で要素数の5が返る

	s2 := make([]int, 5, 10) // 6～10個目の値を格納できる領域があらかじめ確保される
	fmt.Println(cap(s2))     // cap関数で確保された領域10が返る

	s3 := []int{1, 2, 3, 4, 5} // 要素数5、容量5で配列と似たように生成するリテラル
	s4 := s3[0:2]              // [1,2] と部分的に新しいスライス。最初だけ、最後だけ、全部なども:で指定可能
	// 文字列抽出にも使えるがバイト列なので、マルチバイト文字には使えない
	sAlphabet := "ABC"[0:3] // [A B]
	sMoji := "あいうえお"[0:3]   // うまくいかない
	fmt.Println(sAlphabet, sMoji)

	// append関数で末尾に複数要素追加。Go側でメモリ領域を適宜増やす
	s5 := append(s4, 3, 4) // [1,2,3,4]
	fmt.Println(s5)
	c1 := []int{1, 2, 3} // [1, 2, 3]
	c2 := []int{10, 11}  //「10, 11]

	//copy関数で第1引数のsliceに塗りつぶしコピー。コピー数が返る
	n := copy(c1, c2)    // n=2 c1=[10,11,3]
	c3 := []int{0, 2, 5} //完全スライス式。[10,11,容量確保,容量確保,容量確保]
	fmt.Println(n, c3)

	sum(1, 2)    // 戻り値3 引数いくつでも
	sum(1, 2, 3) // 戻り値6
	sliceArg := []int{1, 2, 3}
	sum(sliceArg...) // 引数にスライスを渡すときはスライス変数名+...で渡せる

	countryMap := make(map[int]string) // この書き方が独特
	countryMap[1] = "Japan"
	countryMap[2] = "China"
	countryMap[1] = "Korea" // キーが重複すると上書き
	//countryMap[3] = ...

	//リテラルでも書ける。複数行の場合は最後にカンマ忘れずに
	fruitsMap := map[int]string{1: "Apple", 2: "Peach"}
	// [キー:要素] の要素に配列やスライス、さらにマップを入れたりもできる。
	// キーが登録されていないので要素が取れず、fruitは空文字
	fruit := fruitsMap[3]
	fmt.Println(fruit)

	if _, ok := fruitsMap[2]; ok {
		// fruitsMap[2]が取れた時の処理。変数名はokを使うのがGoのお作法
	}
	for k, v := range fruitsMap {
		fmt.Println(k, v) // 配列やスライス同様、1件1件にアクセスできる
	}
	fmt.Println(len(fruitsMap)) // スライスと同じく、要素数の2が返る。
	// ちなみにスライスと違って容量を得るcap関数は使えない。
	delete(fruitsMap, 1) //キーの値で取り除いて、 [2: Peach] のみになる

	// 要素数に応じた初期スペースを第2引数で指定
	largeLargeFruitsMap := make(map[int]string, 100)
	// なおこれはメモリ領域確保のヒント的なもので、スライスの容量とはちょっと違う。
	fmt.Println(largeLargeFruitsMap)

}

// ...でx個のint型の可変長引数が取れる
func sum(s ...int) int {
	// rangeを使い配列とスライスの各要素に簡単にアクセス
	for _, v := range s {
		// 取り出して何か処理
		fmt.Println(v)
	}
	return 0
}
