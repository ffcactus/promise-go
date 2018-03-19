package com.promise.integrationtest.server;

import org.junit.After;
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
import com.promise.integrationtest.server.message.ServerGroupMessage;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;

public class ServerGroupTest extends PromiseIntegrationTest
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
        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                getRootURL() + "/promise/v1/servergroup",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }

    @After
    public void tearDown()
            throws Exception
    {
    }

    @Test
    public void testHappyPath()
    {
        final String name = "MyServerGroup";
        final String description = "MyServerGroup description.";
        final PostServerGroupRequest request = new PostServerGroupRequest();
        request.setName(name);
        request.setDescription(description);

        // Create a server group.
        final ResponseEntity<GetServerGroupResponse> response1 = RestClient.post(
                getRootURL() + "/promise/v1/servergroup/",
                request,
                GetServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, response1.getStatusCode());
        PromiseAssertUtil.isResource(response1.getBody());

        // Get it.
        final ResponseEntity<GetServerGroupResponse> response2 = RestClient.get(
                getRootURL() + "/promise/v1/servergroup/" + response1.getBody().getId(),
                GetServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.OK, response2.getStatusCode());
        PromiseAssertUtil.isResource(response2.getBody());
        Assert.assertEquals(name, response2.getBody().getName());
        Assert.assertEquals(description, response2.getBody().getDescription());

        // Delete it.
        final ResponseEntity<DeleteResourceResponse> response3 = RestClient.delete(
                getRootURL() + "/promise/v1/servergroup/" + response1.getBody().getId(),
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response3.getStatusCode());
    }

    @Test
    public void testPostExist()
    {
        // Create the default "all" server group should fail.
        final PostServerGroupRequest request = new PostServerGroupRequest("all", "default server group");
        PromiseAssertUtil.assertPostMessage(ServerGroupMessage.EXIST.getId(), getRootURL() + "/promise/v1/servergroup/", request);
    }
}
