package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.promise.integrationtest.base.ResourceResponse;

public class PostServerServerGroupResponse extends ResourceResponse
{
    @JsonProperty(value = "ServerID", required = true)
    private String serverId;
    @JsonProperty(value = "ServerGroupID", required = true)
    private String serverGroupId;
    @JsonProperty(value = "ServerURI", required = true)
    private String serverUri;
    @JsonProperty(value = "ServerGroupURI", required = true)
    private String serverGroupUri;

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

    public String getServerUri()
    {
        return serverUri;
    }

    public void setServerUri(String serverUri)
    {
        this.serverUri = serverUri;
    }

    public String getServerGroupUri()
    {
        return serverGroupUri;
    }

    public void setServerGroupUri(String serverGroupUri)
    {
        this.serverGroupUri = serverGroupUri;
    }

}
