package com.promise.vm.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.promise.common.model.JwtUserDto;
import com.promise.common.security.util.JwtTokenGenerator;

@RestController
@RequestMapping("/token")
public class TokenController
{
    @Value("${jwt.secret}")
    private String secret;

    @PostMapping
    public ResponseEntity<Void> login(@RequestBody final JwtUserDto request)
    {
        final HttpHeaders responseHeaders = new HttpHeaders();
        responseHeaders.set("Authorization", "Bearer " + JwtTokenGenerator.generateToken(request, secret));
        return ResponseEntity.noContent().headers(responseHeaders).build();
    }
}
