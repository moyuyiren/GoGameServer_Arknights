package game

const TASK_STATE_FINISH = 1

type TaskInfo struct {
	TaskId int
	State  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
}

// IsTaskFinish 任务完成确认
func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {
	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.State == TASK_STATE_FINISH
}
