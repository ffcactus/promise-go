package com.promise.aa.dto;

import java.util.UUID;

import com.promise.aa.model.User;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString(exclude = {
        "password"
})
public class CreateUserRequest
{
    private String username;
    private String password;
    private String email;

    public User toModel()
    {
        return new User(
                UUID.randomUUID().toString(),
                username,
                password,
                null,
                null,
                email
        );
    }
}
