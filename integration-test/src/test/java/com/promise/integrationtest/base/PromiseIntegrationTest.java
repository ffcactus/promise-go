package com.promise.integrationtest.base;

import com.promise.integrationtest.Environment;

public class PromiseIntegrationTest
{

    private static Environment environment = new Environment();

    public PromiseIntegrationTest()
    {
    }

    public Environment getEnvironment()
    {
        return environment;
    }

    public static String getRootURL()
    {
        return "http://" + Environment.getHostname();
    }
}
