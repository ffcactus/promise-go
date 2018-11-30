package com.promise.aa.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseStatus;

import com.promise.aa.dto.LoginRequest;
import com.promise.aa.dto.LoginResponse;
import com.promise.common.PromiseRestException;
import com.promise.common.model.PromiseErrorResponse;

public class SessionController
{
    
    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public LoginResponse login(@RequestBody LoginRequest request) {
        return null;
    }
    
    @ExceptionHandler
    public ResponseEntity<PromiseErrorResponse> handle(Throwable ex) {
        if (ex instanceof PromiseRestException) {
            PromiseRestException restException = (PromiseRestException) ex;
            return ResponseEntity.badRequest().body(restException.toPromiseErrorResponse());
        }
        return new ResponseEntity<PromiseErrorResponse>(HttpStatus.INTERNAL_SERVER_ERROR);
    }
}
