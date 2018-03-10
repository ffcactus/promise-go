package com.promise.integrationtest.servergroup.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * The DTO of post server group.
 */
public class PostServerGroupRequest
{
    @JsonProperty("Name")
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
