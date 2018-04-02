package com.promise.integrationtest.util;

import org.junit.Assert;

import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.idpool.dto.GetIPv4PoolResponse;
import com.promise.integrationtest.idpool.dto.PostIPv4PoolRequest;

public class IPv4PoolAssertUtil
{

    /*
     * Assert the IPv4 pool has been posted.
     */
    public static GetIPv4PoolResponse assertIPv4PoolPosted(PostIPv4PoolRequest request)
    {
        GetIPv4PoolResponse response = PromiseAssertUtil.assertPostResponse(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/id-pool/ipv4",
                request,
                GetIPv4PoolResponse.class);
        Assert.assertEquals(request.getName(), response.getName());
        Assert.assertEquals(request.getDescription(), response.getDescription());
        Assert.assertEquals(request.getSubnetMask(), response.getSubnetMask());
        Assert.assertEquals(request.getGateway(), response.getGateway());
        Assert.assertEquals(request.getDomain(), response.getDomain());
        PromiseAssertUtil.assertSameElements(request.getRanges(), response.getRanges());
        PromiseAssertUtil.assertSameElements(request.getDnsServers(), response.getDnsServers());
        return response;
    }
}
