package com.promise.server.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.promise.common.PromiseException;
import com.promise.common.security.ContextAwarePolicyEnforcement;
import com.promise.server.dto.AddServerRequest;
import com.promise.server.dto.GetServerResponse;
import com.promise.server.dto.PatchServerRequest;
import com.promise.server.dto.ServerCollection;
import com.promise.server.model.Server;
import com.promise.server.service.ServerService;

@RestController
public class ServerController
{
    
    @Autowired
    private ContextAwarePolicyEnforcement policy;
    
    @Autowired
    private ServerService service;

    @RequestMapping(value = "/rest/v1/vm", method = RequestMethod.GET, produces = {
            "application/json"
    })
    @PreAuthorize("hasRole('ADMIN')")
    //@Secured("ADMIN")
    public ResponseEntity<ServerCollection> getVmCollection()
    {
        return null;
    }
    
    public ResponseEntity<GetServerResponse> addServer(@RequestBody AddServerRequest request) {
        Server vm = new Server(request);
        policy.checkPermission(vm, "AddServer");
        return new ResponseEntity<GetServerResponse>(new GetServerResponse(service.addServer(vm)), HttpStatus.OK);
    }   

    @RequestMapping(value = "/rest/v1/vm/{id}", method = RequestMethod.PATCH, produces = { "application/json"})
    public ResponseEntity<GetServerResponse> patchVm(@RequestBody PatchServerRequest request)
    {
        Server vm = new Server();
        policy.checkPermission(vm, "PATCH");
        return new ResponseEntity<GetServerResponse>(HttpStatus.OK);
    }
    
    @RequestMapping(value = "/rest/v1/vm/{id}", method = RequestMethod.DELETE, produces = { "application/json"})
    public ResponseEntity<GetServerResponse> removeServer(@PathVariable String id) throws PromiseException {
        return new ResponseEntity<GetServerResponse>(new GetServerResponse(service.removeServer(id)), HttpStatus.OK);
    }
}
