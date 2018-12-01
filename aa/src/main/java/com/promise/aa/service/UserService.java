package com.promise.aa.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.promise.aa.dto.CreateUserRequest;
import com.promise.aa.dto.GetUserResponse;
import com.promise.aa.model.User;
import com.promise.aa.repository.UserRepository;
import com.promise.common.PromiseException;

@Service
public class UserService
{
    @Autowired
    UserRepository userRepository;
    
    public GetUserResponse create(CreateUserRequest request)
            throws PromiseException
    {
        User user = userRepository.save(request.toModel());
        return new GetUserResponse(user);
    }
}
