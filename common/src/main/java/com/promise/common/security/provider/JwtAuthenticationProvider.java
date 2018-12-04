package com.promise.common.security.provider;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.authentication.dao.AbstractUserDetailsAuthenticationProvider;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.AuthorityUtils;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Component;

import com.promise.common.model.JwtUserDto;
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

        final JwtUserDto parsedUser = jwtTokenValidator.parseToken(token);

        if (parsedUser == null)
        {
            throw new JwtTokenMalformedException("JWT token is not valid");
        }

        final List<GrantedAuthority> authorityList = AuthorityUtils.commaSeparatedStringToAuthorityList(parsedUser.getRole());

        return new PromiseUserDetails(parsedUser.getId(), parsedUser.getUsername(), token, authorityList);
    }
}
