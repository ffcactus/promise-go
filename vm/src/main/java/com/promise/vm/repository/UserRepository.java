package com.promise.vm.repository;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.promise.vm.model.User;

public interface UserRepository extends MongoRepository<User, Long>
{
    public List<User> findByName(String name);
    public List<User> findByNameContains(String sub);
}
