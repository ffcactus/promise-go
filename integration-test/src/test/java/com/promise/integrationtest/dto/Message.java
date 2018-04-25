package com.promise.integrationtest.dto;

import java.util.Arrays;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Message
{
    @JsonProperty(value = "ID", required = true)
    private String id;
    @JsonProperty(value = "Severity", required = true)
    private String severity;
    @JsonProperty(value = "CreatedAt", required = true)
    private String createdAt;
    @JsonProperty(value = "Description", required = true)
    private String description;
    @JsonProperty(value = "Arguments", required = false)
    private List<Argument> argument;
    @JsonProperty(value = "Supports", required = false)
    private List<Support> support;

    public String getId()
    {
        return id;
    }

    public void setId(String id)
    {
        this.id = id;
    }

    public String getSeverity()
    {
        return severity;
    }

    public void setSeverity(String severity)
    {
        this.severity = severity;
    }

    public String getCreatedAt()
    {
        return createdAt;
    }

    public void setCreatedAt(String createAt)
    {
        this.createdAt = createAt;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }

    public List<Argument> getArgument()
    {
        return argument;
    }

    public void setArgument(List<Argument> argument)
    {
        this.argument = argument;
    }

    public List<Support> getSupport()
    {
        return support;
    }

    public void setSupport(List<Support> support)
    {
        this.support = support;
    }

    @Override
    public boolean equals(Object obj)
    {
        if (obj == null)
            return false;
        if (obj == this)
            return true;
        if (!(obj instanceof Message))
            return false;
        Message other = (Message) obj;
        if (!other.id.equals(this.id))
            return false;
        if (!other.severity.equals(this.severity))
            return false;
        if (!other.description.equals(this.description))
            return false;

        if (!Arrays.equals(this.argument.toArray(), other.argument.toArray()))
        {
            return false;
        }
        if (!Arrays.equals(this.support.toArray(), other.support.toArray()))
        {
            return false;
        }
        return true;
    }
}
