package com.promise.integrationtest.util;

import java.util.Collection;

import org.junit.Assert;

import com.promise.common.PromiseResource;
import com.promise.common.response.PromiseOperationResponse;

public class CommonTestUtil
{

    public static <T> boolean collectionEquals(Collection<T> first, Collection<T> second)
    {
        if (first.size() != second.size())
        {
            return false;
        }
        return first.containsAll(second);
    }

    public static void assertPromiseResource(PromiseResource r)
    {
        Assert.assertNotNull(r.getId());
        Assert.assertNotNull(r.getUri());
        Assert.assertNotNull(r.getCategory());
        Assert.assertEquals("/rest/" + r.getCategory().getValue() + "/" + r.getId(), r.getUri());
    }

    public static void assertPromiseOperationResponse(PromiseOperationResponse r)
    {
        Assert.assertNotNull(r.getState());
        Assert.assertNotNull(r.getUri());
    }
}
