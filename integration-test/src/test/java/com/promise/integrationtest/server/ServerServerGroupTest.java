package com.promise.integrationtest.server;

import java.io.UnsupportedEncodingException;

import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.util.RestClient;
import com.promise.integrationtest.util.ServerGroupAssertUtil;

public class ServerServerGroupTest extends PromiseIntegrationTest
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
        // Remove all the server.
        final ResponseEntity<DeleteResourceResponse> response1 = RestClient.delete(
                getRootURL() + "/promise/v1/server",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response1.getStatusCode());

        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response2 = RestClient.delete(
                getRootURL() + "/promise/v1/servergroup",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response2.getStatusCode());

        // Remove all the server-servergroup.
        final ResponseEntity<DeleteResourceResponse> response3 = RestClient.delete(
                getRootURL() + "/promise/v1/server-servergroup",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response3.getStatusCode());
    }

    @Test
    public void testHappyPath()
            throws UnsupportedEncodingException
    {
        // Add 3 servers.
        //        final String server1_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_1", "Username", "Password").getId();
        //        final String server2_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_2", "Username", "Password").getId();
        //        final String server3_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_3", "Username", "Password").getId();

        // Now the default server group should contain these 3 servers.
        final String serverGroup_ID = ServerGroupAssertUtil.assertGetServerGroupByName("all").getId();

    }
}
