package com.promise.integrationtest.task;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.List;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.MessageEnum;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.task.dto.GetTaskResponse;
import com.promise.integrationtest.task.dto.PostTaskRequest;
import com.promise.integrationtest.task.dto.PostTaskStepRequest;
import com.promise.integrationtest.task.dto.TaskCollectionMemberResponse;
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

    /**
     * The task should support get collection with filter.
     * @throws UnsupportedEncodingException
     */
    @Test
    public void testGetCollection() throws UnsupportedEncodingException
    {
        final String name = "MyTask1";
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

        // Create a task.
        request.setName("MyTask2");
        final GetTaskResponse response2 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/task/",
                request,
                GetTaskResponse.class);

        // Test the collection is right.
        final List<TaskCollectionMemberResponse> members1 = PromiseAssertUtil
                .assertGetCollection(getRootURL() + "/promise/v1/task", 2, 2, TaskCollectionMemberResponse.class);
        Assert.assertTrue(members1.contains(response1));
        Assert.assertTrue(members1.contains(response2));
        
        
        // Test filter.
        final String filter1 = URLEncoder.encode("Name eq 'MyTask1'", "UTF-8");
        final List<TaskCollectionMemberResponse> members2 = PromiseAssertUtil
                .assertGetCollection(getRootURL() + "/promise/v1/task?$filter=" + filter1, 2, 1, TaskCollectionMemberResponse.class);
        Assert.assertTrue(members2.contains(response1));
        
        PromiseAssertUtil.assertUnknownFilter(getRootURL() + "/promise/v1/task", "xxx", "yyy");
        
        // Delete a task.
        PromiseAssertUtil.assertDeleteResource(getRootURL() + response1.getUri());
        
        // Test start and count.
        PromiseAssertUtil.assertGetColletcionWithStartCount(getRootURL() + "/promise/v1/task", 1);        
    }
    
    /**
     * You should specify task step.
     */
    @Test
    public void testPostTaskRequest()
    {
        final String name = "MyTask1";
        final String description = "MyTask description";

        final PostTaskRequest request = new PostTaskRequest();
        request.setName(name);
        request.setDescription(description);

        PromiseAssertUtil.assertPostMessage(
                getRootURL() + "/promise/v1/task/",
                MessageEnum.TaskNoStep.getId(),
                request);
    }
}
