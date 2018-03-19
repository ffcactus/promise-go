package com.promise.integrationtest.util;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.server.dto.PostServerRequest;
import com.promise.integrationtest.server.dto.PostServerResponse;

public class ServerAssertUtil
{
    /*
     * Assert the server has been posted.
     */
    public static PostServerResponse assertServerPosted(String hostname, String username, String password)
    {
        final PostServerRequest request = new PostServerRequest(hostname, username, password);
        // Create a server group.
        final ResponseEntity<PostServerResponse> response = RestClient.post(
                PromiseIntegrationTest.getRootURL() + "/promise/v1/server/",
                request,
                PostServerResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, response.getStatusCode());
        PromiseAssertUtil.isResource(response.getBody());
        return response.getBody();
    }
}
