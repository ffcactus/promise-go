package com.promise.common.security;

import java.io.IOException;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.authentication.AbstractAuthenticationProcessingFilter;
import org.springframework.stereotype.Component;

import com.promise.common.security.exception.JwtTokenMissingException;
import com.promise.common.security.manager.JwtAuthenticationManager;
import com.promise.common.security.manager.JwtAuthenticationSuccessHandler;

/**
 * This class is the entry point of our JWT authentication process; the filter
 * extracts the JWT token from the request headers and delegates authentication
 * to the injected AuthenticationManager. If the token is not found, an
 * exception is thrown that stops the request from processing. We also need an
 * override for successful authentication because the default Spring flow would
 * stop the filter chain and proceed with a redirect. Keep in mind we need the
 * chain to execute fully, including generating the response, as explained
 * above.
 *
 */
@Component
public class JwtAuthenticationTokenFilter extends AbstractAuthenticationProcessingFilter
{

    private static final String prefix = "Bearer ";

    @Autowired
    public JwtAuthenticationTokenFilter(
            JwtAuthenticationSuccessHandler jwtAuthenticationSuccessHandler,
            JwtAuthenticationManager jwtAuthenticationManager)
    {
        super("/rest/**");
        this.setAuthenticationSuccessHandler(jwtAuthenticationSuccessHandler);
        this.setAuthenticationManager(jwtAuthenticationManager);
    }

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response)
            throws AuthenticationException,
            IOException,
            ServletException
    {
        final String header = request.getHeader("Authorization");

        if (header == null || !header.startsWith(prefix))
        {
            throw new JwtTokenMissingException("No JWT token found in request headers");
        }

        final String authToken = header.substring(prefix.length());

        final JwtAuthenticationToken authRequest = new JwtAuthenticationToken(authToken);

        return getAuthenticationManager().authenticate(authRequest);
    }

    @Override
    protected void successfulAuthentication(
            HttpServletRequest request,
            HttpServletResponse response,
            FilterChain chain,
            Authentication authResult)
            throws IOException,
            ServletException
    {
        super.successfulAuthentication(request, response, chain, authResult);

        // As this authentication is in HTTP header, after success we need to continue the request normally
        // and return the response as if the resource was not secured at all
        chain.doFilter(request, response);
    }

}
