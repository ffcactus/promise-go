package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class UpdateTaskStepRequest
{
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "ExecutionState", required = false)
    private String executionState;
    @JsonProperty(value = "executionResult", required = false)
    private UpdateExecutionResultRequest executionResult;

    public String getName()
    {
        return name;
    }

    public void setName(String name)
    {
        this.name = name;
    }

    public String getExecutionState()
    {
        return executionState;
    }

    public void setExecutionState(String executionState)
    {
        this.executionState = executionState;
    }

    public UpdateExecutionResultRequest getExecutionResult()
    {
        return executionResult;
    }

    public void setExecutionResult(UpdateExecutionResultRequest executionResult)
    {
        this.executionResult = executionResult;
    }
}
