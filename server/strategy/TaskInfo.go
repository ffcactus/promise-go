package strategy

import (
	"promise/common/object/constValue"
	"promise/server/object/model"
	taskDto "promise/task/object/dto"
)

func init() {
	refreshTaskTotalTime = 0
	for i := range ServerTaskRefreshStepLIST {
		refreshTaskTotalTime += ServerTaskRefreshStepLIST[i].ExpectedExecutionMs
	}
}

var (
	refreshTaskTotalTime int

	// ServerTaskRefresh is an ID.
	ServerTaskRefresh = "Server.Task.Refresh"

	// ServerRefreshTaskStepIDPower is an ID.
	ServerRefreshTaskStepIDPower = "Server.Refresh.Power"
	// ServerRefreshTaskStepIDThermal is an ID.
	ServerRefreshTaskStepIDThermal = "Server.Refresh.Thermal"
	// ServerRefreshTaskStepIDBoards is an ID.
	ServerRefreshTaskStepIDBoards = "Server.Refresh.OemHuaweiBoards"
	// ServerRefreshTaskStepIDNetworkAdapters is an ID.
	ServerRefreshTaskStepIDNetworkAdapters = "Server.Refresh.NetworkAdapters"
	// ServerRefreshTaskStepIDDrives is an ID.
	ServerRefreshTaskStepIDDrives = "Server.Refresh.Drives"
	// ServerRefreshTaskStepIDPCIeDevices is an ID.
	ServerRefreshTaskStepIDPCIeDevices = "Server.Refresh.PCIeDevices"
	// ServerRefreshTaskStepIDProcessors is an ID.
	ServerRefreshTaskStepIDProcessors = "Server.Refresh.Processors"
	// ServerRefreshTaskStepIDMemory is an ID.
	ServerRefreshTaskStepIDMemory = "Server.Refresh.Memory"
	// ServerRefreshTaskStepIDEthernetInterfaces is an ID.
	ServerRefreshTaskStepIDEthernetInterfaces = "Server.Refresh.EthernetInterfaces"
	// ServerRefreshTaskStepIDNetworkInterfaces is an ID.
	ServerRefreshTaskStepIDNetworkInterfaces = "Server.Refresh.NetworkInterfaces"
	// ServerRefreshTaskStepIDStorages is an ID.
	ServerRefreshTaskStepIDStorages = "Server.Refresh.Storages"

	// ServerRefreshTaskStepNamePower is task step name.
	ServerRefreshTaskStepNamePower = "Power"
	// ServerRefreshTaskStepNameThermal is task step name.
	ServerRefreshTaskStepNameThermal = "Thermal"
	// ServerRefreshTaskStepNameBoards is task step name.
	ServerRefreshTaskStepNameBoards = "OemHuaweiBoards"
	// ServerRefreshTaskStepNameNetworkAdapters is task step name.
	ServerRefreshTaskStepNameNetworkAdapters = "NetworkAdapters"
	// ServerRefreshTaskStepNameDrives is task step name.
	ServerRefreshTaskStepNameDrives = "Drives"
	// ServerRefreshTaskStepNamePCIeDevices is task step name.
	ServerRefreshTaskStepNamePCIeDevices = "PCIeDevices"
	// ServerRefreshTaskStepNameProcessors is task step name.
	ServerRefreshTaskStepNameProcessors = "Processors"
	// ServerRefreshTaskStepNameMemory is task step name.
	ServerRefreshTaskStepNameMemory = "Memory"
	// ServerRefreshTaskStepNameEthernetInterfaces is task step name.
	ServerRefreshTaskStepNameEthernetInterfaces = "EthernetInterfaces"
	// ServerRefreshTaskStepNameNetworkInterfaces is task step name.
	ServerRefreshTaskStepNameNetworkInterfaces = "NetworkInterfaces"
	// ServerRefreshTaskStepNameStorages is task step name.
	ServerRefreshTaskStepNameStorages = "Storages"

	// ServerTaskRefreshStepPower is a task step.
	ServerTaskRefreshStepPower = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPower,
		Name:                ServerRefreshTaskStepNamePower,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepThermal is a task step.
	ServerTaskRefreshStepThermal = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDThermal,
		Name:                ServerRefreshTaskStepNameThermal,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepBoards is a task step.
	ServerTaskRefreshStepBoards = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDBoards,
		Name:                ServerRefreshTaskStepNameBoards,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepNetworkAdapters is a task step.
	ServerTaskRefreshStepNetworkAdapters = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkAdapters,
		Name:                ServerRefreshTaskStepNameNetworkAdapters,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepDrives is a task step.
	ServerTaskRefreshStepDrives = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDDrives,
		Name:                ServerRefreshTaskStepNameDrives,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepPCIeDevices is a task step.
	ServerTaskRefreshStepPCIeDevices = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDPCIeDevices,
		Name:                ServerRefreshTaskStepNamePCIeDevices,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepProcessors is a task step.
	ServerTaskRefreshStepProcessors = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDProcessors,
		Name:                ServerRefreshTaskStepNameProcessors,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepMemory is a task step.
	ServerTaskRefreshStepMemory = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDMemory,
		Name:                ServerRefreshTaskStepNameMemory,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepEthernetInterfaces is a task step.
	ServerTaskRefreshStepEthernetInterfaces = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDEthernetInterfaces,
		Name:                ServerRefreshTaskStepNameEthernetInterfaces,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepNetworkInterfaces is a task step.
	ServerTaskRefreshStepNetworkInterfaces = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDNetworkInterfaces,
		Name:                ServerRefreshTaskStepNameNetworkInterfaces,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepStorages is a task step.
	ServerTaskRefreshStepStorages = taskDto.PostTaskStepRequest{
		MessageID:           &ServerRefreshTaskStepIDStorages,
		Name:                ServerRefreshTaskStepNameStorages,
		ExpectedExecutionMs: 4000,
	}

	// ServerTaskRefreshStepLIST is a task step list.
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

// CreateRefreshTaskRequest will return a refresh server task request.
func CreateRefreshTaskRequest(server *model.Server) *taskDto.PostTaskRequest {
	var request taskDto.PostTaskRequest

	request.MessageID = &ServerTaskRefresh
	request.Name = "Refresh Server"
	request.Description = "Refresh server resources and re-configure it."
	request.CreatedByName = "CreatedByName???"
	request.CreatedByURI = "CreatedByURI???"
	request.TargetName = server.Name
	request.TargetURI = constValue.ToServerURI(server.ID)
	request.TaskSteps = ServerTaskRefreshStepLIST
	return &request
}
