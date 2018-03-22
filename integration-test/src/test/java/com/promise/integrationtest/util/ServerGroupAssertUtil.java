package com.promise.integrationtest.util;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.promise.integrationtest.base.MemberResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.base.ResourceCollectionResponse;
import com.promise.integrationtest.server.dto.GetServerGroupResponse;
import com.promise.integrationtest.server.dto.PostServerGroupRequest;

public class ServerGroupAssertUtil
{
    /*
     * Assert the servergroup has been posted.
     */
    public static GetServerGroupResponse assertServerGroupPosted(String name, String description)
    {
        
        final PostServerGroupRequest request = new PostServerGroupRequest(name, description);
        // Create a servergroup.
        GetServerGroupResponse response = PromiseAssertUtil.assertPostResponse(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/servergroup", 
                request, 
                GetServerGroupResponse.class);
        Assert.assertEquals(name, response.getName());
        Assert.assertEquals(description, response.getDescription());
        return response;
    }

    /*
     * Assert get servergroup by name exist, and only one exist.
     */
    public static GetServerGroupResponse assertGetServerGroupByName(String name)
            throws UnsupportedEncodingException
    {
        final String filter = URLEncoder.encode("Name eq '" + name + "'", "UTF-8");
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response1 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/servergroup?$filter=" + filter,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, response1.getStatusCode());
        Assert.assertEquals(1, response1.getBody().getMember().size());
        final ResponseEntity<GetServerGroupResponse> response2 = RestClient.get(
                PromiseIntegrationTest.getRootURL() + response1.getBody().getMember().get(0).getUri(),
                GetServerGroupResponse.class);
        Assert.assertEquals(HttpStatus.OK, response2.getStatusCode());
        PromiseAssertUtil.isResourceResponse(response2.getBody());
        Assert.assertEquals(name, response2.getBody().getName());
        return response2.getBody();

    }
}
