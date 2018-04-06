package com.promise.integrationtest.idpool.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.ResourceResponse;

public class IPv4RangeResponse
{
    @JsonProperty(value = "Start", required = true)
    private String start;
    @JsonProperty(value = "End", required = true)
    private String end;
    @JsonProperty(value = "Total", required = true)
    private String total;
    @JsonProperty(value = "Free", required = true)
    private String free;
    @JsonProperty(value = "Allocatable", required = true)
    private String allocatable;

    public IPv4RangeResponse()
    {

    }

    public IPv4RangeResponse(String start, String end)
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

    public String getTotal()
    {
        return total;
    }

    public void setTotal(String total)
    {
        this.total = total;
    }

    public String getFree()
    {
        return free;
    }

    public void setFree(String free)
    {
        this.free = free;
    }

    public String getAllocatable()
    {
        return allocatable;
    }

    public void setAllocatable(String allocatable)
    {
        this.allocatable = allocatable;
    }

    @Override
    public boolean equals(Object obj)
    {
        if (obj == null)
            return false;
        if (obj == this)
            return true;
        if (!(obj instanceof IPv4RangeResponse))
            return false;
        IPv4RangeResponse other = (IPv4RangeResponse) obj;
        if (!other.getStart().equals(this.start))
            return false;
        if (!other.getEnd().equals(this.end))
            return false;
        return true;
    }
}
