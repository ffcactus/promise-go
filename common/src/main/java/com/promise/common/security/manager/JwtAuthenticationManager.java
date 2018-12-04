package com.promise.common.security.manager;

import java.util.Arrays;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.ProviderManager;
import org.springframework.stereotype.Component;

import com.promise.common.security.provider.JwtAuthenticationProvider;

@Component
public class JwtAuthenticationManager extends ProviderManager
{
    @Autowired
    private static JwtAuthenticationProvider provider;

    public JwtAuthenticationManager()
    {
        super(Arrays.asList(provider));
    }

}
