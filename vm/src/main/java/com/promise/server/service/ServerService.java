package com.promise.server.service;

import java.util.Optional;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.promise.common.PromiseException;
import com.promise.common.security.ContextAwarePolicyEnforcement;
import com.promise.server.model.Server;
import com.promise.server.repository.ServerRepository;

@Service
public class ServerService
{
    @Autowired
    private ContextAwarePolicyEnforcement policy;
    
    @Autowired
    private ServerRepository repository;
    
    public Server addServer(Server basic) {
        basic.id = UUID.randomUUID().toString();
        return repository.save(basic);
    }
    
    public Server removeServer(String id) throws PromiseException {
        Optional<Server> optionalServer = repository.findById(id);
        if (optionalServer.isEmpty()) {
            throw new PromiseException("Not found.");
        }
        Server server = optionalServer.get();
        policy.checkPermission(server, "RemoveServer");
        repository.delete(server);
        return server;
    }
}
