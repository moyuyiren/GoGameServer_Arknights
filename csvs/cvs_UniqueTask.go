package csvs

import (
	"GoGameServer_Arknights/utils"
)

type ConfigUniqueTask struct {
	TaskId        int `json:"TaskId"`        //任务编号
	OpenTaskLevel int `json:"OpenTaskLevel"` //任务阶段
	TaskType      int `json:"TaskType"`      //任务类型
	Condition     int `json:"Condition"`     //任务条件
}

var ConfigUniqueTaskMap map[int]*ConfigUniqueTask

func init() {
	ConfigUniqueTaskMap = make(map[int]*ConfigUniqueTask)
	utils.GetCsvUtilMgr().LoadCsv("UniqueTask", &ConfigUniqueTaskMap)
	return
}
