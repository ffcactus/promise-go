package com.promise.common.security.exception;

import org.springframework.security.core.AuthenticationException;

public class JwtTokenMalformedException extends AuthenticationException
{
    /**
     * 
     */
    private static final long serialVersionUID = 2995466721023880669L;

    public JwtTokenMalformedException(String msg)
    {
        super(msg);
        // TODO Auto-generated constructor stub
    }

}
