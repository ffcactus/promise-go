package com.promise.event.dto;

import com.promise.common.PromiseCategory;

public class PostEventRequest
{
    public Severity severity;
    public PromiseCategory category;
    public String eventId;
    public String source;
    public String index;
}
