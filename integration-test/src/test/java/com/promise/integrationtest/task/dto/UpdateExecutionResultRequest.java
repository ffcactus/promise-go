package com.promise.integrationtest.task.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.ErrorResponse;

public class UpdateExecutionResultRequest
{
    @JsonProperty(value = "State", required = false)
    private String state;
    @JsonProperty(value = "ErrorResponse", required = false)
    private ErrorResponse errorResp;

    public String getState()
    {
        return state;
    }

    public void setState(String state)
    {
        this.state = state;
    }

    public ErrorResponse getErrorResponse()
    {
        return errorResp;
    }

    public void setErrorResponse(ErrorResponse errorResp)
    {
        this.errorResp = errorResp;
    }
}
