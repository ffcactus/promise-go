package com.promise.integrationtest.util;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.server.dto.GetServerResponse;
import com.promise.integrationtest.server.dto.PostServerRequest;

public class ServerAssertUtil
{
    /*
     * Assert the server has been posted.
     */
    public static GetServerResponse assertServerPosted(String hostname, String username, String password)
    {
        final PostServerRequest request = new PostServerRequest(hostname, username, password);
        // Create a server group.
        final ResponseEntity<GetServerResponse> response = RestClient.post(
                "/promise/v1/server/action/discover",
                request,
                GetServerResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
        PromiseAssertUtil.isResourceResponse(response.getBody());
        return response.getBody();
    }
}
