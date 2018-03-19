package com.promise.integrationtest.server;

import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.server.dto.GetServerGroupResponse;
import com.promise.integrationtest.server.dto.PostServerGroupRequest;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;

public class ServerServerGroupTest extends PromiseIntegrationTest {

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
    {
    	
    }
}
