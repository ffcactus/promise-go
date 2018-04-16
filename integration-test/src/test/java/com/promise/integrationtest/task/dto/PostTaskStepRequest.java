package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class PostTaskStepRequest
{
    @JsonProperty(value = "MessageID", required = false)
    private String messageID;
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "ExpectedExecutionMs", required = true)
    private long expectedExecutionMs;

    public PostTaskStepRequest()
    {

    }

    public PostTaskStepRequest(String name, long expectedExecutionMs)
    {
        this.name = name;
        this.expectedExecutionMs = expectedExecutionMs;
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

    public long getExpectedExecutionMs()
    {
        return expectedExecutionMs;
    }

    public void setExpectedExecutionMs(long expectedExecutionMs)
    {
        this.expectedExecutionMs = expectedExecutionMs;
    }
}
