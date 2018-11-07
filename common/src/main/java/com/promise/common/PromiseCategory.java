package com.promise.common;

public enum PromiseCategory
{
    Auth ("Auth"),
    Task ("Task"),
    Global ("Global");

    private final String value;

    PromiseCategory(String value)
    {
        this.value = value;
    }

    String value()
    {
        return this.value;
    }
}
