package com.promise.common.security;

import java.io.IOException;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;

import org.springframework.web.filter.GenericFilterBean;

/**
 * 
 * This filter should replace <tt>SecurityContextPersistenceFilter</tt> for JWT based authentication.
 * It get the principle by the from the token and save it to the <tt>SecurityContext</tt>
 *
 */
public class JwtFilter extends GenericFilterBean 
{

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException,
            ServletException
    {
        HttpServletRequest httpRequest = (HttpServletRequest) request;
        System.out.println("JWT: " + httpRequest.getHeader("x-auth-token"));
        chain.doFilter(request, response);
    }

}
