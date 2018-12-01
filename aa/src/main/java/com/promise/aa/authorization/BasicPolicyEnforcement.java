package com.promise.aa.authorization;

import org.springframework.beans.factory.annotation.Autowired;

public class BasicPolicyEnforcement implements PolicyEnforcement 
{
    @Autowired
    private PolicyDefinition policyDefinition;
    
    @Override
    public boolean check(Object subject, Object resource, Object action, Object environment)
    {
        // TODO Auto-generated method stub
        return false;
    }

}
