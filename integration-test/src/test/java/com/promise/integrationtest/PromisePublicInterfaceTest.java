package com.promise.integrationtest;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.auth.sdk.dto.PostLoginRequest;
import com.promise.auth.sdk.dto.PostLoginResponse;
import com.promise.common.PromiseToken;
import com.promise.integrationtest.util.HttpJsonClient;

public class PromisePublicInterfaceTest extends PromiseTestBase
{
    protected PromiseToken token;

    public PromisePublicInterfaceTest()
            throws Exception
    {
        final PostLoginRequest request = new PostLoginRequest();
        request.setUserName(USERNAME);
        request.setPassword(PASSWORD);
        request.setDomain(DOMAIN);
        final ResponseEntity<PostLoginResponse> response = HttpJsonClient.post(
                HOSTNAME + "/rest/login",
                null,
                request,
                PostLoginResponse.class);
        if (response.getStatusCode() == HttpStatus.OK)
        {
            token = new PromiseToken(response.getBody().getToken());
        }
        else
        {
            throw new Exception("Failed to login.");
        }
    }
}
