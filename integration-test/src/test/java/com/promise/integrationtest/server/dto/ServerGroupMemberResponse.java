package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.MemberResponse;

public class ServerGroupMemberResponse extends MemberResponse
{
    @JsonProperty(value = "Name", required = true)
    private String name;

    public String getName()
    {
        return name;
    }

    public void setName(String name)
    {
        this.name = name;
    }
}
