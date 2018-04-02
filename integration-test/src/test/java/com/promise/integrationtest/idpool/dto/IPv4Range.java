package com.promise.integrationtest.idpool.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.ResourceResponse;

public class IPv4Range
{
    @JsonProperty(value = "Start", required = true)
    private String start;
    @JsonProperty(value = "End", required = true)
    private String end;

    public IPv4Range()
    {

    }

    public IPv4Range(String start, String end)
    {
        this.start = start;
        this.end = end;
    }

    public String getStart()
    {
        return start;
    }

    public void setStart(String start)
    {
        this.start = start;
    }

    public String getEnd()
    {
        return end;
    }

    public void setEnd(String end)
    {
        this.end = end;
    }
    
    @Override
    public boolean equals(Object obj)
    {
        if (obj == null)
            return false;
        if (obj == this)
            return true;
        if (!(obj instanceof IPv4Range))
            return false;
        IPv4Range other = (IPv4Range) obj;
        if (!other.getStart().equals(this.start))
            return false;
        if (!other.getEnd().equals(this.end))
            return false;
        return true;
    }
}
