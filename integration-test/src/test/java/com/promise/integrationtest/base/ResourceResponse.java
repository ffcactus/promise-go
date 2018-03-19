package com.promise.integrationtest.base;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * The common get resource response DTO.
 */
public class ResourceResponse
{
    @JsonProperty(value = "ID", required = true)
    private String id;
    @JsonProperty(value = "URI", required = true)
    private String uri;

    public String getId()
    {
        return id;
    }

    public void setId(String id)
    {
        this.id = id;
    }

    public String getUri()
    {
        return uri;
    }

    public void setUri(String uri)
    {
        this.uri = uri;
    }
}
