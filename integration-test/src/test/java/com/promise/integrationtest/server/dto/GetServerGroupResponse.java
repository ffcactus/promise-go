package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.ResourceResponse;

/**
 * The DTO of get server group response.
 */
public class GetServerGroupResponse extends ResourceResponse
{
    @JsonProperty(value = "Name", required = true)
    private String name;

    @JsonProperty("Description")
    private String description;

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
}
