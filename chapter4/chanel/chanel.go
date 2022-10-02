package main

import "fmt"

func receiveSample(ch chan int) {
	// 無限ループでchから受け取った整数をfmt.Printlnする
}

func main() {
	var ch chan int          // int型のチャネル。chanをつける
	var ch4r <-chan int      //受信専用 常に矢印は左向き
	var ch4s chan<- chan int //送信専用

	ch = make(chan int, 4) // 組み込み関数makeで生成、バッファが格納領域のサイズ
	ch <- 999              // チャネルに向かって整数を送信
	i := <-ch              // チャネルから受信

	ch = make(chan int)
	go receiveSample(ch) // ゴルーチン起動！
	// forループでchに向かって整数を送信
	// 数字がずっと出力される
	fmt.Println(ch4r, ch4s)

	// バッファが0か中が空のチャネルから受信、バッファに空きがないチャネルへの送信でエラー
	// Fatal error: all goroutines are asleep - deadlock!

	chFruits := make(chan string, 2)
	chFruits <- "Apple"
	chFruits <- "Cherry"
	result1 := len(chFruits) // 2 データの個数が入る。cap(chFruits)でバッファサイズが取れる
	fmt.Println(result1)

	ch1 := make(chan int, 2)
	close(ch1)
	//クローズ済みのチャネルに送信はできない。
	//ch <- 1 // panic: send on closed channel
	fmt.Println(<-ch1) // クローズ済みでもチャネルからの受信はできる。
	i, ok := <-ch1
	// チャネル内のバッファが空 && クローズ済みだと戻り値2つめがfalseになる。
	fmt.Println(i, ok) // 0 false

	ch4for := make(chan int, 2)
	ch4for <- 1
	ch4for <- 2
	close(ch4for) // クローズしていないと最後の受信分の後で deadlock!
	for i := range ch4for {
		fmt.Println(i) // 1つづつ取り出せる
	}

	ch4select1 := make(chan int, 1)
	ch4select2 := make(chan string, 1)
	ch4select1 <- 123
	ch4select2 <- "もじれつ"
	// このselect文を実行すると、Goランタイムはどのcase節を実行するか
	// ランダムに選択する。実行のたびに変わる。
	select {
	case v1 := <-ch4select1:
		// 来たらv1には123が入っている。
		fmt.Println(v1)
	case v2 := <-ch4select2:
		// 来たらv2には"もじれつ" が入っている。
		fmt.Println(v2)
	default:
		// このコードではどちらにも当てはまらない例は来ない。
	}
}
