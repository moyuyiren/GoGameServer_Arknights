package main

import (
	"GoGameServer_Arknights/csvs"
	"GoGameServer_Arknights/game"
	"GoGameServer_Arknights/manage"
	"time"
)

func main() {
	//加载配置
	csvs.CheckLoad()
	//公共管理模块
	go manage.GetManageBanWord()
	//监听客户端模块

	//模拟登陆，每个玩家拥有自己的协程
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			player := game.NewTestPlayer()
			go player.Testfunction()
		}
	}
	return
}
