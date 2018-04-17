package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class UpdateTaskRequest
{
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "ExecutionState", required = false)
    private String executionState;
    @JsonProperty(value = "ExpectedExecutionMs", required = false)
    private long expectedExecutionMs;
    @JsonProperty(value = "Percentage", required = false)
    private int percentage;
    @JsonProperty(value = "ExecutionResult", required = false)
    private UpdateExecutionResultRequest executionResult;

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }

    public String getExecutionState()
    {
        return executionState;
    }

    public void setExecutionState(String executionState)
    {
        this.executionState = executionState;
    }

    public long getExpectedExecutionMs()
    {
        return expectedExecutionMs;
    }

    public void setExpectedExecutionMs(long expectedExecutionMs)
    {
        this.expectedExecutionMs = expectedExecutionMs;
    }

    public int getPercentage()
    {
        return percentage;
    }

    public void setPercentage(int percentage)
    {
        this.percentage = percentage;
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
