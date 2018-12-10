package com.promise.common.security.manager;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.Authentication;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;
import org.springframework.stereotype.Component;

/**
 * This removes the default behavior of a successful authentication (redirecting
 * to home or any other page the user requested).
 *
 */
@Component
public class JwtAuthenticationSuccessHandler implements AuthenticationSuccessHandler
{
    private static final Logger logger = LoggerFactory.getLogger(JwtAuthenticationSuccessHandler.class);

    @Override
    public void onAuthenticationSuccess(HttpServletRequest request, HttpServletResponse response, Authentication authentication)
    {
        // We do not need to do anything extra on REST authentication success, because there is no page to redirect to
        logger.info("Authentication done. URI = " + request.getRequestURI());
    }

}
