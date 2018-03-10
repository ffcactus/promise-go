package com.promise.integrationtest.auth;

import java.net.HttpURLConnection;
import java.util.ArrayList;
import java.util.List;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.auth.sdk.dto.CreateScopeRequest;
import com.promise.auth.sdk.dto.GetScopeListResponse;
import com.promise.auth.sdk.dto.GetScopeResponse;
import com.promise.auth.sdk.dto.PostLoginRequest;
import com.promise.auth.sdk.dto.PostLoginResponse;
import com.promise.common.PromiseAccessPoint;
import com.promise.common.PromiseToken;
import com.promise.common.response.PromiseGetResponse;
import com.promise.common.response.PromiseOperationResponse;
import com.promise.integrationtest.PromisePublicInterfaceTest;
import com.promise.integrationtest.util.CommonTestUtil;
import com.promise.integrationtest.util.HttpJsonClient;

public class ScopeTest extends PromisePublicInterfaceTest
{
    private static final String URI_HEAD = HOSTNAME + "/rest/scope";
    private static final CreateScopeRequest createRequest0;
    private static final CreateScopeRequest createRequest1;
    private static final List<PromiseAccessPoint> accessPointList0;
    private static final List<PromiseAccessPoint> accessPointList1;
    private static PromiseToken token;

    static
    {
        createRequest0 = new CreateScopeRequest();
        createRequest0.setName("Admin Scope");
        createRequest0.setDescription("Admin scope that has all the rights.");
        accessPointList0 = new ArrayList<>();
        accessPointList0.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/auth"));
        accessPointList0.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/task"));
        accessPointList0.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/scope"));
        createRequest0.setAccessPointList(accessPointList0);

        createRequest1 = new CreateScopeRequest();
        createRequest1.setName("User Scope");
        createRequest1.setDescription("User scope that doesn't have all the rights.");
        accessPointList1 = new ArrayList<>();
        accessPointList1.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/auth"));
        accessPointList1.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/task"));
        createRequest1.setAccessPointList(accessPointList1);
    }

    public ScopeTest()
            throws Exception
    {
        super();
    }

