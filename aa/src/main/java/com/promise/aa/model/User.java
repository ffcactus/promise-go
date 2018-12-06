package com.promise.aa.model;

import javax.validation.constraints.NotNull;

import org.springframework.data.annotation.Id;

import com.promise.common.model.PromiseUserDetails;

import lombok.Getter;
import lombok.Setter;

public class User extends PromiseUserDetails
{ 

    /**
     * 
     */
    private static final long serialVersionUID = -5031495025980981775L;
    @Getter
    @Setter
    @NotNull
    @Id
    private String id;
    
    @Getter
    @Setter
    @NotNull
    private String email;
    
    public User(String id, String username, String password, String email, String partition, String scope, String rawAuthorities) {
        super(username, partition, scope, rawAuthorities, password);
        this.setPassword(password);
        this.id = id;
        this.email = email;
    }
}
