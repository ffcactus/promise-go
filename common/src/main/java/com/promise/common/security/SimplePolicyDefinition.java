package com.promise.common.security;

import java.util.ArrayList;
import java.util.List;

import javax.annotation.PostConstruct;

import org.springframework.expression.ExpressionParser;
import org.springframework.expression.spel.standard.SpelExpressionParser;
import org.springframework.stereotype.Component;

@Component
public class SimplePolicyDefinition implements PolicyDefinition
{
    private List<PolicyRule> rules;
    
    @PostConstruct
    private void init(){
        ExpressionParser exp = new SpelExpressionParser();
        rules = new ArrayList<>();
        
        PolicyRule newRule = new PolicyRule(
                "ResourceOwner", 
                "Resource owner should have access to it.", 
                exp.parseExpression("subject.username == 'username'"),  // target
                exp.parseExpression("true")); // condition
        rules.add(newRule);
    }
    
    public List<PolicyRule> getAllPolicyRules() {
        return rules;
    }
}
