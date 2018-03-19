package com.promise.integrationtest.base;

import java.io.File;
import java.io.IOException;
import java.net.URL;
import java.util.Scanner;

import org.junit.Assert;

import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class PromiseIntegrationTest
{

    private static Environment environment;

    public PromiseIntegrationTest()
    {
        final ObjectMapper mapper = new ObjectMapper();
        try
        {
            final String text = getFile("environment.json");
            environment = mapper.readValue(text, Environment.class);
        }
        catch (final JsonParseException e)
        {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        catch (final JsonMappingException e)
        {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        catch (final IOException e)
        {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }

    public Environment getEnvironment()
    {
        return environment;
    }

    public static String getRootURL()
    {
        return "http://" + environment.getHostname();
    }

    private String getFile(String fileName)
    {

        final StringBuilder result = new StringBuilder("");

        // Get file from resources folder
        final ClassLoader classLoader = getClass().getClassLoader();
        final URL url = classLoader.getResource(fileName);
        if (url == null || url.getFile() == null)
        {
            Assert.fail("No environment.json file.");
        }
        final File file = new File(url.getFile());

        try (Scanner scanner = new Scanner(file))
        {

            while (scanner.hasNextLine())
            {
                final String line = scanner.nextLine();
                result.append(line).append("\n");
            }

            scanner.close();

        }
        catch (final IOException e)
        {
            e.printStackTrace();
        }

        return result.toString();

    }
}
