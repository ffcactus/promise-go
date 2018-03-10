package com.promise.integrationtest.auth;

import java.util.ArrayList;
import java.util.List;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Ignore;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.auth.sdk.dto.CreateScopeRequest;
import com.promise.auth.sdk.dto.CreateUserRequest;
import com.promise.auth.sdk.dto.GetUserListResponse;
import com.promise.auth.sdk.dto.GetUserResponse;
import com.promise.auth.sdk.dto.PostLoginRequest;
import com.promise.auth.sdk.dto.PostLoginResponse;
import com.promise.common.PromiseAccessPoint;
import com.promise.common.PromiseToken;
import com.promise.common.response.PromiseGetResponse;
import com.promise.common.response.PromiseOperationResponse;
import com.promise.common.response.PromiseResponseState;
import com.promise.integrationtest.PromisePublicInterfaceTest;
import com.promise.integrationtest.util.CommonTestUtil;
import com.promise.integrationtest.util.HttpJsonClient;

public class UserTest extends PromisePublicInterfaceTest
{
    private static final CreateScopeRequest createScopeRequest;
    private static String scopeUri;
    private static List<String> scopeUriList;
    private static PromiseToken token;

    static
    {
        createScopeRequest = new CreateScopeRequest();
        createScopeRequest.setName("Admin Scope");
        createScopeRequest.setDescription("Admin scope that has all the rights.");
        final List<PromiseAccessPoint> accessPointList = new ArrayList<>();
        accessPointList.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/auth"));
        accessPointList.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/task"));
        accessPointList.add(new PromiseAccessPoint(PromiseAccessPoint.URI, "rest/scope"));
        createScopeRequest.setAccessPointList(accessPointList);
    }

    public UserTest()
            throws Exception
    {
        super();
        // TODO Auto-generated constructor stub
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

        // Create scopes.
        final ResponseEntity<PromiseOperationResponse> createScopeResponse = HttpJsonClient
                .post(HOSTNAME + "/rest/scope", token, createScopeRequest, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, createScopeResponse.getStatusCode());
        scopeUri = createScopeResponse.getBody().getUri();
        scopeUriList = new ArrayList<>();
        scopeUriList.add(scopeUri);
    }

