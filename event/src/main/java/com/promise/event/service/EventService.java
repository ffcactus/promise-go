package com.promise.event.service;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.promise.event.dto.PostEventRequest;

@Service
public class EventService
{
    public ResponseEntity<String> postEvent(PostEventRequest request)
    {
        return new ResponseEntity<>("", HttpStatus.OK);
    }
}
