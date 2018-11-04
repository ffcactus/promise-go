package com.promise.event.controller;

import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import com.promise.event.dto.PostEventRequest;
import com.promise.event.service.EventService;

@RestController
@EnableAutoConfiguration
public class EventController
{
    @Autowired
    private EventService service;

    /**
     * Post an event.
     *
     * @param header  The request header.
     * @param request The request body.
     * @return The HTTP status and body.
     */
    ResponseEntity<String> postEvent(
            @RequestHeader Map<String, String> header,
            @RequestBody PostEventRequest request)
    {
        return service.postEvent(request);
    }
}
