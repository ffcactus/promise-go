package com.promise.integrationtest.idpool.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.MemberResponse;

public class IPv4PoolCollectionMemberResponse extends MemberResponse
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
