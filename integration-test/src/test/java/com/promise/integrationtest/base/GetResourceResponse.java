package com.promise.integrationtest.base;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * The common get resource response DTO.
 */
public class GetResourceResponse
{
    @JsonProperty("ID")
    private String id;
    @JsonProperty("URI")
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
