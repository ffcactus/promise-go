package com.promise.aa.dto;

import com.promise.aa.model.User;

import lombok.Data;

@Data
public class GetUserResponse
{
    private String id;
    private String username;
    private String email;
    
    public GetUserResponse(User user) {
        this.id = user.getId();
        this.username = user.getUsername();
        this.email = user.getEmail();
    }
}
