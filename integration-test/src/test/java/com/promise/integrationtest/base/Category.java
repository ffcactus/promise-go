package com.promise.integrationtest.base;

public enum Category
{
    Promise ("Promise"),
    AA ("AA"),
    Task ("Task"),
    Server ("Server"),
    ServerGroup ("ServerGroup"),
    ServerServerGroup ("ServerServerGroup");

    private String id;

    Category(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
