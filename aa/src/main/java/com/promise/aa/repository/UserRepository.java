package com.promise.aa.repository;

import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.promise.aa.model.User;

public interface UserRepository extends MongoRepository<User, String>
{
    Optional<User> findByUsername(String username);

    Optional<User> findByToken(String token);
}
