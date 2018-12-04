package com.promise.common.model;

import lombok.Data;

/**
 * Simple placeholder for info extracted from the JWT
 *
 */
@Data
public class JwtUserDto
{
    private String id;
    private String username;
    private String role;

}
