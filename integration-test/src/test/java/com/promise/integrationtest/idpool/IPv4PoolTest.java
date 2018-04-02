package com.promise.integrationtest.idpool;

import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.idpool.dto.GetIPv4PoolResponse;
import com.promise.integrationtest.idpool.dto.IPv4Range;
import com.promise.integrationtest.idpool.dto.PostIPv4PoolRequest;
import com.promise.integrationtest.server.dto.GetServerGroupResponse;
import com.promise.integrationtest.util.IPv4PoolAssertUtil;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;

public class IPv4PoolTest extends PromiseIntegrationTest
{
    @BeforeClass
    public static void setUpBeforeClass()
            throws Exception
    {

    }

    @AfterClass
    public static void tearDownAfterClass()
            throws Exception
    {
    }

    @Before
    public void setUp()
            throws Exception
    {
        // Remove all IPv4 pool.
        final ResponseEntity<DeleteResourceResponse> response1 = RestClient.delete(
                getRootURL() + "/promise/v1/id-pool/ipv4",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response1.getStatusCode());
    }
    
    @Test
    public void testHappyPath()
            throws UnsupportedEncodingException
    {
        IPv4Range range1 = new IPv4Range("0.0.0.0", "0.0.0.255");
        IPv4Range range2 = new IPv4Range("0.0.1.0", "0.0.1.255");

        List<IPv4Range> ranges1 = new ArrayList<IPv4Range>();

        
        ranges1.add(range1);
        ranges1.add(range2);

        PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();

        
        request1.setName("pool1");
        request1.setDescription("description.");
        request1.setRanges(ranges1);
        request1.setSubnetMask("subnetMask");
        request1.setGateway("gateway");
        request1.setDomain("domain");
        String[] dns = {"dns1", "dns2"};
        request1.setDnsServers(Arrays.asList(dns));

        
        GetIPv4PoolResponse response1 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);        
        
        final GetServerGroupResponse response2 = PromiseAssertUtil.assertGetResponse(
                getRootURL() + response1.getUri(),
                GetServerGroupResponse.class);
        Assert.assertEquals("pool1", response2.getName());

        PromiseAssertUtil.assertDeleteResource(getRootURL() + response1.getUri());
    }
}
