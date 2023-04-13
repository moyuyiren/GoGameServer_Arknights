package game

type ModIcon struct {
}

// 是否有这个头像
func (self *ModIcon) IsHasIcon(iconID int) bool {
	return true
}
