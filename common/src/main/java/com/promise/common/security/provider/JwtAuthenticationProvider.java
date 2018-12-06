package com.promise.common.security.provider;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.authentication.dao.AbstractUserDetailsAuthenticationProvider;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import com.promise.common.model.JwtUser;
import com.promise.common.model.PromiseUserDetails;
import com.promise.common.security.JwtAuthenticationToken;
import com.promise.common.security.exception.JwtTokenMalformedException;
import com.promise.common.security.util.JwtTokenValidator;

@Component
public class JwtAuthenticationProvider extends AbstractUserDetailsAuthenticationProvider
{
    @Autowired
    private JwtTokenValidator jwtTokenValidator;

    @Override
    public boolean supports(Class<?> authentication)
    {
        return (JwtAuthenticationToken.class.isAssignableFrom(authentication));
    }

    @Override
    protected void additionalAuthenticationChecks(UserDetails userDetails, UsernamePasswordAuthenticationToken authentication)
            throws AuthenticationException
    {
    }

    @Override
    protected UserDetails retrieveUser(String username, UsernamePasswordAuthenticationToken authentication)
            throws AuthenticationException
    {
        final JwtAuthenticationToken jwtAuthenticationToken = (JwtAuthenticationToken) authentication;
        final String token = jwtAuthenticationToken.getToken();

        final JwtUser parsedUser = jwtTokenValidator.parseToken(token);

        if (parsedUser == null)
        {
            throw new JwtTokenMalformedException("JWT token is not valid");
        }
        
        PromiseUserDetails userDetails = new PromiseUserDetails(parsedUser.getUsername(), parsedUser.getPartition(), parsedUser.getScope(), parsedUser.getRole(), "");
        return userDetails;
    }
}
