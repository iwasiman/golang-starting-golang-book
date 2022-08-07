package main

import "fmt"

//変数宣言では型の前に*をつける。*int型となる。ポインタ型の変数初期値はnil。
var p *int
var p4array *[3]int 

var n int = 100
//アドレス演算子"&"をつけると実際の番地を指す。
fmt.Println(n, &n) // n: 100 &n: 0xな番地
DoubleFunc(n) //引数を2倍する関数を呼んでも、基本型は値渡しなので変わらない。

var p *int = &n
//p := &n // 暗黙的な宣言で同じ
//変数の前に*をつけると番地が示すデータの実体になる。「デリファレンス」という。
fmt.Println(p, *p) // p:元から&nが入っているので、0xな番地nのまま *p:実体なので100 
*p = 300 //ポインタでデリファレンスし、指している実体を書き換える
fmt.Println(n, &n) // n:300 &n:0xな番地 // 値が書き換わる。番地は同じ。

// 引数にポインタ、関数内でもポインタでアクセスする関数を設定
DoubleFunc4Pointer(p *int) {
	*i = *i * 2
}

// nには300が入っているので、アドレスを渡せば実体を書き換えられる。
DoubleFunc4Pointer(&n)
fmt.Println(n, &n) // n:600 &n:0xな番地のまま // 値が書き換わる。番地は同じ。
fmt.Println(*p, p) // *p:600 p:0xな番地のまま // pはnを指しているので値、番地は同じ。

// ポインター変数を渡すときは変数名を渡すだけ。
DoubleFunc4Pointer(p)
fmt.Println(n, &n) // n:1200 &n:0xな番地のまま // pが指しているnが書き換わる。番地は同じ。
fmt.Println(*p, p) // *p:1200 p:0xな番地のまま // nを指しているので同じ。


// 配列へのポインタ型
drinks := [2]string{"Milk", "Juice"}
pd := &drinks
fmt.Println(drinks[0]) // "Milk"
fmt.Println(pd[0]) // C的には(*pd)[0]だが、コンパイラが検知して展開してくれる

// 文字列
s := "ABCDE"
&s // ポインタで番地が出る
s[0] // インデックス参照だがbyte型、マルチバイト文字には使えない
&s[0] // ポインタ参照するとコンパイルエラー。
// Goでは文字列は不変(immutable)で、破壊的変更できないようになっている。
// また関数の引数にstring型を使って操作しても、元から参照型で文字列の実体は新しい領域にコピーされない。
// したがって、Goで*string型で文字列のポインターを使うシーンは基本的にない。




type MyInt int // エイリアスが定義できる。
// 元のint型にも型変換できる。なおエイリアス型同士では変換できない。
var mi1 MyInt = 5 
type Strings []string //String{"aaa", "bbb"}
type AreaMap map[string][2]float64 // AreaMap{"SomeCity", {12.345, 23.456}}

type Point struct {
	X int
	Y int
}
pt1 := Point{1, 2} // フィールドの順番に代入される
pt2 := Point{X: 1, Y: 2} // こちらの方式がよい
pt1.X = 100 // ドットで代入も普通にできる

type P struct {
desc string
Point // "Point Point" となる場合は型名を省略できる
}
pSample := P{desc: "説明", Point: Point{X: 1, Y: 2}}
// 下はpSample.Point.X の途中のPoint.が省略できる。
pSample.X // これでアクセス可能、VSCodeでも補完が出る。

func swap(p *Point) {
// Point構造体の中身を入れ替える処理
}
swap(pt1) // 構造体は値渡しなのでpt1は変化しない、この呼び方はほぼ使わない
swap(&pt1) // ポインタを渡すと参照渡し、構造体の中が変わる
pByPointer := &Point{X:100, y:200} // 構造体は最初からポインタを生成して使うのが良い。
pByPointer2 := new(Point) // new演算子を使って書くやり方
swap(pByPointer2)

// Point構造体用のメソッドの定義
func (p *Point) showValues() {
	// レシーバーがp Point でも p *Pointでも、関数内ではp.で呼び出せる。
	// C/C++の世界では . と -> 両方が出てくるのより分かりやすい。
	fmt.Println("X:::" p.X, "Y:::", p.x)
}

pt1.showValues() // インスタンスメソッド的に呼び出せる。
// Point型でも*Point型でも呼び出し側では同じ。ドットで補完が出る。
pByPointer2.showValues()

// 慣例で"New"+構造体名で「型のコンストラクタ」という関数を作る。
// 戻り値は型のポインタ型を使う。
func NewPoint(x int, y int) *Pointer {
return &Pointer{X: x, Y: y}
}
pFromConstructor := newPoint(500, 600) // *Point型で生成される

