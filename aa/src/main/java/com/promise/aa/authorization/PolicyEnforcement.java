package com.promise.aa.authorization;

public interface PolicyEnforcement {

    boolean check(Object subject, Object resource, Object action, Object environment);

}