package com.promise.integrationtest.base;

public enum EventType
{
    Create ("Promise"),
    Update ("AA"),
    Delete ("Task"),
    DeleteCollection ("ServerServerGroup");

    private String id;

    EventType(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
