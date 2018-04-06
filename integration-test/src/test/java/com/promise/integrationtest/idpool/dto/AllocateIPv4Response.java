package com.promise.integrationtest.idpool.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class AllocateIPv4Response
{

    @JsonProperty(value = "Address", required = true)
    private String address;
    @JsonProperty(value = "Pool", required = true)
    private GetIPv4PoolResponse pool;

    public String getAddress()
    {
        return address;
    }

    public void setAddress(String address)
    {
        this.address = address;
    }

    public GetIPv4PoolResponse getPool()
    {
        return pool;
    }

    public void setPool(GetIPv4PoolResponse pool)
    {
        this.pool = pool;
    }
}
