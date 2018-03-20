package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class PostServerServerGroupRequest
{
    @JsonProperty(value = "ServerID", required = true)
    private String serverId;
    @JsonProperty(value = "ServerGroupID", required = true)
    private String serverGroupId;

    public PostServerServerGroupRequest()
    {

    }

    public PostServerServerGroupRequest(String serverId, String serverGroupId)
    {
        this.serverId = serverId;
        this.serverGroupId = serverGroupId;
    }

    public String getServerId()
    {
        return serverId;
    }

    public void setServerId(String serverId)
    {
        this.serverId = serverId;
    }

    public String getServerGroupId()
    {
        return serverGroupId;
    }

    public void setServerGroupId(String serverGroupId)
    {
        this.serverGroupId = serverGroupId;
    }
}
