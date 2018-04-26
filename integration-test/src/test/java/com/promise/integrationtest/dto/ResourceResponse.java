package com.promise.integrationtest.dto;

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
    @JsonProperty(value = "Category", required = true)
    private String category;
    @JsonProperty(value = "CreatedAt", required = true)
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

    public String getCategory()
    {
        return category;
    }

    public void setCategory(String category)
    {
        this.category = category;
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

    @Override
    public boolean equals(Object obj)
    {
        if (obj == null)
            return false;
        if (obj == this)
            return true;
        if (!(obj instanceof ResourceResponse))
            return false;
        ResourceResponse other = (ResourceResponse) obj;
        if (!other.getId().equals(this.id))
            return false;
        if (!other.getUri().equals(this.uri))
            return false;
        return true;
    }
}
