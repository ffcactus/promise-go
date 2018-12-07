package com.promise.server.dto;

import com.promise.server.model.Server;

import lombok.Data;
import lombok.ToString;

@Data
@ToString
public class GetServerResponse
{
    public GetServerResponse()
    {

    }

    public GetServerResponse(Server server)
    {
        this.id = server.id;
        this.name = server.name;
        this.uri = "/rest/v1/vm/" + id;
        this.partition = server.partition;
        this.scope = server.scope;
    }

    public String uri;
    public String id;
    public String name;
    public String partition;
    public String scope;
}
