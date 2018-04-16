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

import com.promise.integrationtest.base.Category;
import com.promise.integrationtest.base.MessageEnum;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.task.dto.GetTaskResponse;
import com.promise.integrationtest.task.dto.GetTaskStepResponse;
import com.promise.integrationtest.task.dto.PostTaskRequest;
import com.promise.integrationtest.task.dto.PostTaskStepRequest;
import com.promise.integrationtest.task.dto.TaskCollectionMemberResponse;
import com.promise.integrationtest.task.dto.UpdateExecutionResultRequest;
import com.promise.integrationtest.task.dto.UpdateTaskStepRequest;
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
     * The task you posted should have all the property you specified and the
     * default values.
     */
    @Test
    public void testPostTaskWithFullProperty()
    {
        final String messageID = "messageID";
        final String name = "MyTask";
        final String description = "MyTask description";
        final String createdByName = "createdByName";
        final String createdByUrl = "createdByUrl";
        final String targetName = "targetName";
        final String targetUrl = "targetUrl";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.setMessageID(messageID);
        request1.setCreatedByName(createdByName);
        request1.setCreatedByURI(createdByUrl);
        request1.setTargetName(targetName);
        request1.setTargetURI(targetUrl);

        PostTaskStepRequest postTaskStepReqeust1 = new PostTaskStepRequest("Step1", 1000);
        postTaskStepReqeust1.setMessageID(messageID);
        postTaskStepReqeust1.setDescription(description);
        PostTaskStepRequest postTaskStepReqeust2 = new PostTaskStepRequest("Step2", 2000);
        postTaskStepReqeust2.setMessageID(messageID);
        postTaskStepReqeust2.setDescription(description);

        request1.addTaskStep(postTaskStepReqeust1);
        request1.addTaskStep(postTaskStepReqeust2);

        // Create a task.
        final GetTaskResponse response1 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/task/",
                request1,
                GetTaskResponse.class);
        Assert.assertEquals(messageID, response1.getMessageID());
        Assert.assertEquals(name, response1.getName());
        Assert.assertEquals(description, response1.getDescription());
        Assert.assertEquals(Category.Task.getId(), response1.getCategory());
        Assert.assertEquals(TaskExecutionStateEnum.Ready.getId(), response1.getExecutionState());
        Assert.assertEquals(TaskExecutionResultStateEnum.Unknown.getId(), response1.getExecutionResult().getState());

        Assert.assertEquals(createdByName, response1.getCreatedByName());
        Assert.assertEquals(createdByUrl, response1.getCreatedByURI());
        Assert.assertEquals(targetName, response1.getTargetName());
        Assert.assertEquals(targetUrl, response1.getTargetURI());
        Assert.assertEquals(3000, response1.getExpectedExecutionMs());
        Assert.assertEquals(0, response1.getPercentage());
        Assert.assertEquals(2, response1.getTaskSteps().size());

        GetTaskStepResponse step1 = response1.getTaskSteps().get(0);
        GetTaskStepResponse step2 = response1.getTaskSteps().get(1);

        Assert.assertEquals(messageID, step1.getMessageID());
        Assert.assertEquals("Step1", step1.getName());
        Assert.assertEquals(description, step1.getDescription());
        Assert.assertEquals(1000, step1.getExpectedExecutionMs());
        Assert.assertEquals(TaskExecutionStateEnum.Ready.getId(), step1.getExecutionState());
        Assert.assertEquals(TaskExecutionResultStateEnum.Unknown.getId(), step1.getExecutionResult().getState());

        Assert.assertEquals(messageID, step2.getMessageID());
        Assert.assertEquals("Step2", step2.getName());
        Assert.assertEquals(description, step2.getDescription());
        Assert.assertEquals(2000, step2.getExpectedExecutionMs());
        Assert.assertEquals(TaskExecutionStateEnum.Ready.getId(), step2.getExecutionState());
        Assert.assertEquals(TaskExecutionResultStateEnum.Unknown.getId(), step2.getExecutionResult().getState());
    }

    /**
     * The task should support get collection with filter.
     * 
     * @throws UnsupportedEncodingException
     */
    @Test
    public void testGetCollection()
            throws UnsupportedEncodingException
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
                .assertGetCollection(
                        getRootURL() + "/promise/v1/task?$filter=" + filter1,
                        2,
                        1,
                        TaskCollectionMemberResponse.class);
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
    
    /**
     * You should provide the right update task step request.
     */
    @Test
    public void testUpdateTaskStepRequest()
    {
        final String name = "MyTask";
        final String description = "MyTask description";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.addTaskStep(new PostTaskStepRequest("Step1", 1000));

        // Create a task.
        final GetTaskResponse response1 = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/task/",
                request1,
                GetTaskResponse.class);

        // name should be valid.
        UpdateTaskStepRequest request2 = new UpdateTaskStepRequest();
        request2.setName("xxxx");
        PromiseAssertUtil.assertPostMessage(
                getRootURL() + response1.getUri() + "/action/updateTaskStep",
                MessageEnum.UnknownPropertyValue.getId(),
                request2);

        // execution state should be valid.
        request2.setName("Step1");
        request2.setExecutionState("xxxx");
        PromiseAssertUtil.assertPostMessage(
                getRootURL() + response1.getUri() + "/action/updateTaskStep",
                MessageEnum.UnknownPropertyValue.getId(),
                request2);

        // execution result should be valid.
        request2.setExecutionState(TaskExecutionStateEnum.Running.getId());
        UpdateExecutionResultRequest updateExecutionResultRequest = new UpdateExecutionResultRequest();
        updateExecutionResultRequest.setState("xxxx");
        request2.setExecutionResult(updateExecutionResultRequest);
        PromiseAssertUtil.assertPostMessage(
                getRootURL() + response1.getUri() + "/action/updateTaskStep",
                MessageEnum.UnknownPropertyValue.getId(),
                request2);
    }

    /**
     * You can update task step, and task service will do some corresponding update for you.
     */
    @Test
    public void testUpdateTaskStep()
    {
        final String name = "MyTask";
        final String description = "MyTask description";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.addTaskStep(new PostTaskStepRequest("Step1", 1000));
        request1.addTaskStep(new PostTaskStepRequest("Step2", 1000));
        request1.addTaskStep(new PostTaskStepRequest("Step3", 1000));

        // Create a task.
        final GetTaskResponse createdTtask = PromiseAssertUtil.assertPostResponse(
                getRootURL() + "/promise/v1/task/",
                request1,
                GetTaskResponse.class);

        // Set step1 to running, percentage should not change.
        final UpdateTaskStepRequest request2 = new UpdateTaskStepRequest();
        request2.setName("Step1");
        request2.setExecutionState(TaskExecutionStateEnum.Running.getId());
        final GetTaskResponse response2 = PromiseAssertUtil.assertActionResponse(
                getRootURL() + createdTtask.getUri() + "/action/updateTaskStep", 
                request2, 
                GetTaskResponse.class);
        Assert.assertEquals(0, response2.getPercentage());
        
    }

    /**
     * The task ID must exist.
     */
    @Test
    public void testNotExist()
    {
        // Get
        PromiseAssertUtil.assertGetMessage(getRootURL() + "/promise/v1/task/not_exist", MessageEnum.NotExist.getId());
        // Delete
        PromiseAssertUtil.assertDeleteMessage(getRootURL() + "/promise/v1/task/not_exist", MessageEnum.NotExist.getId());
        // Update task step action.
        UpdateTaskStepRequest request1 = new UpdateTaskStepRequest();
        PromiseAssertUtil.assertActionMessage(
                getRootURL() + "/promise/v1/task/not_exist/action/updateTaskStep",
                MessageEnum.NotExist.getId(),
                request1);

    }
}
