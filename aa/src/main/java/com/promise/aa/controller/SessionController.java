package com.promise.aa.controller;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.promise.aa.dto.GetUserResponse;
import com.promise.aa.dto.LoginRequest;
import com.promise.aa.dto.LoginResponse;
import com.promise.aa.model.User;
import com.promise.aa.service.SessionService;
import com.promise.common.PromiseException;
import com.promise.common.model.PromiseError;

@RestController
@RequestMapping("/rest/v1/session")
public class SessionController
{
    @Autowired
    SessionService service;

    @PostMapping("/login")
    @ResponseStatus(HttpStatus.CREATED)
    public LoginResponse login(@RequestBody LoginRequest request)
    {
        return null;
    }

    @PostMapping("/logout")
    public void logout(
            @RequestHeader("promise-token") String token)
    {
        service.logout(token);
    }
    
    @GetMapping("/info")
    public ResponseEntity<GetUserResponse> info(@RequestHeader("promise-token") String token) throws PromiseException
    {
        Optional<User> user = service.info(token);
        if (user.isEmpty()) {
            return new ResponseEntity<GetUserResponse>(HttpStatus.NOT_FOUND);
        }
        
        return new ResponseEntity<GetUserResponse>(new GetUserResponse(user.get()), HttpStatus.OK);
    }

    @ExceptionHandler
    public ResponseEntity<PromiseError> handle(Throwable ex)
    {
        if (ex instanceof PromiseException)
        {
            PromiseException restException = (PromiseException) ex;
            return ResponseEntity.badRequest().body(restException.toPromiseErrorResponse());
        }
        return new ResponseEntity<PromiseError>(HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
