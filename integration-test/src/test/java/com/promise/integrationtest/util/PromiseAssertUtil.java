package com.promise.integrationtest.util;

import java.util.List;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.MemberResponse;
import com.promise.integrationtest.base.ResourceCollectionResponse;
import com.promise.integrationtest.base.ResourceResponse;
import com.promise.integrationtest.common.object.message.PromiseMessage;
import com.promise.integrationtest.dto.Message;

public class PromiseAssertUtil
{
    /**
     * Assert the resource response is a standard Promise resource.
     * 
     * @param resp The response DTO.
     */
    public static void isResourceResponse(ResourceResponse resp)
    {
        Assert.assertNotNull(resp.getId());
        Assert.assertNotNull(resp.getUri());
        Assert.assertTrue(resp.getUri().contains(resp.getId()));
    }
    
    /**
     * Assert the resource response is a standard Promise resource.
     * 
     * @param resp The response DTO.
     */
    public static void isMemberResponse(MemberResponse resp)
    {
        Assert.assertNotNull(resp.getId());
        Assert.assertNotNull(resp.getUri());
        Assert.assertTrue(resp.getUri().contains(resp.getId()));
    }    

    /**
     * Assert the resource should be posted.
     * 
     * @param url The URL for the POST.
     * @param request The request DTO.
     * @param responseClass The response DTO class.
     * @return The DTO if success.
     */
    public static <R, T extends ResourceResponse> T assertPostResponse(String url, R request, Class<T> responseClass)
    {
        final ResponseEntity<T> response = RestClient.post(
                url,
                request,
                responseClass);
        Assert.assertEquals(HttpStatus.CREATED, response.getStatusCode());
        T ret = response.getBody();
        PromiseAssertUtil.isResourceResponse(ret);
        return ret;
    }

    /**
     * Assert the POST should fail.
     * 
     * @param url The URL to post.
     * @param expectedMessageID The expected message's ID.
     * @param request The request DTO.
     */
    public static <R> void assertPostMessage(String url, String expectedMessageID, R request)
    {
        final ResponseEntity<List<Message>> response = RestClient.post(
                url,
                request,
                new TypeReference<List<Message>>()
                {
                });
        Assert.assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        final List<Message> message = response.getBody();
        Assert.assertEquals(expectedMessageID, message.get(0).getId());
    }

    /**
     * Assert the resource could be get.
     * 
     * @param url The URL for the GET.
     * @param responseClass The response DTO.
     * @return The response DTO.
     */
    public static <T extends ResourceResponse> T assertGetResponse(String url, Class<T> responseClass)
    {
        final ResponseEntity<T> response = RestClient.get(
                url,
                responseClass);
        Assert.assertEquals(HttpStatus.OK, response.getStatusCode());
        T body = response.getBody();
        PromiseAssertUtil.isResourceResponse(body);
        return body;
    }

    /**
     * Assert the GET should fail, and the message should match.
     * 
     * @param url The URL to GET.
     * @param expectedMessageID The expected message ID.
     */
    public static void assertGetMessage(String url, String expectedMessageID)
    {
        final ResponseEntity<List<Message>> response = RestClient.get(
                url,
                new TypeReference<List<Message>>()
                {
                });
        Assert.assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        final List<Message> message = response.getBody();
        Assert.assertEquals(expectedMessageID, message.get(0).getId());
    }

    /**
     * Assert the DELETE should success, and the resource can't be GET anymore.
     * 
     * @param url The URL to DELETE.
     */
    public static void assertDeleteResource(String url)
    {
        final ResponseEntity<DeleteResourceResponse> response1 = RestClient.delete(url, DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response1.getStatusCode());

        final ResponseEntity<List<Message>> response2 = RestClient.get(
                url,
                new TypeReference<List<Message>>()
                {
                });
        Assert.assertEquals(HttpStatus.BAD_REQUEST, response2.getStatusCode());
        final List<Message> message = response2.getBody();
        Assert.assertEquals(PromiseMessage.ResourceNotExist.getId(), message.get(0).getId());
    }

    /**
     * Assert the DELETE should fail, and the error message should match.
     * 
     * @param url The URL to DELETE.
     */
    public static void assertDeleteMessage(String url, String expectedMessageID)
    {
        final ResponseEntity<List<Message>> response = RestClient.delete(url, new TypeReference<List<Message>>()
        {
        });

        Assert.assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        final List<Message> message = response.getBody();
        Assert.assertEquals(expectedMessageID, message.get(0).getId());
    }

    /**
     * Assert the GET collection should success.
     * Assert total member match.
     * Assert each member is MemberResponse.
     * 
     * @param url The URL to GET.
     * @param expectedTotal The expected total members.
     * @param memberClass The DTO of the member.
     */
    public static <T extends MemberResponse> List<T> assertGetCollection(String url, int expectedTotal, Class<T> memberClass)
    {
        final ResponseEntity<ResourceCollectionResponse<T>> response = RestClient.get(
                url,
                new TypeReference<ResourceCollectionResponse<T>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, response.getStatusCode());
        final ResourceCollectionResponse<T> body = response.getBody();
        Assert.assertEquals(expectedTotal, body.getTotal());
        for (T each : body.getMember()) {
            PromiseAssertUtil.isMemberResponse(each);
        }
        return body.getMember();
    }
}
