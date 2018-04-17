package com.promise.integrationtest.task;

public enum TaskExecutionResultStateEnum
{
    Finished ("Finished"),
    Warning ("Warning"),
    Error ("Error"),
    Abort ("Abort"),
    Unknown ("Unknown"),
    ;
    private String id;

    TaskExecutionResultStateEnum(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
