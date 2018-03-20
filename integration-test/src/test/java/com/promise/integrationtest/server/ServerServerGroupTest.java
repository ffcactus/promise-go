package com.promise.integrationtest.server;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;

import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.MemberResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.base.ResourceCollectionResponse;
import com.promise.integrationtest.server.dto.PostServerServerGroupRequest;
import com.promise.integrationtest.server.dto.PostServerServerGroupResponse;
import com.promise.integrationtest.util.RestClient;
import com.promise.integrationtest.util.ServerAssertUtil;
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
        final String server1_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_1", "Username", "Password").getId();
        final String server2_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_2", "Username", "Password").getId();
        final String server3_ID = ServerAssertUtil.assertServerPosted("Mock_Hostname_3", "Username", "Password").getId();

        // Now the default server group should contain these 3 servers.
        final String defaultServerGroupID = ServerGroupAssertUtil.assertGetServerGroupByName("all").getId();

        final String filter1 = URLEncoder.encode("ServerGroupID eq '" + defaultServerGroupID + "'", "UTF-8");
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response1 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup?$filter=" + filter1,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, response1.getStatusCode());
        Assert.assertEquals(3, response1.getBody().getMember().size());

        // Create 2 servergroup.
        final String group1_ID = ServerGroupAssertUtil.assertServerGroupPosted("Group1", "Description1").getId();
        final String group2_ID = ServerGroupAssertUtil.assertServerGroupPosted("Group2", "Description2").getId();

        // Put server1 in group1
        final PostServerServerGroupRequest ssgRequest1 = new PostServerServerGroupRequest(server1_ID, group1_ID);
        final ResponseEntity<PostServerServerGroupResponse> ssgResponse1 = RestClient.post(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup",
                ssgRequest1,
                PostServerServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, ssgResponse1.getStatusCode());

        // Put server2 and server 3 in group2
        final PostServerServerGroupRequest ssgRequest2 = new PostServerServerGroupRequest(server2_ID, group2_ID);
        final ResponseEntity<PostServerServerGroupResponse> ssgResponse2 = RestClient.post(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup",
                ssgRequest2,
                PostServerServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, ssgResponse2.getStatusCode());

        final PostServerServerGroupRequest ssgRequest3 = new PostServerServerGroupRequest(server3_ID, group2_ID);
        final ResponseEntity<PostServerServerGroupResponse> ssgResponse3 = RestClient.post(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup",
                ssgRequest3,
                PostServerServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, ssgResponse3.getStatusCode());

        // group 1 should have 1 server.
        final String filter2 = URLEncoder.encode("ServerGroupID eq '" + group1_ID + "'", "UTF-8");
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> ssgResponse4 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup?$filter=" + filter2,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, ssgResponse4.getStatusCode());
        Assert.assertEquals(1, ssgResponse4.getBody().getMember().size());

        // group 2 should have 2 server.
        final String filter3 = URLEncoder.encode("ServerGroupID eq '" + group2_ID + "'", "UTF-8");
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> ssgResponse5 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup?$filter=" + filter3,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, ssgResponse5.getStatusCode());
        Assert.assertEquals(2, ssgResponse5.getBody().getMember().size());

        // server 1 should belong to 2 groups.
        final String filter4 = URLEncoder.encode("ServerID eq '" + server1_ID + "'", "UTF-8");
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> ssgResponse6 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server-servergroup?$filter=" + filter4,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, ssgResponse6.getStatusCode());
        Assert.assertEquals(2, ssgResponse5.getBody().getMember().size());
    }
}
