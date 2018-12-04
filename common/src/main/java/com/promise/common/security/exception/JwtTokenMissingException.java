package com.promise.common.security.exception;

import org.springframework.security.core.AuthenticationException;

public class JwtTokenMissingException extends AuthenticationException {


    /**
     * 
     */
    private static final long serialVersionUID = 2720832927506740003L;

    public JwtTokenMissingException(String msg) {
        super(msg);
    }
}
