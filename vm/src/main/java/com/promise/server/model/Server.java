package com.promise.server.model;

import org.springframework.data.annotation.Id;

import com.promise.server.dto.AddServerRequest;

import lombok.AllArgsConstructor;
import lombok.ToString;

@AllArgsConstructor
@ToString
public class Server
{
    @Id
    public String id;
    public String name;
    public String partition;
    public String scope;
    public String type;

    public Server() {
        
    }
    
    public Server(AddServerRequest request) {
        this.name = request.name;
        this.partition = request.partition;
        this.scope = request.scope;
        this.type = request.type;
    }
}
