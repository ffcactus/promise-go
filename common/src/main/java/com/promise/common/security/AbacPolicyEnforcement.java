package com.promise.common.security;

public interface AbacPolicyEnforcement
{
    boolean check(Object subject, Object resource, Object action, Object environment);
}
