package com.promise.integrationtest.server.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * The DTO of post server.
 */
public class PostServerRequest
{
    @JsonProperty(value = "Hostname", required = true)
    private String hostname;
    @JsonProperty(value = "Username", required = true)
    private String username;
    @JsonProperty(value = "Password", required = true)
    private String password;

    public PostServerRequest()
    {

    }

    public PostServerRequest(String hostname, String username, String password)
    {
        this.hostname = hostname;
        this.username = username;
        this.password = password;
    }

    public String getHostname()
    {
        return hostname;
    }

    public void setHostname(String hostname)
    {
        this.hostname = hostname;
    }

    public String getUsername()
    {
        return username;
    }

    public void setUsername(String username)
    {
        this.username = username;
    }

    public String getPassword()
    {
        return password;
    }

    public void setPassword(String password)
    {
        this.password = password;
    }
}
