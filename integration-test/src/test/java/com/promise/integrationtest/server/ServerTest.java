package com.promise.integrationtest.server;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.server.dto.GetServerResponse;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;
import com.promise.integrationtest.util.ServerAssertUtil;

public class ServerTest extends PromiseIntegrationTest
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
        // Remove all the server.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                "/promise/v1/server",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }

    @Before
    public void setUp()
            throws Exception
    {
        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                "/promise/v1/server",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }

    @After
    public void tearDown()
            throws Exception
    {
    }
    
    /**
     * Happy path test add, get and delete operation.
     *
     */
    @Test
    public void testHappyPath()
    {
        final String uri = ServerAssertUtil.assertServerPosted("Mock_Hostname_1", "Username", "Password").getUri();
        PromiseAssertUtil.assertGetResponse(uri, GetServerResponse.class);
        PromiseAssertUtil.assertDeleteResource(uri);
    }
}
