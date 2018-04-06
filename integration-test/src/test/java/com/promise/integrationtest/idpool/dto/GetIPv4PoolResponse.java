package com.promise.integrationtest.idpool.dto;

import java.util.ArrayList;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.ResourceResponse;

public class GetIPv4PoolResponse extends ResourceResponse
{
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "Ranges", required = true)
    private List<IPv4RangeResponse> ranges;
    @JsonProperty(value = "SubnetMask", required = true)
    private String subnetMask;
    @JsonProperty(value = "Gateway", required = true)
    private String gateway;
    @JsonProperty(value = "Domain", required = true)
    private String Domain;
    @JsonProperty(value = "DNSServers", required = true)
    private List<String> dnsServers;
    @JsonProperty(value = "Total", required = true)
    private int total;
    @JsonProperty(value = "Free", required = true)
    private int free;
    @JsonProperty(value = "Allocatable", required = true)
    private int allocatable;

    public GetIPv4PoolResponse()
    {
        ranges = new ArrayList<>();
        dnsServers = new ArrayList<>();
    }

    public String getName()
    {
        return name;
    }

    public void setName(String name)
    {
        this.name = name;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }

    public List<IPv4RangeResponse> getRanges()
    {
        return ranges;
    }

    public void setRanges(List<IPv4RangeResponse> ranges)
    {
        this.ranges = ranges;
    }

    public String getSubnetMask()
    {
        return subnetMask;
    }

    public void setSubnetMask(String subnetMask)
    {
        this.subnetMask = subnetMask;
    }

    public String getGateway()
    {
        return gateway;
    }

    public void setGateway(String gateway)
    {
        this.gateway = gateway;
    }

    public String getDomain()
    {
        return Domain;
    }

    public void setDomain(String domain)
    {
        Domain = domain;
    }

    public List<String> getDnsServers()
    {
        return dnsServers;
    }

    public void setDnsServers(List<String> dnsServers)
    {
        this.dnsServers = dnsServers;
    }

    public int getTotal()
    {
        return total;
    }

    public void setTotal(int total)
    {
        this.total = total;
    }

    public int getFree()
    {
        return free;
    }

    public void setFree(int free)
    {
        this.free = free;
    }

    public int getAllocatable()
    {
        return allocatable;
    }

    public void setAllocatable(int allocatable)
    {
        this.allocatable = allocatable;
    }
}
