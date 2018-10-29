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
import com.promise.integrationtest.base.ErrorResponseEnum;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.task.dto.GetTaskResponse;
import com.promise.integrationtest.task.dto.GetTaskStepResponse;
import com.promise.integrationtest.task.dto.PostTaskRequest;
import com.promise.integrationtest.task.dto.PostTaskStepRequest;
import com.promise.integrationtest.task.dto.TaskCollectionMemberResponse;
import com.promise.integrationtest.task.dto.UpdateExecutionResultRequest;
import com.promise.integrationtest.task.dto.UpdateTaskRequest;
import com.promise.integrationtest.task.dto.UpdateTaskStepRequest;
import com.promise.integrationtest.util.ErrorResponseUtil;
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
                "/promise/v1/task",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response.getStatusCode());
    }

    @Before
    public void setUp()
            throws Exception
    {
        // Remove all the server group.
        final ResponseEntity<DeleteResourceResponse> response = RestClient.delete(
                "/promise/v1/task",
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
                "/promise/v1/task/",
                request,
                GetTaskResponse.class);

        // Get it.
        final GetTaskResponse response2 = PromiseAssertUtil.assertGetResponse(
                response1.getUri(),
                GetTaskResponse.class);
        Assert.assertEquals(name, response2.getName());
        Assert.assertEquals(description, response2.getDescription());

        // Delete it.
        PromiseAssertUtil.assertDeleteResource(response1.getUri());
    }

    /**
     * The task ID must exist.
     */
    @Test
    public void testNotExist()
    {
        // Get
        PromiseAssertUtil.assertGetErrorResponse("/promise/v1/task/not_exist", HttpStatus.NOT_FOUND, ErrorResponseEnum.NotExist.getId());
        // Delete
        PromiseAssertUtil.assertDeleteErrorResponse("/promise/v1/task/not_exist", HttpStatus.NOT_FOUND, ErrorResponseEnum.NotExist.getId());
        // Update task step action.
        UpdateTaskStepRequest request1 = new UpdateTaskStepRequest();
        PromiseAssertUtil.assertActionErrorResponse(
                "/promise/v1/task/not_exist/action/updateTaskStep",
                HttpStatus.NOT_FOUND,
                ErrorResponseEnum.NotExist.getId(),
                request1);
        // Update task action.
        UpdateTaskRequest request2 = new UpdateTaskRequest();
        PromiseAssertUtil.assertActionErrorResponse(
                "/promise/v1/task/not_exist/action/updateTask",
                HttpStatus.NOT_FOUND,
                ErrorResponseEnum.NotExist.getId(),
                request2);
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

        PromiseAssertUtil.assertPostErrorResponse(
                "/promise/v1/task/",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.TaskNoStep.getId(),
                request);
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
                "/promise/v1/task/",
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
                "/promise/v1/task/",
                request,
                GetTaskResponse.class);

        // Create a task.
        request.setName("MyTask2");
        final GetTaskResponse response2 = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/task/",
                request,
                GetTaskResponse.class);

        // Test the collection is right.
        final List<TaskCollectionMemberResponse> members1 = PromiseAssertUtil
                .assertGetCollection("/promise/v1/task", 2, 2, TaskCollectionMemberResponse.class);
        Assert.assertTrue(members1.contains((Object)response1));
        Assert.assertTrue(members1.contains((Object)response2));

        // Test filter.
        final String filter1 = URLEncoder.encode("Name eq 'MyTask1'", "UTF-8");
        final List<TaskCollectionMemberResponse> members2 = PromiseAssertUtil
                .assertGetCollection(
                        "/promise/v1/task?filter=" + filter1,
                        2,
                        1,
                        TaskCollectionMemberResponse.class);
        Assert.assertTrue(members2.contains((Object)response1));

        PromiseAssertUtil.assertUnknownFilter("/promise/v1/task", "xxx", "yyy");

        // Delete a task.
        PromiseAssertUtil.assertDeleteResource(response1.getUri());

        // Test start and count.
        PromiseAssertUtil.assertGetColletcionWithStartCount("/promise/v1/task", 1);
    }

    /**
     * Task service will validate your update task request.
     */
    public void testUpdateTaskRequest()
    {
        final String name = "MyTask";
        final String description = "Description to be changed.";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.addTaskStep(new PostTaskStepRequest("Step1", 1000));

        // Create a task.
        final GetTaskResponse response1 = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/task/",
                request1,
                GetTaskResponse.class);
        UpdateTaskRequest updateTaskRequest = new UpdateTaskRequest();
        // Validate execution state.
        updateTaskRequest.setExecutionState("xxxx");
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTask",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                updateTaskRequest);
        // Validate percentage.
        updateTaskRequest.setExecutionState(TaskExecutionStateEnum.Running.getId());
        updateTaskRequest.setPercentage(100 + 1);
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTask",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                updateTaskRequest);
        // Validate execution result state.
        updateTaskRequest.setPercentage(50);
        UpdateExecutionResultRequest updateExecutionResultRequest = new UpdateExecutionResultRequest();
        updateExecutionResultRequest.setState("xxxx");
        updateTaskRequest.setExecutionResult(updateExecutionResultRequest);
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTask",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                updateTaskRequest);
    }

    /**
     * The task step should have all the properties that you updated.
     */
    @Test
    public void testUpdateTaskWithFullProperty()
    {
        final String name = "MyTask";
        final String description = "MyTask description";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.addTaskStep(new PostTaskStepRequest("Step1", 1000));

        // Create a task.
        final GetTaskResponse createdTask = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/task/",
                request1,
                GetTaskResponse.class);
        UpdateTaskRequest updateTaskRequest = new UpdateTaskRequest();
        updateTaskRequest.setDescription("description");
        updateTaskRequest.setPercentage(55);
        updateTaskRequest.setExpectedExecutionMs(1234);
        updateTaskRequest.setExecutionState(TaskExecutionStateEnum.Running.getId());

        UpdateExecutionResultRequest executionResult = new UpdateExecutionResultRequest();
        executionResult.setState(TaskExecutionResultStateEnum.Finished.getId());
        executionResult.setErrorResponse(ErrorResponseUtil.newTestErrorResponse());
        updateTaskRequest.setExecutionResult(executionResult);

        final GetTaskResponse updatedTask = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTask",
                updateTaskRequest,
                GetTaskResponse.class);

        Assert.assertEquals("description", updatedTask.getDescription());
        Assert.assertEquals(55, updatedTask.getPercentage());
        Assert.assertEquals(1234, updatedTask.getExpectedExecutionMs());
        Assert.assertEquals(TaskExecutionStateEnum.Running.getId(), updatedTask.getExecutionState());
        Assert.assertEquals(TaskExecutionResultStateEnum.Finished.getId(), updatedTask.getExecutionResult().getState());
        Assert.assertEquals(ErrorResponseUtil.newTestErrorResponse(), updatedTask.getExecutionResult().getErrorResponse());
    }

    /**
     * Task service will validate your update task step request.
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
                "/promise/v1/task/",
                request1,
                GetTaskResponse.class);

        // name should be valid.
        UpdateTaskStepRequest request2 = new UpdateTaskStepRequest();
        request2.setName("xxxx");
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTaskStep",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                request2);

        // execution state should be valid.
        request2.setName("Step1");
        request2.setExecutionState("xxxx");
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTaskStep",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                request2);

        // execution result should be valid.
        request2.setExecutionState(TaskExecutionStateEnum.Running.getId());
        UpdateExecutionResultRequest updateExecutionResultRequest = new UpdateExecutionResultRequest();
        updateExecutionResultRequest.setState("xxxx");
        request2.setExecutionResult(updateExecutionResultRequest);
        PromiseAssertUtil.assertPostErrorResponse(
                response1.getUri() + "/action/updateTaskStep",
                HttpStatus.BAD_REQUEST,
                ErrorResponseEnum.UnknownPropertyValue.getId(),
                request2);
    }

    /**
     * The task step should have all the properties that you updated.
     */
    @Test
    public void testUpdateTaskStepWithFullProperty()
    {
        final String name = "MyTask";
        final String description = "MyTask description";

        final PostTaskRequest request1 = new PostTaskRequest();
        request1.setName(name);
        request1.setDescription(description);
        request1.addTaskStep(new PostTaskStepRequest("Step1", 1000));

        // Create a task.
        final GetTaskResponse createdTask = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/task/",
                request1,
                GetTaskResponse.class);

        // Set all the properties
        final UpdateTaskStepRequest taskStepRequest = new UpdateTaskStepRequest();
        taskStepRequest.setName("Step1");
        taskStepRequest.setExecutionState(TaskExecutionStateEnum.Terminated.getId());

        UpdateExecutionResultRequest executionResult = new UpdateExecutionResultRequest();
        executionResult.setState(TaskExecutionResultStateEnum.Finished.getId());
        executionResult.setErrorResponse(ErrorResponseUtil.newTestErrorResponse());

        taskStepRequest.setExecutionResult(executionResult);
        final GetTaskResponse updatedTask = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTaskStep",
                taskStepRequest,
                GetTaskResponse.class);
        Assert.assertEquals(1, updatedTask.getTaskSteps().size());
        GetTaskStepResponse taskStep = updatedTask.getTaskSteps().get(0);
        Assert.assertEquals(TaskExecutionStateEnum.Terminated.getId(), taskStep.getExecutionState());
        Assert.assertEquals(TaskExecutionResultStateEnum.Finished.getId(), taskStep.getExecutionResult().getState());
        Assert.assertEquals(ErrorResponseUtil.newTestErrorResponse(), taskStep.getExecutionResult().getErrorResponse());
    }

    /**
     * You can update task step, and task service will do some corresponding
     * update for you.
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
        final GetTaskResponse createdTask = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/task/",
                request1,
                GetTaskResponse.class);

        // Set step1 to running, percentage should not change.
        final UpdateTaskStepRequest taskStepRequest = new UpdateTaskStepRequest();
        taskStepRequest.setName("Step1");
        taskStepRequest.setExecutionState(TaskExecutionStateEnum.Running.getId());
        final GetTaskResponse response2 = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTaskStep",
                taskStepRequest,
                GetTaskResponse.class);
        Assert.assertEquals(0, response2.getPercentage());
        Assert.assertEquals("Step1", response2.getCurrentStep());

        // Set step1 to terminated, percentage should increase.
        taskStepRequest.setExecutionState(TaskExecutionStateEnum.Terminated.getId());
        final GetTaskResponse response3 = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTaskStep",
                taskStepRequest,
                GetTaskResponse.class);
        Assert.assertEquals(33, response3.getPercentage());
        Assert.assertEquals("Step1", response3.getCurrentStep());

        // Set step2 to terminated, current task should be step2.
        taskStepRequest.setName("Step2");
        taskStepRequest.setExecutionState(TaskExecutionStateEnum.Terminated.getId());
        final GetTaskResponse response4 = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTaskStep",
                taskStepRequest,
                GetTaskResponse.class);
        Assert.assertEquals(67, response4.getPercentage());
        Assert.assertEquals("Step2", response4.getCurrentStep());

        // Set step3 to terminated, percentage should be 100.
        taskStepRequest.setName("Step3");
        taskStepRequest.setExecutionState(TaskExecutionStateEnum.Terminated.getId());
        final GetTaskResponse response5 = PromiseAssertUtil.assertActionResponse(
                createdTask.getUri() + "/action/updateTaskStep",
                taskStepRequest,
                GetTaskResponse.class);
        Assert.assertEquals(100, response5.getPercentage());
        Assert.assertEquals("Step3", response5.getCurrentStep());
    }

}
