package csvs

import "GoGameServer_Arknights/utils"

type ConfigIcon struct {
	IconID int `json:"iconId"`
	Check  int `json:"Check"`
}

var (
	ConfigIconMap         map[int]*ConfigIcon
	ConfigIconMapByRoleId map[int]*ConfigIcon
)

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)
	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)
	ConfigIconMapByRoleId = make(map[int]*ConfigIcon)
	for _, v := range ConfigIconMap {
		ConfigIconMapByRoleId[v.Check] = v
	}
	return
}

func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func GetIconConfigByRoleId(roleId int) *ConfigIcon {
	return ConfigIconMapByRoleId[roleId]
}
