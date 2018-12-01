package com.promise.vm.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.promise.vm.dto.VmCollection;

@RestController
@RequestMapping("/rest/v1/vm")
public class VmController
{
    @GetMapping
    public ResponseEntity<VmCollection> getVmCollection() {
        return null;
    }
}
