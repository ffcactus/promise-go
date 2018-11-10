package com.promise.secure;

import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.stereotype.Component;

@Component
public class MySecondAuthenticationProvider implements AuthenticationProvider
{

    @Override
    public Authentication authenticate(Authentication authentication)
            throws AuthenticationException
    {
        System.out.println(this.getClass().toString());
        return authentication;
    }

    @Override
    public boolean supports(Class<?> authentication)
    {
        return true;
    }

}
