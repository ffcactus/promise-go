package com.promise.vm.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.promise.common.model.JwtUserDto;
import com.promise.common.security.util.JwtTokenGenerator;
import com.promise.vm.dto.VmCollection;

@RestController
public class VmController
{
    @Value("${jwt.secret}")
    private String secret;
    
    @RequestMapping(value = "/rest/v1/vm", method = RequestMethod.GET, produces = {"application/json"})
    @PreAuthorize("hasPermission(null,'PROJECTS_LIST')")
    public ResponseEntity<VmCollection> getVmCollection() {
        return null;
    }
    
    @PostMapping("/rest/v1/login")
    public ResponseEntity<Void> login(@RequestBody JwtUserDto request)
    {
        HttpHeaders responseHeaders = new HttpHeaders();
        responseHeaders.set("Authorization", "Bearer " + JwtTokenGenerator.generateToken(request, secret));
        return ResponseEntity.noContent().headers(responseHeaders).build();
    }

}
