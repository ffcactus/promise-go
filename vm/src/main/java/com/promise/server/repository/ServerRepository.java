package com.promise.server.repository;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.promise.server.model.Server;

public interface ServerRepository extends MongoRepository<Server, String>
{
    public List<Server> findByName(String name);
}
