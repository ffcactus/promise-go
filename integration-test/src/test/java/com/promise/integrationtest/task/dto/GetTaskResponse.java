package com.promise.integrationtest.task.dto;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.ResourceResponse;

public class GetTaskResponse extends ResourceResponse
{
    @JsonProperty(value = "MessageID", required = false)
    private String messageID;
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "ExecutionState", required = false)
    private String executionState;
    @JsonProperty(value = "CreatedByName", required = true)
    private String createdByName;
    @JsonProperty(value = "CreatedByURI", required = true)
    private String createdByURI;
    @JsonProperty(value = "TargetName", required = true)
    private String targetName;
    @JsonProperty(value = "TargetURI", required = true)
    private String targetURI;
    @JsonProperty(value = "ExpectedExecutionMs", required = true)
    private long expectedExecutionMs;
    @JsonProperty(value = "Percentage", required = true)
    private int percentage;
    @JsonProperty(value = "CurrentStep", required = true)
    private String currentStep;
    @JsonProperty(value = "TaskSteps", required = true)
    private List<GetTaskStepResponse> taskSteps;
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

    public String getCreatedByName()
    {
        return createdByName;
    }

    public void setCreatedByName(String createdByName)
    {
        this.createdByName = createdByName;
    }

    public String getCreatedByURI()
    {
        return createdByURI;
    }

    public void setCreatedByURI(String createdByURI)
    {
        this.createdByURI = createdByURI;
    }

    public String getTargetName()
    {
        return targetName;
    }

    public void setTargetName(String targetName)
    {
        this.targetName = targetName;
    }

    public String getTargetURI()
    {
        return targetURI;
    }

    public void setTargetURI(String targetURI)
    {
        this.targetURI = targetURI;
    }

    public int getPercentage()
    {
        return percentage;
    }

    public void setPercentage(int percentage)
    {
        this.percentage = percentage;
    }

    public String getCurrentStep()
    {
        return currentStep;
    }

    public void setCurrentStep(String currentStep)
    {
        this.currentStep = currentStep;
    }

    public List<GetTaskStepResponse> getTaskSteps()
    {
        return taskSteps;
    }

    public void setTaskSteps(List<GetTaskStepResponse> taskSteps)
    {
        this.taskSteps = taskSteps;
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
