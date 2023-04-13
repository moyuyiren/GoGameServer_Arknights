package test

import (
	"GoGameServer_Arknights/game"
	"fmt"
	"testing"
	"time"
)

// 测试修改Icon
func TestSetIcon(t *testing.T) {
	player := game.NewTestPlayer()
	player.RecvSetIcon(1111)
}

// 测试修改名字
func TestSetName(t *testing.T) {
	player := game.NewTestPlayer()

	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%3 == 0 {
				player.RecvSetName("外挂")
				time.Sleep(time.Second)
			} else {
				fmt.Println("好人")
				time.Sleep(time.Second)
			}
		}
	}
}
