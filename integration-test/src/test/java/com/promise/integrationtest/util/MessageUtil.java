package com.promise.integrationtest.util;

import java.util.ArrayList;
import java.util.List;

import com.promise.integrationtest.dto.Argument;
import com.promise.integrationtest.dto.ErrorResponse;
import com.promise.integrationtest.dto.Support;

public class ErrorResponseUtil
{
    public static ErrorResponse testErrorResponse;

    static
    {
        testErrorResponse = new ErrorResponse();
        testErrorResponse.setId("id");
        testErrorResponse.setDescription("description");
        testErrorResponse.setSeverity("severity");

        List<Argument> arguments = new ArrayList<Argument>();

        Argument argument1 = new Argument();
        argument1.setName("name");
        argument1.setType("type");
        argument1.setValue("value");

        Argument argument2 = new Argument();
        argument2.setName("name");
        argument2.setType("type");
        argument2.setValue("value");

        arguments.add(argument1);
        arguments.add(argument2);
        testErrorResponse.setArgument(arguments);

        List<Support> supports = new ArrayList<Support>();
        Support support1 = new Support();
        support1.setId("id");
        support1.setReason("reason");
        support1.setSolution("solution");
        support1.setReasonArgument(arguments);
        support1.setSolutionArgument(arguments);

        Support support2 = new Support();
        support2.setId("id");
        support2.setReason("reason");
        support2.setSolution("solution");
        support2.setReasonArgument(arguments);
        support2.setSolutionArgument(arguments);
        supports.add(support1);
        supports.add(support2);
        testErrorResponse.setSupport(supports);
    }

    /**
     * Return a error response that is used for test.
     * 
     * @return
     */
    public static ErrorResponse newTestErrorResponse()
    {
        return testErrorResponse;
    }
}
