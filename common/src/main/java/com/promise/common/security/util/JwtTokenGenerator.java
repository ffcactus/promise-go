package com.promise.common.security.util;

import org.springframework.beans.factory.annotation.Value;

import com.promise.common.model.JwtUser;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.impl.TextCodec;

public class JwtTokenGenerator
{
    @Value("${jwt.secret}")
    private String secret;
    
    /**
     * Generates a JWT token containing username as subject, and userId and role as
     * additional claims. These properties are taken from the specified
     * User object. Tokens validity is infinite.
     *
     * @param u the user for which the token will be generated
     * @return the JWT token
     */
    public static String generateToken(JwtUser u, String secret)
    {
        final Claims claims = Jwts.claims().setSubject(u.getUsername());
        claims.put("partition", u.getUsername());
        claims.put("scope", u.getScope());
        claims.put("role", u.getRole());      
        return Jwts.builder()
                .setClaims(claims)
                .signWith(SignatureAlgorithm.HS256, TextCodec.BASE64.decode(secret))
                .compact();
    }

}
