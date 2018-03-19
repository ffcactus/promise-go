package com.promise.integrationtest.base;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ResourceCollectionResponse<T>
{
    @JsonProperty(value = "Start", required = true)
    private int start;
    @JsonProperty(value = "Count", required = true)
    private int count;
    @JsonProperty(value = "Total", required = true)
    private int total;
    @JsonProperty(value = "Members", required = true)
    private List<T> member;

    public int getStart()
    {
        return start;
    }

    public void setStart(int start)
    {
        this.start = start;
    }

    public int getCount()
    {
        return count;
    }

    public void setCount(int count)
    {
        this.count = count;
    }

    public int getTotal()
    {
        return total;
    }

    public void setTotal(int total)
    {
        this.total = total;
    }

    public List<T> getMember()
    {
        return member;
    }

    public void setMember(List<T> member)
    {
        this.member = member;
    }
}
