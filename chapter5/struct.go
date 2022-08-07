package main

type User struct {
	Name string
	Age int
}
type Users []*User

user1 := User{Name: "user1", Age: 10}
user2 := User{Name: "user2", Age: 20}
users1 := Users[]
// 複数追加できる。ポインタを使うのに注意
users.append(users, &user1, &user2) 
for _1, u := range users1 {
fmt.println(u) // &{user1, 10} ...
}

users2 := make([]*User, 0) // 中身0件で作れる
users3 := make(Users, 0) // 同じこと

// キーが数値、値が構造体のマップ
userMap := map[int]User {
	1: {Name: "user1", Age: 10},
	2: {Name: "user2", Age: 20},
}

type User struct {
	Name string `名前でウィス`
	Age int `年でウィス`
}


type error interface {
	Error() string
}

// 独自定義のエラー型
type MyError struct {
	Message string
	ErrorCode string
}
// errorインターフェースで指定された「メソッド」を実装
func (e *MyError) Error() string {
	return "特製のエラーメッセージを返すでウィス" + e.Message
}
// 独自定義のエラー型を返す関数 しかし戻り値の型はerrorインターフェース
func RaiseMyError() error {
	return &MyError{Message: "エラーだニャン", ErrorCode: NKB48}
}

err := RaiseMyError() // ここでは変数errはerror型と扱われる
err.Error() // 呼べる。「特製の〜」が返る
// .(ポインタ+型)の型アサーション使うと本来の型に戻る
e, ok := err.(*MyError)
if ok {
	// ここでeは構造体MyError型として扱われるのでフィールドにアクセス可能。
	e.ErrorCode // NKB48 
	// VSCodeだとドットを打つとちゃんと補完が出る。素晴らしい！
}

// スペック表示ができることを示すインターフェース
type SpecShowable interface {
	ShowSpec() string
}

// 構造体パーソンと、
// SpecShowableインターフェースで実装が必須になるメソッドを定義
type Person struct {
	Name string
	Age int
}
func (p *Person) ShowSpec() string {
	return "スペック？普通なんだけど..."
}

// 構造体ロボットと、
// SpecShowableインターフェースで実装が必須になるメソッドを定義
type Robot struct {
	Number string
	Model string
}
func (r *Robot) ShowSpec() string {
	return "ナンバー:" + r.Number + "モデル:" + r.Model
}

// 型が違っても、全てSpecShowableインターフェース型で揃ったスライスとして扱える
// 構造体なのでアドレス演算子&を最初から全部つける
mens := []SpecShowable {
	&Person{Name: "ケイタ", Age: 10}
	&Robot{Number: "ROBONYAN", Model:"F型"}
}

// 1件1件がSpecShowableインターフェース型なので、共通したメソッドが呼べる！
for _, v := range vs {
// スペックの文字列が返るので表示できる
fmt.Println(v.ShowSpec())
// 2件のメソッド実行結果が順に出力:
// スペック？普通なんだけど...
// ナンバー:ROBONYAN モデル:F型
}