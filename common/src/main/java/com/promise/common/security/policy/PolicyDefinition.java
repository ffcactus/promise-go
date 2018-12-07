package com.promise.common.security.policy;

import java.util.List;

import com.promise.common.security.PolicyRule;

public interface PolicyDefinition {
    public List<PolicyRule> getAllPolicyRules();
}