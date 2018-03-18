package com.promise.integrationtest.servergroup.message;

public enum ServerGroupMessage
{
    EXIST ("MessageIDGroupExist"),
    NOT_EXIST ("MessageIDGroupNotExist");

    private String id;

    ServerGroupMessage(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
