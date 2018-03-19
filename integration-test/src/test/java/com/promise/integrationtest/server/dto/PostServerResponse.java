package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.ResourceResponse;

/**
 * DTO of post server response.
 *
 */
public class PostServerResponse extends ResourceResponse
{
    @JsonProperty(value = "Hostname", required = true)
    private String hostname;
    @JsonProperty(value = "Type", required = true)
    private String type;
    @JsonProperty(value = "PhysicalUUID", required = true)
    private String physicalUUID;

    public String getHostname()
    {
        return hostname;
    }

    public void setHostname(String hostname)
    {
        this.hostname = hostname;
    }

    public String getType()
    {
        return type;
    }

    public void setType(String type)
    {
        this.type = type;
    }

    public String getPhysicalUUID()
    {
        return physicalUUID;
    }

    public void setPhysicalUUID(String physicalUUID)
    {
        this.physicalUUID = physicalUUID;
    }
}
