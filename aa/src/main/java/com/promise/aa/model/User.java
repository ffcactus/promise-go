package com.promise.aa.model;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@ToString(exclude= {"password", "token"})
public class User
{
    private String id;
    private String username;
    private String password;
    private String token;
    private String role;
    private String email;
}
