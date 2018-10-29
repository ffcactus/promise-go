package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.dto.ResourceResponse;

// todo : this class should not have this annotation.

/**
 * DTO of post server response.
 *
 */
@JsonIgnoreProperties(ignoreUnknown = true)
public class GetServerResponse extends ResourceResponse
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
