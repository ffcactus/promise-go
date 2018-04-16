package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.Message;

public class UpdateExecutionResultRequest
{
    @JsonProperty(value = "State", required = false)
    private String state;
    @JsonProperty(value = "Message", required = false)
    private Message message;

    public String getState()
    {
        return state;
    }

    public void setState(String state)
    {
        this.state = state;
    }

    public Message getMessage()
    {
        return message;
    }

    public void setMessage(Message message)
    {
        this.message = message;
    }
}
