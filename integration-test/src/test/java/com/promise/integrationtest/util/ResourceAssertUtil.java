package com.promise.integrationtest.util;

import org.junit.Assert;

import com.promise.integrationtest.base.GetResourceResponse;

public class ResourceAssertUtil
{
    public static void isResource(GetResourceResponse resp)
    {
        Assert.assertNotNull(resp.getId());
        Assert.assertNotNull(resp.getUri());
        Assert.assertTrue(resp.getUri().contains(resp.getId()));
    }
}
