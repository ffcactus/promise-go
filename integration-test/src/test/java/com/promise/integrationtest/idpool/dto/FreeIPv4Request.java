package com.promise.integrationtest.idpool.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class FreeIPv4Request
{
    @JsonProperty(value="Address", required=true)
    private String address;

    public String getAddress()
    {
        return address;
    }

    public void setAddress(String address)
    {
        this.address = address;
    }
}
