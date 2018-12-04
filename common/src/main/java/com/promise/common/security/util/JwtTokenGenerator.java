package com.promise.common.security.util;

import com.promise.common.model.JwtUserDto;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

public class JwtTokenGenerator
{
    /**
     * Generates a JWT token containing username as subject, and userId and role as
     * additional claims. These properties are taken from the specified
     * User object. Tokens validity is infinite.
     *
     * @param u the user for which the token will be generated
     * @return the JWT token
     */
    public static String generateToken(JwtUserDto u, String secret)
    {
        final Claims claims = Jwts.claims().setSubject(u.getUsername());
        claims.put("userId", u.getId());
        claims.put("role", u.getRole());

        return Jwts.builder()
                .setClaims(claims)
                .signWith(SignatureAlgorithm.HS512, secret)
                .compact();
    }

}
