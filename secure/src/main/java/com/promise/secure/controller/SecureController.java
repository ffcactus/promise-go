package com.promise.secure.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

@RestController
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
        String username;
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        Object principal = auth.getPrincipal();
        if (principal instanceof UserDetails)
        {
            username = ((UserDetails) principal).getUsername();
        }
        else
        {
            username = principal.toString();
        }
        return new ResponseEntity<>(username, HttpStatus.OK);
    }

}
