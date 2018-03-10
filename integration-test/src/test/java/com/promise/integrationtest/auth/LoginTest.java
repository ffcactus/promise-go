package com.promise.integrationtest.auth;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

public class LoginTest
{
    @BeforeClass
    public static void setUpBeforeClass()
            throws Exception
    {
        System.out.println("setUpBeforeClass");
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
    public void testLogin()
    {
        Assert.assertEquals(true, true);
    }
}
