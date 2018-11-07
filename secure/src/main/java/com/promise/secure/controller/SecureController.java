package com.promise.secure.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;

public class SecureController
{
    @GetMapping("/home")
    ResponseEntity<String> getMessage(
            @RequestHeader Map<String, String> header,
            @RequestBody String request)
    {
        return new ResponseEntity<>("Hello World.", HttpStatus.OK);
    }

    @GetMapping("/student/{id}")
    ResponseEntity<String> getStudent(@PathVariable String id)
    {
        return new ResponseEntity<>("id", HttpStatus.OK);
    }

}
