package com.promise.integrationtest.server.message;

public enum ServerServerGroupMessage
{
    DeleteDefault ("Server.Message.ServerServerGroupDeleteDefault");
    private String id;

    ServerServerGroupMessage(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
