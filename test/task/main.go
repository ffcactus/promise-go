package main

import (
	"fmt"
	"promise/sdk/task"
	"promise/task/object/dto"
	"time"
)

var (
	concurrent = make(chan bool, 2)
	done       = make(chan string, 100)
)

func test(name string) {
	var (
		create     = dto.PostTaskRequest{}
		update     = dto.UpdateTaskRequest{}
		percentage = []uint32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	)
	create.Name = name
	create.TaskSteps = []dto.PostTaskStepRequest{dto.PostTaskStepRequest{
		Name:                "Step",
		ExpectedExecutionMs: 1000,
	}}
	response, errorResp, err := task.CreateTask(&create)
	if errorResp != nil || err != nil {
		fmt.Printf("Create %s failed.\n", name)
		done <- name
		<-concurrent
		return
	}
	fmt.Printf("Create %s.\n", name)
	for _, v := range percentage {
		update.Percentage = &v
		_, errorResp, err := task.UpdateTask(response.ID, &update)
		if errorResp != nil || err != nil {
			fmt.Printf("Update %s to %d failed.\n", name, *update.Percentage)
			done <- name
			<-concurrent
			return
		}
		fmt.Printf("Update %s to %d.\n", name, *update.Percentage)
	}
	fmt.Printf("%s before done.\n", name)
	done <- name
	fmt.Printf("%s before concurrent.\n", name)
	<-concurrent
	fmt.Printf("%s Done.\n", name)
}

func main() {
	var (
		instance = 10
		index    = 0
	)
	for {
		select {
		case concurrent <- true:
			go test(fmt.Sprintf("Task%d", index))
			index++
		default:
			fmt.Printf("Wait.\n")
			time.Sleep(time.Duration(1) * time.Second)
		}
		if index == 10 {
			break
		}
	}
	for i := 0; i < instance; i++ {
		name := <-done
		fmt.Printf("%s Done.\n", name)
	}
	fmt.Printf("Done.\n")
}
