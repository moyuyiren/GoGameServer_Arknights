package game

type AssisTeam struct {
	Table     int //槽位
	RoleId    int //干员id
	RoleStar  int //干员星级
	RoleLevel int //干员等级
	RoleSP    int //干员技能
}

func NewAssisTeamBean() *AssisTeam {
	AssisTeamBean := AssisTeam{
		Table:     0,
		RoleId:    0,
		RoleStar:  0,
		RoleLevel: 0,
		RoleSP:    0,
	}
	return &AssisTeamBean
}