    @BeforeClass
    public static void setUpBeforeClass()
            throws Exception
    {
        final PostLoginRequest request = new PostLoginRequest();
        request.setUserName("Administrator");
        request.setPassword("admin");
        request.setDomain("local");
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

    @AfterClass
    public static void tearDownAfterClass()
            throws Exception
    {
    }

    @Before
    public void setUp()
            throws Exception
    {
    }

    @After
    public void tearDown()
            throws Exception
    {
    }

    @Test
    public void testCreateScope()
    {
        final ResponseEntity<PromiseOperationResponse> postRet = HttpJsonClient
                .post(URI_HEAD, token, createRequest0, PromiseOperationResponse.class);
        final PromiseOperationResponse postResponse = postRet.getBody();
        Assert.assertEquals(HttpURLConnection.HTTP_CREATED, postRet.getStatusCodeValue());
        CommonTestUtil.assertPromiseOperationResponse(postResponse);
        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteRet = HttpJsonClient
                .delete(HOSTNAME + postResponse.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet.getStatusCode());
    }

    @Test
    public void testDeleteExistScope()
    {
        final ResponseEntity<PromiseOperationResponse> postRet = HttpJsonClient
                .post(URI_HEAD, token, createRequest0, PromiseOperationResponse.class);
        final PromiseOperationResponse postResponse = postRet.getBody();
        Assert.assertEquals(HttpURLConnection.HTTP_CREATED, postRet.getStatusCodeValue());
        CommonTestUtil.assertPromiseOperationResponse(postResponse);

        final ResponseEntity<PromiseOperationResponse> deleteRet = HttpJsonClient
                .delete(HOSTNAME + postResponse.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet.getStatusCode());
    }

    @Test
    public void testDeleteNoneExistScope()
    {
        final ResponseEntity<PromiseOperationResponse> deleteRet = HttpJsonClient
                .delete(URI_HEAD + "/xxxx", token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.NOT_FOUND, deleteRet.getStatusCode());
    }

    @Test
    public void testGetScope()
    {
        final ResponseEntity<PromiseOperationResponse> postScopeRet = HttpJsonClient
                .post(URI_HEAD, token, createRequest0, PromiseOperationResponse.class);
        final PromiseOperationResponse postResponse = postScopeRet.getBody();

        Assert.assertEquals(HttpURLConnection.HTTP_CREATED, postScopeRet.getStatusCodeValue());
        CommonTestUtil.assertPromiseOperationResponse(postResponse);

        final ResponseEntity<PromiseGetResponse<GetScopeResponse>> getScopeRet = HttpJsonClient
                .getWithType(HOSTNAME + postResponse.getUri(), token, GetScopeResponse.class);

        final PromiseGetResponse<GetScopeResponse> getResponse = getScopeRet.getBody();
        final GetScopeResponse scope = getResponse.getData();
        Assert.assertEquals(HttpStatus.OK, getScopeRet.getStatusCode());
        CommonTestUtil.assertPromiseResource(scope);
        Assert.assertEquals(createRequest0.getName(), scope.getName());
        Assert.assertEquals(createRequest0.getDescription(), scope.getDescription());
        Assert.assertTrue(CommonTestUtil.collectionEquals(createRequest0.getAccessPointList(), scope.getAccessPointList()));

        Assert.assertEquals(postResponse.getUri(), getResponse.getUri());

        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteRet = HttpJsonClient
                .delete(HOSTNAME + getResponse.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet.getStatusCode());
    }

    @Test
    public void testGetScopeList()
    {
        final ResponseEntity<PromiseOperationResponse> postScopeRet0 = HttpJsonClient
                .post(URI_HEAD, token, createRequest0, PromiseOperationResponse.class);
        final PromiseOperationResponse postResponse0 = postScopeRet0.getBody();

        final ResponseEntity<PromiseOperationResponse> postScopeRet1 = HttpJsonClient
                .post(URI_HEAD, token, createRequest1, PromiseOperationResponse.class);
        final PromiseOperationResponse postResponse1 = postScopeRet1.getBody();

        Assert.assertEquals(HttpURLConnection.HTTP_CREATED, postScopeRet0.getStatusCodeValue());
        CommonTestUtil.assertPromiseOperationResponse(postResponse0);

        Assert.assertEquals(HttpURLConnection.HTTP_CREATED, postScopeRet1.getStatusCodeValue());
        CommonTestUtil.assertPromiseOperationResponse(postResponse1);

        final ResponseEntity<GetScopeListResponse> getScopeListRet = HttpJsonClient
                .get(URI_HEAD, token, GetScopeListResponse.class);
        final GetScopeListResponse getScopeListResponse = getScopeListRet.getBody();

        Assert.assertEquals(HttpStatus.OK, getScopeListRet.getStatusCode());
        Assert.assertEquals(0, getScopeListResponse.getStart());
        Assert.assertEquals(2, getScopeListResponse.getCount());
        Assert.assertEquals(2, getScopeListResponse.getMember().size());

        if (getScopeListResponse.getMember().get(0).getUri().equals(postResponse0.getUri()))
        {
            final GetScopeResponse t0 = getScopeListResponse.getMember().get(0);
            CommonTestUtil.assertPromiseResource(t0);
            Assert.assertEquals(createRequest0.getName(), t0.getName());
            Assert.assertEquals(createRequest0.getDescription(), t0.getDescription());
            CommonTestUtil.assertPromiseResource(t0);
            Assert.assertTrue(CommonTestUtil.collectionEquals(createRequest0.getAccessPointList(), t0.getAccessPointList()));

            final GetScopeResponse t1 = getScopeListResponse.getMember().get(1);
            CommonTestUtil.assertPromiseResource(t1);
            Assert.assertEquals(createRequest1.getName(), t1.getName());
            Assert.assertEquals(createRequest1.getDescription(), t1.getDescription());
            CommonTestUtil.assertPromiseResource(t1);
            Assert.assertTrue(CommonTestUtil.collectionEquals(createRequest1.getAccessPointList(), t1.getAccessPointList()));
        }
        else
        {
            final GetScopeResponse t0 = getScopeListResponse.getMember().get(0);
            CommonTestUtil.assertPromiseResource(t0);
            Assert.assertEquals(createRequest1.getName(), t0.getName());
            Assert.assertEquals(createRequest1.getDescription(), t0.getDescription());
            CommonTestUtil.assertPromiseResource(t0);
            Assert.assertTrue(CommonTestUtil.collectionEquals(createRequest1.getAccessPointList(), t0.getAccessPointList()));

            final GetScopeResponse t1 = getScopeListResponse.getMember().get(1);
            CommonTestUtil.assertPromiseResource(t1);
            Assert.assertEquals(createRequest0.getName(), t1.getName());
            Assert.assertEquals(createRequest0.getDescription(), t1.getDescription());
            CommonTestUtil.assertPromiseResource(t1);
            Assert.assertTrue(CommonTestUtil.collectionEquals(createRequest0.getAccessPointList(), t1.getAccessPointList()));
        }

        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteRet0 = HttpJsonClient
                .delete(HOSTNAME + postResponse0.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet0.getStatusCode());

        final ResponseEntity<PromiseOperationResponse> deleteRet1 = HttpJsonClient
                .delete(HOSTNAME + postResponse1.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet1.getStatusCode());
    }

}
