package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class GetTaskStepResponse
{
    @JsonProperty(value = "MessageID", required = false)
    private String messageID;
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "ExpectedExecutionMs", required = true)
    private long expectedExecutionMs;
    @JsonProperty(value = "ExecutionState", required = true)
    private String executionState;
    @JsonProperty(value = "ExecutionResult", required = true)
    private ExecutionResultResponse executionResult;

    public String getMessageID()
    {
        return messageID;
    }

    public void setMessageID(String messageID)
    {
        this.messageID = messageID;
    }

    public String getName()
    {
        return name;
    }

    public void setName(String name)
    {
        this.name = name;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }

    public long getExpectedExecutionMs()
    {
        return expectedExecutionMs;
    }

    public void setExpectedExecutionMs(long expectedExecutionMs)
    {
        this.expectedExecutionMs = expectedExecutionMs;
    }

    public String getExectuionState()
    {
        return executionState;
    }

    public void setExectuionState(String exectuionState)
    {
        this.executionState = exectuionState;
    }

    public ExecutionResultResponse getExecutionResult()
    {
        return executionResult;
    }

    public void setExecutionResult(ExecutionResultResponse executionResult)
    {
        this.executionResult = executionResult;
    }
}
