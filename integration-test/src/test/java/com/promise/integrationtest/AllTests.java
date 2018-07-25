package com.promise.integrationtest;

import org.junit.runner.RunWith;
import org.junit.runners.Suite;
import org.junit.runners.Suite.SuiteClasses;

import com.promise.integrationtest.idpool.IPv4PoolTest;
import com.promise.integrationtest.server.ServerGroupTest;
import com.promise.integrationtest.server.ServerServerGroupTest;
import com.promise.integrationtest.server.ServerTest;
import com.promise.integrationtest.task.TaskTest;

@RunWith(Suite.class)
@SuiteClasses({
        TaskTest.class,
        IPv4PoolTest.class,
        ServerTest.class,
        ServerGroupTest.class,
        ServerServerGroupTest.class
})
public class AllTests
{

}
