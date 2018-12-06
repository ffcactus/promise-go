package com.promise.aa.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.promise.aa.dto.LoginRequest;
import com.promise.aa.model.User;
import com.promise.aa.repository.UserRepository;
import com.promise.common.PromiseException;
import com.promise.common.model.JwtUser;

@Service
public class SessionService
{
    @Autowired
    UserRepository userRepository;

    /**
     * Do login operation.
     * 
     * @param request The login request.
     * @return The login response in which token is included.
     * @throws PromiseException In case like the user has already login.
     */
    public JwtUser Login(LoginRequest request)
            throws PromiseException
    {
        Optional<User> user = userRepository.findByUsername(request.getUsername());
        if (user.isEmpty())
        {
            throw new PromiseException("Invalid username or password.");
        }
        User savedUser = user.get();
        if (!savedUser.getPassword().equals(request.getPassword()))
        {
            throw new PromiseException("Invalid username or password.");
        }
        
        return new JwtUser(savedUser.getUsername(), savedUser.getPartition(), savedUser.getScope(), savedUser.getRole());
    }

    public void logout(String token)
    {
        // TODO Auto-generated method stub

    }

    public Optional<User> info(String token)
            throws PromiseException
    {
        Optional<User> user = userRepository.findByToken(token);
        if (user.isEmpty())
        {
            throw new PromiseException("Unknown token.");
        }
        return user;
    }
}
