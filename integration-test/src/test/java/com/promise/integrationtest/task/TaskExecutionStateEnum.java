package com.promise.integrationtest.task;

public enum TaskExecutionStateEnum
{
    Ready("Ready"),
    Running("Running"),
    Suspended("Suspended"),
    Terminated("Terminated"),
    ;
    private String id;

    TaskExecutionStateEnum(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
