package com.promise.integrationtest.util;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.List;

import org.junit.Assert;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.promise.integrationtest.base.MessageEnum;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.dto.MemberResponse;
import com.promise.integrationtest.dto.Message;
import com.promise.integrationtest.dto.ResourceCollectionResponse;
import com.promise.integrationtest.dto.ResourceResponse;

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
     * Assert the action should be accepted.
     * 
     * @param url The URL for the POST.
     * @param request The request DTO.
     * @param responseClass The response DTO class.
     * @return The DTO if success.
     */
    public static <R, T> T assertActionResponse(String url, R request, Class<T> responseClass)
    {
        final ResponseEntity<T> response = RestClient.post(
                url,
                request,
                responseClass);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
        T ret = response.getBody();
        return ret;
    }

    /**
     * Assert the POST action should fail.
     * 
     * @param url The URL to post.
     * @param expectedMessageID The expected message's ID.
     * @param request The request DTO.
     */
    public static <R> void assertActionMessage(String url, String expectedMessageID, R request)
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
        Assert.assertEquals(MessageEnum.NotExist.getId(), message.get(0).getId());
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
    public static <T extends MemberResponse> List<T> assertGetCollection(
            String url,
            int expectedTotal,
            int execptedCount,
            Class<T> memberClass)
    {
        final ResponseEntity<ResourceCollectionResponse<T>> response = RestClient.get(
                url,
                new TypeReference<ResourceCollectionResponse<T>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, response.getStatusCode());
        final ResourceCollectionResponse<T> body = response.getBody();
        Assert.assertEquals(expectedTotal, body.getTotal());
        Assert.assertEquals(execptedCount, body.getCount());
        for (T each : body.getMember())
        {
            PromiseAssertUtil.isMemberResponse(each);
        }
        return body.getMember();
    }

    /**
     * Assert the common get collection operation.
     * When start, count, filter is omitted, the count and total should match.
     * If the total greater than 0, you can use start and skip to get the first
     * element and the last element.
     * If you can't get more then total elements.
     * 
     * @param url
     * @throws UnsupportedEncodingException
     */
    public static void assertGetColletcionWithStartCount(String url, int expectedTotal)
            throws UnsupportedEncodingException
    {
        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response1 = RestClient.get(
                url,
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        Assert.assertEquals(HttpStatus.OK, response1.getStatusCode());
        final ResourceCollectionResponse<MemberResponse> body1 = response1.getBody();
        Assert.assertEquals(expectedTotal, body1.getTotal());
        Assert.assertEquals(expectedTotal, body1.getCount());
        if (expectedTotal == 0)
        {
            return;
        }

        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response2 = RestClient.get(
                url + "?start=0&&count=1",
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        final ResourceCollectionResponse<MemberResponse> body2 = response2.getBody();
        Assert.assertEquals(expectedTotal, body2.getTotal());
        Assert.assertEquals(0, body2.getStart());
        Assert.assertEquals(1, body2.getCount());

        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response3 = RestClient.get(
                url + "?start=" + (expectedTotal - 1) + "&&count=1",
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        final ResourceCollectionResponse<MemberResponse> body3 = response3.getBody();
        Assert.assertEquals(expectedTotal, body3.getTotal());
        Assert.assertEquals(expectedTotal - 1, body3.getStart());
        Assert.assertEquals(1, body3.getCount());

        final ResponseEntity<ResourceCollectionResponse<MemberResponse>> response4 = RestClient.get(
                url + "?start=0&&count=" + (expectedTotal + 1),
                new TypeReference<ResourceCollectionResponse<MemberResponse>>()
                {
                });
        final ResourceCollectionResponse<MemberResponse> body4 = response4.getBody();
        Assert.assertEquals(expectedTotal, body4.getTotal());
        Assert.assertEquals(0, body4.getStart());
        Assert.assertEquals(expectedTotal, body4.getCount());
    }

    /**
     * Assert the filter name is unknown
     * 
     * @param url
     * @throws UnsupportedEncodingException
     */
    public static void assertUnknownFilter(String url, String name, String value)
            throws UnsupportedEncodingException
    {
        String filter1 = URLEncoder.encode(name + " eq '" + value + "'", "UTF-8");
        PromiseAssertUtil.assertGetMessage(url + "?$filter=" + filter1, MessageEnum.UnknownFilterName.getId());
    }

    public static void assertSameElements(List<?> firstList, List<?> secondList)
    {
        Assert.assertEquals(firstList.size(), secondList.size());

        // Iterate over the elements of the first list.
        for (int index = 0; index < firstList.size(); index++)
        {
            Assert.assertTrue(secondList.contains(firstList.get(index)));
        }
    }

}
