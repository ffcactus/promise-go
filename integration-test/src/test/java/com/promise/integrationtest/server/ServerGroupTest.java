package com.promise.integrationtest.server;

import java.io.UnsupportedEncodingException;
import java.util.List;

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
import com.promise.integrationtest.common.object.message.PromiseMessage;
import com.promise.integrationtest.server.dto.GetServerGroupResponse;
import com.promise.integrationtest.server.dto.PostServerGroupRequest;
import com.promise.integrationtest.server.dto.ServerGroupMemberResponse;
import com.promise.integrationtest.server.message.ServerGroupMessage;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;
import com.promise.integrationtest.util.ServerGroupAssertUtil;

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
        final GetServerGroupResponse response1 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/servergroup/",
                request,
                GetServerGroupResponse.class);

        // Get it.
        final GetServerGroupResponse response2 = PromiseAssertUtil.assertGetResponse(
                getRootURL() + "/promise/v1/servergroup/" + response1.getId(),
                GetServerGroupResponse.class);
        Assert.assertEquals(name, response2.getName());
        Assert.assertEquals(description, response2.getDescription());

        // Delete it.
        PromiseAssertUtil.assertDeleteResource(getRootURL() + "/promise/v1/servergroup/" + response1.getId());
    }

    /**
     * The default servergroup should exist.
     * 
     * @throws UnsupportedEncodingException
     */
    @Test
    public void testDefaultServerGroupExist()
            throws UnsupportedEncodingException
    {
        ServerGroupAssertUtil.assertGetServerGroupByName("all");
    }

    /**
     * When you post a servergroup that exist, it will fail.
     */
    @Test
    public void testPostExist()
    {
        // Create the default "all" server group should fail.
        final PostServerGroupRequest request = new PostServerGroupRequest("all", "default server group");
        PromiseAssertUtil
                .assertPostMessage(getRootURL() + "/promise/v1/servergroup/", PromiseMessage.ResourceDuplicate.getId(), request);
    }

    /**
     * When you delete a servergroup that is not exist, it should fail.
     */
    @Test
    public void testDeleteNotExist()
    {
        PromiseAssertUtil.assertDeleteMessage(
                getRootURL() + "/promise/v1/servergroup/i_am_not_exist",
                PromiseMessage.ResourceNotExist.getId());
    }

    /**
     * When you delete the default server group, it will fail.
     * 
     * @throws UnsupportedEncodingException
     */
    @Test
    public void testDeleteDefault()
            throws UnsupportedEncodingException
    {
        GetServerGroupResponse response = ServerGroupAssertUtil.assertGetServerGroupByName("all");
        PromiseAssertUtil.assertDeleteMessage(getRootURL() + response.getUri(), ServerGroupMessage.DeleteDefault.getId());
    }

    /**
     * You can get collection.
     * @throws UnsupportedEncodingException 
     */
    @Test
    public void testGetCollection() throws UnsupportedEncodingException
    {
        GetServerGroupResponse r1 = ServerGroupAssertUtil.assertServerGroupPosted("group1", "description1");
        GetServerGroupResponse r2 = ServerGroupAssertUtil.assertServerGroupPosted("group2", "description2");
        GetServerGroupResponse r3 = ServerGroupAssertUtil.assertGetServerGroupByName("all");
        
        List<ServerGroupMemberResponse> members = PromiseAssertUtil.assertGetCollection(getRootURL() + "/promise/v1/servergroup", 3, ServerGroupMemberResponse.class);
        Assert.assertTrue(members.contains(r1));
        Assert.assertTrue(members.contains(r2));
        Assert.assertTrue(members.contains(r3));
        
        
    }
}
