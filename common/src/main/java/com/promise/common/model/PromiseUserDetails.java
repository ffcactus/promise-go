package com.promise.common.model;

import java.util.Collection;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

/**
 * In Promise, the request token is taken as the user, since it has all the information about the user.
 *
 */
public class PromiseUserDetails implements UserDetails
{

    /**
     * 
     */
    private static final long serialVersionUID = 8192382807339090886L;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities()
    {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public String getPassword()
    {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public String getUsername()
    {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public boolean isAccountNonExpired()
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean isAccountNonLocked()
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean isCredentialsNonExpired()
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean isEnabled()
    {
        // TODO Auto-generated method stub
        return false;
    }

}
