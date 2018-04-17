package com.promise.integrationtest.dto;

import java.util.Arrays;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Support
{
    @JsonProperty(value = "ID", required = true)
    private String id;
    @JsonProperty(value = "Reason", required = true)
    private String reason;
    @JsonProperty(value = "ReasonArguments", required = false)
    private List<Argument> reasonArgument;
    @JsonProperty(value = "Solution", required = true)
    private String solution;
    @JsonProperty(value = "SolutionArguments", required = false)
    private List<Argument> solutionArgument;

    public String getId()
    {
        return id;
    }

    public void setId(String id)
    {
        this.id = id;
    }

    public String getReason()
    {
        return reason;
    }

    public void setReason(String reason)
    {
        this.reason = reason;
    }

    public List<Argument> getReasonArgument()
    {
        return reasonArgument;
    }

    public void setReasonArgument(List<Argument> reasonArgument)
    {
        this.reasonArgument = reasonArgument;
    }

    public String getSolution()
    {
        return solution;
    }

    public void setSolution(String solution)
    {
        this.solution = solution;
    }

    public List<Argument> getSolutionArgument()
    {
        return solutionArgument;
    }

    public void setSolutionArgument(List<Argument> solutionArgument)
    {
        this.solutionArgument = solutionArgument;
    }

    @Override
    public boolean equals(Object obj)
    {
        if (obj == null)
            return false;
        if (obj == this)
            return true;
        if (!(obj instanceof Support))
            return false;
        Support other = (Support) obj;
        if (!other.getId().equals(this.id))
            return false;
        if (!other.getReason().equals(this.reason))
            return false;
        if (!other.getSolution().equals(this.solution))
            return false;

        if (!Arrays.equals(this.reasonArgument.toArray(), other.reasonArgument.toArray()))
        {
            return false;
        }
        if (!Arrays.equals(this.solutionArgument.toArray(), other.solutionArgument.toArray()))
        {
            return false;
        }
        return true;
    }

}
