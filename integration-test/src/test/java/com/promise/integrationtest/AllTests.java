package com.promise.integrationtest;

import org.junit.runner.RunWith;
import org.junit.runners.Suite;
import org.junit.runners.Suite.SuiteClasses;

import com.promise.integrationtest.idpool.IPv4PoolTest;
import com.promise.integrationtest.server.ServerGroupTest;
import com.promise.integrationtest.server.ServerServerGroupTest;
import com.promise.integrationtest.task.TaskTest;
import com.promise.integrationtest.ws.WsTest;

@RunWith(Suite.class)
@SuiteClasses({
        WsTest.class,
        TaskTest.class,
        IPv4PoolTest.class,
        ServerGroupTest.class,
        ServerServerGroupTest.class
})
public class AllTests
{

}
