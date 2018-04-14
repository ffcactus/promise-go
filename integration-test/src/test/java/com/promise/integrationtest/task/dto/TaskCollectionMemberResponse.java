package com.promise.integrationtest.task.dto;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.MemberResponse;

@JsonIgnoreProperties(ignoreUnknown = false)
public class TaskCollectionMemberResponse extends MemberResponse
{
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "Percentage", required = true)
    private int percentage;
    @JsonProperty(value = "CurrentStep", required = true)
    private String currentStep;
    @JsonProperty(value = "TaskSteps", required = true)
    private List<GetTaskStepResponse> taskSteps;
    @JsonProperty(value = "ExecutionResult", required = true)
    private ExecutionResultResponse executionResult;

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
