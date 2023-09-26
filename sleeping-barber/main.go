package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// 1 人の理容師、1 つの理容椅子、そして待っている顧客用に n 個の椅子 (n は 0) を備えた待合室がある理髪店
// 次のルールが適用される
// - 客がいないと床屋は椅子で眠ってしまう
// - 理容師が寝ている場合、顧客は理容師を起こさなければならない
// - 理容師が作業中に顧客が到着した場合、椅子がすべて埋まっていればその顧客は立ち去り、椅子が空いていれば空いた椅子に座る
// - 散髪後、理髪師は待合室をチェックして待っている客がいるかどうかを確認し、誰もいない場合は眠りにつく

// 生じうる問題
// - 待合室の確認、入店、待合室の椅子に座るなどのあらゆる動作が、顧客が散髪を待っている間に理容師が寝ているという競合状態を引き起こす
// - 待合室に空席が 1 つしかないときに 2 人の顧客が同時に到着し、2 人とも 1 つの椅子に座りたい場合、最初に椅子に着いた人だけが座ることになる

// variables
var seetingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed a random number generator
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("--------------------------------------------------------")

	clientChan := make(chan string, seetingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seetingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The Shop is open for the day!")

	shop.addBarbers("Frank")
	shop.addBarbers("Gerad")
	shop.addBarbers("Milton")
	shop.addBarbers("Susan")
	shop.addBarbers("Kelly")
	shop.addBarbers("Pat")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMillSeconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillSeconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed
}
