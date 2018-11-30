package com.promise.common;

public enum PromiseCategory
{
    Auth ("Auth"),
    Task ("Task"),
    Vm ("Vm"),
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
