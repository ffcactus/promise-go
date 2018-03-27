package com.promise.integrationtest.ws.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.JsonNode;

public class PostEventRequest
{
    @JsonProperty(value = "CreatedAt", required = true)
    private String createdAt;
    @JsonProperty(value = "Category", required = true)
    private String category;
    @JsonProperty(value = "Type", required = true)
    private String type;
    @JsonProperty(value = "ResourceID", required = true)
    private String resourceId;
    @JsonProperty(value = "Data")
    private JsonNode data;

    public String getCreatedAt()
    {
        return createdAt;
    }

    public void setCreatedAt(String createdAt)
    {
        this.createdAt = createdAt;
    }

    public String getCategory()
    {
        return category;
    }

    public void setCategory(String category)
    {
        this.category = category;
    }

    public String getType()
    {
        return type;
    }

    public void setType(String type)
    {
        this.type = type;
    }

    public String getResourceId()
    {
        return resourceId;
    }

    public void setResourceId(String resourceId)
    {
        this.resourceId = resourceId;
    }

    public JsonNode getData()
    {
        return data;
    }

    public void setData(JsonNode data)
    {
        this.data = data;
    }
}
