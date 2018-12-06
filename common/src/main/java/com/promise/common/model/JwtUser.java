package com.promise.common.model;

import javax.validation.constraints.NotNull;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * This class holds the information that will be included in JWT token.
 *
 */
@Data
@AllArgsConstructor(access = AccessLevel.PUBLIC)
public class JwtUser
{
    @NotNull
    protected String username;
    @NotNull
    protected String partition;
    @NotNull
    protected String scope;
    @NotNull
    protected String role;
}
