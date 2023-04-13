package manage

import (
	"GoGameServer_Arknights/csvs"
	"fmt"
	"regexp"
)

var manageBanWord *MangeBanWord

type MangeBanWord struct {
	BanWordBase  []string //配置生成
	BanWordExtra []string //更新
}

func GetManageBanWord() *MangeBanWord {
	if manageBanWord == nil {
		manageBanWord = new(MangeBanWord)
		manageBanWord.BanWordBase = csvs.GetBanWordBase()
		manageBanWord.BanWordExtra = []string{"秒图", "刷关"}
	}
	return manageBanWord
}

func (self *MangeBanWord) IsBanWord(txt string) bool {
	for _, s := range self.BanWordBase {
		match, _ := regexp.MatchString(s, txt)
		fmt.Println(match, s)
		if match {
			return match
		}
	}

	for _, s := range self.BanWordExtra {
		match, _ := regexp.MatchString(s, txt)
		fmt.Println(match, s)
		if match {
			return match
		}
	}
	return false
}
