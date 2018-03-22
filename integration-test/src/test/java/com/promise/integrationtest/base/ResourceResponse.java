package com.promise.integrationtest.base;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * The common get resource response DTO.
 */
@JsonIgnoreProperties(ignoreUnknown = true)
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
