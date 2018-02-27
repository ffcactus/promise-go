package strategy

import (
	// commonDto "promise/common/object/dto"
	"promise/server/object/model"
	taskDto "promise/task/object/dto"
	// taskModel "promise/task/object/model"
)

func init() {
	refreshTaskTotalTime = 0
	for i, _ := range ServerTaskRefreshStepLIST {
		refreshTaskTotalTime += ServerTaskRefreshStepLIST[i].ExpectedExecutionMs
	}
}

var (
	refreshTaskTotalTime int

	// Task MessageID
	SERVER_TASK_ID_REFRESH = "Server.Task.Refresh"

	// TaskStep MessageID
	ServerRefreshTaskStepIDPower              = "Server.Refresh.Power"
	ServerRefreshTaskStepIDThermal            = "Server.Refresh.Thermal"
	ServerRefreshTaskStepIDBoards             = "Server.Refresh.OemHuaweiBoards"
	ServerRefreshTaskStepIDNetworkAdapters    = "Server.Refresh.NetworkAdapters"
	ServerRefreshTaskStepIDDrives             = "Server.Refresh.Drives"
	ServerRefreshTaskStepIDPCIeDevices        = "Server.Refresh.PCIeDevices"
	ServerRefreshTaskStepIDProcessors         = "Server.Refresh.Processors"
	ServerRefreshTaskStepIDMemory             = "Server.Refresh.Memory"
	ServerRefreshTaskStepIDEthernetInterfaces = "Server.Refresh.EthernetInterfaces"
	ServerRefreshTaskStepIDNetworkInterfaces  = "Server.Refresh.NetworkInterfaces"
	ServerRefreshTaskStepIDStorages           = "Server.Refresh.Storages"

	// Name
	ServerRefreshTaskStepNamePower              = "Power"
	ServerRefreshTaskStepNameThermal            = "Thermal"
	ServerRefreshTaskStepNameBoards             = "OemHuaweiBoards"
	ServerRefreshTaskStepNameNetworkAdapters    = "NetworkAdapters"
	ServerRefreshTaskStepNameDrives             = "Drives"
	ServerRefreshTaskStepNamePCIeDevices        = "PCIeDevices"
	ServerRefreshTaskStepNameProcessors         = "Processors"
	ServerRefreshTaskStepNameMemory             = "Memory"
	ServerRefreshTaskStepNameEthernetInterfaces = "EthernetInterfaces"
	ServerRefreshTaskStepNameNetworkInterfaces  = "NetworkInterfaces"
	ServerRefreshTaskStepNameStorages           = "Storages"

	// Refresh TaskSteps
	ServerTaskRefreshStepPower = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPower,
		Name:                ServerRefreshTaskStepNamePower,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepThermal = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDThermal,
		Name:                ServerRefreshTaskStepNameThermal,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepBoards = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDBoards,
		Name:                ServerRefreshTaskStepNameBoards,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepNetworkAdapters = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkAdapters,
		Name:                ServerRefreshTaskStepNameNetworkAdapters,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepDrives = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDDrives,
		Name:                ServerRefreshTaskStepNameDrives,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepPCIeDevices = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPCIeDevices,
		Name:                ServerRefreshTaskStepNamePCIeDevices,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepProcessors = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDProcessors,
		Name:                ServerRefreshTaskStepNameProcessors,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepMemory = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDMemory,
		Name:                ServerRefreshTaskStepNameMemory,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepEthernetInterfaces = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDEthernetInterfaces,
		Name:                ServerRefreshTaskStepNameEthernetInterfaces,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepNetworkInterfaces = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkInterfaces,
		Name:                ServerRefreshTaskStepNameNetworkInterfaces,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepStorages = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDStorages,
		Name:                ServerRefreshTaskStepNameStorages,
		ExpectedExecutionMs: 4000,
	}
	ServerTaskRefreshStepLIST = []taskDto.PostTaskStepRequest{
		ServerTaskRefreshStepPower,
		ServerTaskRefreshStepThermal,
		ServerTaskRefreshStepBoards,
		ServerTaskRefreshStepNetworkAdapters,
		ServerTaskRefreshStepDrives,
		ServerTaskRefreshStepPCIeDevices,
		ServerTaskRefreshStepProcessors,
		ServerTaskRefreshStepMemory,
		ServerTaskRefreshStepEthernetInterfaces,
		ServerTaskRefreshStepNetworkInterfaces,
		ServerTaskRefreshStepStorages,
	}
)

func CreateRefreshTaskRequest(server *model.Server) *taskDto.PostTaskRequest {
	var request taskDto.PostTaskRequest

	request.MessageID = &SERVER_TASK_ID_REFRESH
	request.Name = "Refresh Server"
	request.Description = "Refresh server resources and re-configure it."
	request.CreatedByName = "CreatedByName???"
	request.CreatedByURI = "CreatedByURI???"
	request.TargetName = server.Name
	request.TargetURI = server.URI
	request.TaskSteps = ServerTaskRefreshStepLIST
	return &request
}
