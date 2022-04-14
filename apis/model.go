// Package apis Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package apis

/**
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

type Model struct {
	taskManager       *TaskManager
	workerManager     *WorkerManager
	constraintManager *ConstraintManager
}

type TaskManager struct {
}

type WorkerManager struct {
}

type ConstraintManager struct {
}

type task struct {
	taskId       string
	taskName     string
	resourceNeed map[string]int
}

type worker struct {
	workerId       string
	workerName     string
	totalResource  map[string]int
	usedResource   map[string]int
	remainResource map[string]int
	deployedTask   map[string]string
}

type constraint struct {
}