    @AfterClass
    public static void tearDownAfterClass()
            throws Exception
    {
        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteRet = HttpJsonClient
                .delete(HOSTNAME + scopeUri, token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteRet.getStatusCode());
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
    public void testCreateUser()
    {
        final CreateUserRequest createUserRequest = new CreateUserRequest();
        final List<String> scopeUriList = new ArrayList<>();
        scopeUriList.add(scopeUri);
        createUserRequest.setUsername("baibin");
        createUserRequest.setEmail("baibin@email.com");
        createUserRequest.setPassword("iforgot".toCharArray());
        createUserRequest.setScopeUriList(scopeUriList);

        final ResponseEntity<PromiseOperationResponse> responseEntity = HttpJsonClient
                .post(HOSTNAME + "/rest/user", token, createUserRequest, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.CREATED, responseEntity.getStatusCode());
        final PromiseOperationResponse postResponse = responseEntity.getBody();
        CommonTestUtil.assertPromiseOperationResponse(postResponse);
        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteResponseEntity = HttpJsonClient
                .delete(HOSTNAME + postResponse.getUri(), token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteResponseEntity.getStatusCode());
    }

    @Test
    public void testGetUser()
    {
        // Create the user to get later.
        final CreateUserRequest createUserRequest = new CreateUserRequest();
        final List<String> scopeUriList = new ArrayList<>();
        scopeUriList.add(scopeUri);
        createUserRequest.setUsername("baibin");
        createUserRequest.setEmail("baibin@email.com");
        createUserRequest.setPassword("iforgot".toCharArray());
        createUserRequest.setScopeUriList(scopeUriList);
        final ResponseEntity<PromiseOperationResponse> createUserResponseEntity = HttpJsonClient
                .post(HOSTNAME + "/rest/user", token, createUserRequest, PromiseOperationResponse.class);
        final String userUri = createUserResponseEntity.getBody().getUri();

        // Get the user that is created before.
        final ResponseEntity<PromiseGetResponse<GetUserResponse>> getUserResponseEntity = HttpJsonClient
                .getWithType(HOSTNAME + userUri, token, GetUserResponse.class);

        final PromiseGetResponse<GetUserResponse> getUserResponse = getUserResponseEntity.getBody();
        final GetUserResponse user = getUserResponse.getData();

        Assert.assertEquals(HttpStatus.OK, getUserResponseEntity.getStatusCode());
        CommonTestUtil.assertPromiseResource(user);
        Assert.assertEquals(createUserRequest.getUsername(), user.getUsername());
        Assert.assertEquals(createUserRequest.getEmail(), user.getEmail());
        Assert.assertTrue(CommonTestUtil.collectionEquals(scopeUriList, user.getScopeUri()));

        Assert.assertEquals(userUri, getUserResponse.getUri());

        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteResponseEntity = HttpJsonClient
                .delete(HOSTNAME + userUri, token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteResponseEntity.getStatusCode());
    }

    @Test
    public void testGetNonexistUser()
    {
        final ResponseEntity<PromiseGetResponse<GetUserResponse>> getUserResponseEntity = HttpJsonClient
                .getWithType(HOSTNAME + "/rest/user/xxxx", token, GetUserResponse.class);
        Assert.assertEquals(HttpStatus.NOT_FOUND, getUserResponseEntity.getStatusCode());

        Assert.assertEquals(PromiseResponseState.WARN, getUserResponseEntity.getBody().getState());
        Assert.assertNotNull(getUserResponseEntity.getBody().getReason());
        Assert.assertNotNull(getUserResponseEntity.getBody().getSolution());
        Assert.assertNull(getUserResponseEntity.getBody().getData());
    }

    @Test
    public void testDeleteUser()
    {
        // Create the user to delete later.
        final CreateUserRequest createUserRequest = new CreateUserRequest();
        final List<String> scopeUriList = new ArrayList<>();
        scopeUriList.add(scopeUri);
        createUserRequest.setUsername("baibin");
        createUserRequest.setEmail("baibin@email.com");
        createUserRequest.setPassword("iforgot".toCharArray());
        createUserRequest.setScopeUriList(scopeUriList);
        final ResponseEntity<PromiseOperationResponse> createUserResponseEntity = HttpJsonClient
                .post(HOSTNAME + "/rest/user", token, createUserRequest, PromiseOperationResponse.class);
        final String userUri = createUserResponseEntity.getBody().getUri();

        // Delete the user that is created before.
        final ResponseEntity<PromiseOperationResponse> deleteResponseEntity = HttpJsonClient
                .delete(HOSTNAME + userUri, token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteResponseEntity.getStatusCode());
    }

    @Test
    public void testDeleteNonexistUser()
    {
        final ResponseEntity<PromiseOperationResponse> deleteResponseEntity = HttpJsonClient
                .delete(HOSTNAME + "/rest/user/xxxx", token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.NOT_FOUND, deleteResponseEntity.getStatusCode());
    }

    @Test
    public void testGetUserList()
    {
        // Create the user to get later.
        final CreateUserRequest createUserRequest = new CreateUserRequest();
        final List<String> scopeUriList = new ArrayList<>();
        scopeUriList.add(scopeUri);
        createUserRequest.setUsername("baibin");
        createUserRequest.setEmail("baibin@email.com");
        createUserRequest.setPassword("iforgot".toCharArray());
        createUserRequest.setScopeUriList(scopeUriList);
        final ResponseEntity<PromiseOperationResponse> createUserResponseEntity = HttpJsonClient
                .post(HOSTNAME + "/rest/user", token, createUserRequest, PromiseOperationResponse.class);
        final String userUri = createUserResponseEntity.getBody().getUri();

        // Get the user list.
        final ResponseEntity<GetUserListResponse> getUserResponseEntity = HttpJsonClient
                .get(HOSTNAME + "/rest/user", token, GetUserListResponse.class);
        Assert.assertEquals(HttpStatus.OK, getUserResponseEntity.getStatusCode());
        final GetUserListResponse getUserListResponse = getUserResponseEntity.getBody();

        // Check count.
        Assert.assertEquals(2, getUserListResponse.getCount());

        // Make sure both created user and Administrator getted.
        boolean foundCreatedUser = false;
        boolean foundAdministrator = false;
        for (final GetUserResponse each : getUserListResponse.getMember())
        {
            CommonTestUtil.assertPromiseResource(each);
            if (each.getUsername().equals("Administrator"))
            {
                foundAdministrator = true;
            }
            if (each.getUsername().equals(createUserRequest.getUsername()))
            {
                Assert.assertEquals(createUserRequest.getUsername(), each.getUsername());
                Assert.assertEquals(createUserRequest.getEmail(), each.getEmail());
                Assert.assertTrue(CommonTestUtil.collectionEquals(scopeUriList, each.getScopeUri()));
                foundCreatedUser = true;
            }
        }
        Assert.assertTrue(foundCreatedUser);
        Assert.assertTrue(foundAdministrator);

        // Clean up.
        final ResponseEntity<PromiseOperationResponse> deleteResponseEntity = HttpJsonClient
                .delete(HOSTNAME + userUri, token, PromiseOperationResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, deleteResponseEntity.getStatusCode());
    }

    @Ignore
    void testCreateUserExist()
    {

    }
}
