// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	p := player{tickets: 10, coin: 100}

	n := inputN(&p)
	results, summary := drawN(&p, n)

	fmt.Println(results)
	fmt.Println(summary)
}

func inputN(p *player) int {

	max := p.drawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		// TODO: nが0より大きくmax以下である場合はbreak
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

// -- gacha.goに移す ここから --

func drawN(p *player, n int) ([]*card, map[rarity]int) {
	p.draw(n)

	results := make([]*card, n)
	summary := make(map[rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &card{rarity: rarityN, name: "スライム"}
	case num < 95:
		return &card{rarity: rarityR, name: "オーク"}
	case num < 99:
		// TODO: rarityフィールドがraritySRで
		// nameフィールドが"ドラゴン"であるcard構造体のポインタを返す
		return &card{rarity: raritySR, name: "ドラゴン"}
	default:
		return &card{rarity: rarityXR, name: "イフリート"}
	}
}

// -- gacha.goに移す ここまで --
