package com.promise.common.security;

import org.springframework.security.core.GrantedAuthority;

public class PromiseGrantedAuthority implements GrantedAuthority
{

    /**
     * 
     */
    private static final long serialVersionUID = 1605007554089822869L;
    private String authority;
    
    public PromiseGrantedAuthority(String authority) {
        this.authority = authority;
    }
    

    @Override
    public String getAuthority()
    {
        return authority;
    }

}
