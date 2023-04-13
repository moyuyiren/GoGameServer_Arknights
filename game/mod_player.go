package game

import (
	"GoGameServer_Arknights/csvs"
	"GoGameServer_Arknights/manage"
	"encoding/json"
	"errors"
	"fmt"
)

type ModPlayer struct {
	UserId      int    //玩家Uid
	Icon        int    //头像
	Name        string //名字
	Sign        string //签名
	PlayerLevel int    //玩家等级
	PlayerExp   int    //玩家经验
	Birth       int    //入职日期
	//=====================简单处理需要干员模块支持=====================
	AssistTeam []*AssisTeam //助战编队
	Assistant  int          //助理
	Hire       int          //雇佣人数
	//============================后勤模块===========================
	Furniture int //家具
	//======================未知的字段-举例可能还有很多=================
	IsProhibit int //封禁
	IsGM       int //是否是Gm
}

// SetIcon 设置头像
func (self *ModPlayer) SetIcon(iconID int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconID) {
		//通知客户端操作非法
		return
	}
	player.ModPlayer.Icon = iconID
	fmt.Println("当前图标：", player.ModPlayer.Icon)
}

// SetName 设置Name名字
func (self *ModPlayer) SetName(Name string, player *Player) {
	//违禁词识别  建议直接封号，没有花里胡哨
	if manage.GetManageBanWord().IsBanWord(Name) {
		return
	}
	//执行逻辑
	player.ModPlayer.Name = Name
	fmt.Println("当前Name：", player.ModPlayer.Name)
}

// SetName 设置Sign签名
func (self *ModPlayer) SetSign(Sign string, player *Player) {
	//违禁词识别  建议直接封号，没有花里胡哨
	if manage.GetManageBanWord().IsBanWord(Sign) {
		return
	}
	//执行逻辑
	player.ModPlayer.Name = Sign
	fmt.Println("当前Name：", player.ModPlayer.Name)
}

// AddExp 添加经验  GM||关卡完成后||任务完成后
func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp
	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		if config.PlayerExp == 0 {
			break
		}
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}
	fmt.Println("Now Level ", self.PlayerLevel, "当前经验", self.PlayerExp)
}

// SetAssistTeam 设置助战
func (self *ModPlayer) SetAssistTeam(AssistTeamStr string) error {
	AssisTeamBean := NewAssisTeamBean()
	err := json.Unmarshal([]byte(AssistTeamStr), AssisTeamBean)
	if err != nil {
		return err
	}
	for _, iteam := range self.AssistTeam {
		if iteam.RoleId == AssisTeamBean.RoleId {
			return errors.New("重复设置")
		}
	}
	self.AssistTeam = append(self.AssistTeam, AssisTeamBean)
	return nil

}

// GetAssistTeam 获取助战  todo
func (self *ModPlayer) GetAssistTeam(iconID int, player *Player, playerUid string) {
}

// SetProhibit 设置封禁
func (self *ModPlayer) SetProhibit(prohibit int) {
	self.IsProhibit = prohibit
}

// SetIsGm 设置为Gm用户
func (self *ModPlayer) SetIsGm(isGm int) {
	self.IsGM = isGm
}
