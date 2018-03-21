package com.promise.integrationtest.common.object.message;

public enum PromiseMessage
{
    InternalError ("Promise.Message.InternalError"),
    ResourceNotExist ("Promise.Message.ResourceNotExist"),
    ResourceDuplicate ("Promise.Message.ResourceDuplicate"),
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
