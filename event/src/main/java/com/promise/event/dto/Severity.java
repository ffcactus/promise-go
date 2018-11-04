package com.promise.event.dto;

public enum Severity
{
    NORMAL ("Normal"),
    WARNING ("Warning"),
    CRITICAL ("Critical");

    private final String value;

    Severity(String value)
    {
        this.value = value;
    }

    String value()
    {
        return this.value;
    }
}
