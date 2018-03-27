package com.promise.integrationtest.ws;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.promise.integrationtest.base.Category;
import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.EmptyResponse;
import com.promise.integrationtest.base.EventType;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.server.dto.GetServerGroupResponse;
import com.promise.integrationtest.server.dto.PostServerGroupRequest;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;
import com.promise.integrationtest.ws.dto.PostEventRequest;

public class WsTest extends PromiseIntegrationTest
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
        final PostServerGroupRequest request1 = new PostServerGroupRequest();
        request1.setName(name);
        request1.setDescription(description);

        // Create a server group.
        final GetServerGroupResponse response1 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/servergroup/",
                request1,
                GetServerGroupResponse.class);

        // Get it.
        final GetServerGroupResponse response2 = PromiseAssertUtil.assertGetResponse(
                getRootURL() + "/promise/v1/servergroup/" + response1.getId(),
                GetServerGroupResponse.class);
        Assert.assertEquals(name, response2.getName());
        Assert.assertEquals(description, response2.getDescription());

        PostEventRequest request2 = new PostEventRequest();
        request2.setCategory(Category.ServerGroup.getId());
        request2.setType(EventType.Create.getId());
        request2.setResourceId(response2.getId());
        final ObjectMapper mapper = new ObjectMapper();
        mapper.valueToTree(response2);
        //ObjectMapper.convertValue(response2, GetServerGroupResponse.class);
        request2.setData(mapper.valueToTree(response2));

        // Test create event.
        final ResponseEntity<EmptyResponse> response3 = RestClient.post(
                getRootURL() + "/promise/v1/ws-sender",
                request2,
                EmptyResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, response3.getStatusCode());
        
        // Test update event.
        request2.setType(EventType.Update.getId());
        final ResponseEntity<EmptyResponse> response4 = RestClient.post(
                getRootURL() + "/promise/v1/ws-sender",
                request2,
                EmptyResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, response4.getStatusCode());
        
        // Test delete event.
        request2.setData(null);
        request2.setType(EventType.Delete.getId());
        final ResponseEntity<EmptyResponse> response5 = RestClient.post(
                getRootURL() + "/promise/v1/ws-sender",
                request2,
                EmptyResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, response5.getStatusCode());
    }
    
}
