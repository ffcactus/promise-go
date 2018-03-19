package com.promise.integrationtest.util;

import java.util.List;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.promise.integrationtest.base.ResourceResponse;
import com.promise.integrationtest.dto.Message;

public class PromiseAssertUtil
{
    public static void isResource(ResourceResponse resp)
    {
        Assert.assertNotNull(resp.getId());
        Assert.assertNotNull(resp.getUri());
        Assert.assertTrue(resp.getUri().contains(resp.getId()));
    }

    public static <R> void assertPostMessage(String expectId, String url, R request)
    {
        final ResponseEntity<List<Message>> response = RestClient.post(
                url,
                request,
                new TypeReference<List<Message>>()
                {
                });
        Assert.assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        final List<Message> message = response.getBody();
        Assert.assertEquals(expectId, message.get(0).getId());
    }

}
