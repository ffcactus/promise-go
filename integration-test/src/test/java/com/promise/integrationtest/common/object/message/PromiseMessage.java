package com.promise.integrationtest.common.object.message;

public enum PromiseMessage
{
    InternalError ("Promise.Message.InternalError"),
    NotExist ("Promise.Message.NotExist"),
    Duplicate ("Promise.Message.Duplicate"),
    InvalidRequest ("Promise.Message.InvalidRequest"),
    Timeout ("Promise.Message.Timeout");

    private String id;

    PromiseMessage(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
