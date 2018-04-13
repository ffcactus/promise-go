package com.promise.integrationtest.task;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.DeleteResourceResponse;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.task.dto.GetTaskResponse;
import com.promise.integrationtest.task.dto.PostTaskRequest;
import com.promise.integrationtest.task.dto.PostTaskStepRequest;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;

public class TaskTest extends PromiseIntegrationTest
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
        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                getRootURL() + "/promise/v1/task",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }
    
    @Before
    public void setUp()
            throws Exception
    {
        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                getRootURL() + "/promise/v1/task",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }

    @After
    public void tearDown()
            throws Exception
    {
    }
    
    @Test
    public void testHappyPath()
    {
        final String name = "MyTask";
        final String description = "MyTask description";
        
        final PostTaskRequest request = new PostTaskRequest();
        request.setName(name);
        request.setDescription(description);
        request.addTaskStep(new PostTaskStepRequest("Step1", 1000));

        // Create a task.
        final GetTaskResponse response1 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/task/",
                request,
                GetTaskResponse.class);

        // Get it.
        final GetTaskResponse response2 = PromiseAssertUtil.assertGetResponse(
                getRootURL() + response1.getUri(),
                GetTaskResponse.class);
        Assert.assertEquals(name, response2.getName());
        Assert.assertEquals(description, response2.getDescription());

        // Delete it.
        PromiseAssertUtil.assertDeleteResource(getRootURL() + response1.getUri());
    }
}
