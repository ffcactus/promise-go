package com.promise.aa.authorization;

import java.io.Serializable;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.PermissionEvaluator;
import org.springframework.security.core.Authentication;

public class AbacPermissionEvaluator implements PermissionEvaluator
{
    @Autowired
    PolicyEnforcement policy;
    
    @Override
    public boolean hasPermission(Authentication authentication, Object targetDomainObject, Object permission)
    {
        // Getting subject
        Object user = authentication.getPrincipal();
        
        Map<String, Object> environment = new HashMap<>();
        environment.put("time", new Date());
        return policy.check(user, targetDomainObject, permission, environment);
    }

    @Override
    public boolean hasPermission(Authentication authentication, Serializable targetId, String targetType, Object permission)
    {
        return false;
    }

}
