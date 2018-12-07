package com.promise.common.security;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.expression.method.DefaultMethodSecurityExpressionHandler;
import org.springframework.security.access.expression.method.MethodSecurityExpressionHandler;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.method.configuration.GlobalMethodSecurityConfiguration;

@EnableGlobalMethodSecurity(prePostEnabled = true, securedEnabled = true)
public class PromiseMethodSecurityConfig extends GlobalMethodSecurityConfiguration
{
    @Autowired
    AbacPermissionEvaluator permissionEvaluator;

    @Override
    protected MethodSecurityExpressionHandler createExpressionHandler()
    {
        final DefaultMethodSecurityExpressionHandler result = new DefaultMethodSecurityExpressionHandler();
        result.setPermissionEvaluator(permissionEvaluator);
        return result;
    }

}
