package game

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModUniqueTask *ModUniqueTask
}

func NewTestPlayer() *Player {
	//**********************模块初始化***********************
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModUniqueTask = new(ModUniqueTask)
	//******************数据库进行数据初始化********************

	//********************模拟数据初始化***********************
	player.ModPlayer.Icon = 0000
	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.PlayerExp = 0
	return player
}

// 测试客户端
func (self *Player) Testfunction() {
	self.RecvSetName("我登录了")
}

// Recv 对外接口
// RecvSetIcon 设置玩家头像
func (self *Player) RecvSetIcon(iconID int) {
	self.ModPlayer.SetIcon(iconID, self)
}

// RecvSetName 设置玩家Name
func (self *Player) RecvSetName(Name string) {
	self.ModPlayer.SetName(Name, self)
}

// RecvSetSign 设置玩家Sign
func (self *Player) RecvSetSign(Sign string) {
	self.ModPlayer.SetSign(Sign, self)
}

// RecvSetAssistTeam 设置助战阵容 需要前端返回用户设置的值
func (self *Player) RecvSetAssistTeam(AssistTeamStr string) {
	self.ModPlayer.SetAssistTeam(AssistTeamStr)
}
