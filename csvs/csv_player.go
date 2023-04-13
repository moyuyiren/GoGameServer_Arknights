package csvs

import (
	"GoGameServer_Arknights/utils"
)

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
}

var ConfigPlayerLevelSlice []*ConfigPlayerLevel

func init() {
	utils.GetCsvUtilMgr().LoadCsv("PlayerLevel", &ConfigPlayerLevelSlice)
	return
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level >= len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level]

}
