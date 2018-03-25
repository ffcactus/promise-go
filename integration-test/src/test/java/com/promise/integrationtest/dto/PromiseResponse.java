package com.promise.integrationtest.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class PromiseResponse
{
    @JsonProperty(value = "ID", required = true)
    private String id;
    @JsonProperty(value = "URI", required = true)
    private String uri;
    @JsonProperty(value = "CreateAt", required = true)
    private String createdAt;
    @JsonProperty(value = "UpdatedAt", required = true)
    private String updatedAt;

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

    public String getCreatedAt()
    {
        return createdAt;
    }

    public void setCreatedAt(String createdAt)
    {
        this.createdAt = createdAt;
    }

    public String getUpdatedAt()
    {
        return updatedAt;
    }

    public void setUpdatedAt(String updatedAt)
    {
        this.updatedAt = updatedAt;
    }
}
