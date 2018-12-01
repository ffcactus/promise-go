package com.promise.aa.model;

import javax.validation.constraints.NotNull;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString(exclude = {
        "password", "token"
})
public class User
{
    @Id
    @NotNull
    private String id;
    
    @NotNull
    @Indexed(unique=true)
    private String username;
    @NotNull 
    private String password;
    private String token;
    private String role;
    @NotNull
    @Indexed(unique=true)    
    private String email;
}
