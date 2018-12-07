package com.promise.common.model;

import java.util.List;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.AuthorityUtils;
import org.springframework.security.core.userdetails.UserDetails;

import com.fasterxml.jackson.annotation.JsonIgnore;

import lombok.Getter;
import lombok.Setter;

/**
 * PromiseUserDetails contains the informations that required for authentication and authorization.
 *
 */
public class PromiseUserDetails extends JwtUser implements UserDetails
{
    /**
     * 
     */
    private static final long serialVersionUID = -3994452351352130007L;
    @Getter
    @Setter
    @JsonIgnore
    protected String password;
    @Getter
    @Setter
    protected String token;
    
    public PromiseUserDetails(
            String username,
            String partition,
            String scope,
            String role,
            String password)
    {
        super(username, partition, scope, role);
        this.password = password;
    }    
    
    @Override
    @JsonIgnore
    public boolean isAccountNonExpired()
    {
        return true;
    }

    @Override
    @JsonIgnore
    public boolean isAccountNonLocked()
    {
        return true;
    }

    @Override
    @JsonIgnore
    public boolean isCredentialsNonExpired()
    {
        return true;
    }
    
    @Override
    @JsonIgnore
    public boolean isEnabled()
    {
        return true;
    }
    
    public List<GrantedAuthority> getAuthorities() {
        AuthorityUtils.commaSeparatedStringToAuthorityList(role);
        return AuthorityUtils.commaSeparatedStringToAuthorityList(role);   
    }    
}
