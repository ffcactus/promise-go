package com.promise.vm.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.promise.vm.dto.VmCollection;

@RestController
public class VmController
{

    @RequestMapping(value = "/rest/v1/vm", method = RequestMethod.GET, produces = {
            "application/json"
    })
    @PreAuthorize("hasPermission(null,'PROJECTS_LIST')")
    public ResponseEntity<VmCollection> getVmCollection()
    {
        return null;
    }
}
