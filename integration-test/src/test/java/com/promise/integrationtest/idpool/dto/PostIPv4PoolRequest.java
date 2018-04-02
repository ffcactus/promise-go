package com.promise.integrationtest.idpool.dto;

import java.util.ArrayList;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public class PostIPv4PoolRequest
{
    @JsonProperty(value = "Name", required = true)
    private String name;
    @JsonProperty(value = "Description", required = false)
    private String description;
    @JsonProperty(value = "Ranges", required = true)
    private List<IPv4Range> ranges;
    @JsonProperty(value = "SubnetMask", required = true)
    private String subnetMask;
    @JsonProperty(value = "Gateway", required = true)
    private String gateway;
    @JsonProperty(value = "Domain", required = true)
    private String domain;
    @JsonProperty(value = "DNSServers", required = true)
    private List<String> dnsServers;

    public PostIPv4PoolRequest()
    {
        name = "";
        subnetMask = "";
        gateway = "";
        domain = "";
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

    public List<IPv4Range> getRanges()
    {
        return ranges;
    }

    public void setRanges(List<IPv4Range> ranges)
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
        return this.domain;
    }

    public void setDomain(String domain)
    {
        this.domain = domain;
    }

    public List<String> getDnsServers()
    {
        return dnsServers;
    }

    public void setDnsServers(List<String> dnsServers)
    {
        this.dnsServers = dnsServers;
    }
}
