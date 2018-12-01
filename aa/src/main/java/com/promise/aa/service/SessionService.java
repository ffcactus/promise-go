package com.promise.aa.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.promise.aa.dto.LoginRequest;
import com.promise.aa.dto.LoginResponse;
import com.promise.aa.model.User;
import com.promise.aa.repository.UserRepository;
import com.promise.common.PromiseException;

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
    public LoginResponse Login(LoginRequest request)
            throws PromiseException
    {
        Optional<User> user = userRepository.findByUsername(request.getUsername());
        if (user.isEmpty())
        {
            throw new PromiseException("Invalid username or password.");
        }

        if (user.get().getPassword() != request.getPassword())
        {
            throw new PromiseException("Invalid username or password.");
        }
        if (user.get().getToken() != null)
        {
            throw new PromiseException("Already login.");
        }
        String token = request.getUsername() + "-" + request.getPassword();
        User savedUser = user.get();
        savedUser.setToken(token);
        userRepository.save(savedUser);
        return new LoginResponse(token);
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
