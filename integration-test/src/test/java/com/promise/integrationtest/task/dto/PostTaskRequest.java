package com.promise.integrationtest.task.dto;

import java.util.ArrayList;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class PostTaskRequest
{
    @JsonProperty(value = "MessageID", required = false)
    private String messageID;
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "CreatedByName", required = true)
    private String createdByName;
    @JsonProperty(value = "CreatedByURI", required = true)
    private String createdByURI;
    @JsonProperty(value = "TargetName", required = true)
    private String targetName;
    @JsonProperty(value = "TargetURI", required = true)
    private String targetURI;
    @JsonProperty(value = "TaskSteps", required = true)
    private List<PostTaskStepRequest> taskSteps;

    public PostTaskRequest()
    {
        taskSteps = new ArrayList<PostTaskStepRequest>();
    }

    public void addTaskStep(PostTaskStepRequest step)
    {
        taskSteps.add(step);
    }

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

    public List<PostTaskStepRequest> getTaskSteps()
    {
        return taskSteps;
    }

    public void setTaskSteps(List<PostTaskStepRequest> taskSteps)
    {
        this.taskSteps = taskSteps;
    }
}
