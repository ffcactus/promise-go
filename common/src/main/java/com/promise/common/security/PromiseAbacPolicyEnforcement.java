package com.promise.common.security;

import java.util.ArrayList;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.expression.EvaluationException;
import org.springframework.stereotype.Component;

import com.promise.common.security.policy.PolicyDefinition;

@Component
public class PromiseAbacPolicyEnforcement implements AbacPolicyEnforcement
{
    private static final Logger logger = LoggerFactory.getLogger(PromiseAbacPolicyEnforcement.class);

    @Autowired
    @Qualifier("promisePolicyDefinition")
    private PolicyDefinition policyDefinition;

    @Override
    public boolean check(Object subject, Object resource, Object action, Object environment)
    {
        List<PolicyRule> rules = policyDefinition.getAllPolicyRules();

        SecurityAccessContext ctx = new SecurityAccessContext(subject, resource, action, environment);
        List<PolicyRule> matchedRules = filterRules(rules, ctx);
        return checkRules(matchedRules, ctx);
    }

    private List<PolicyRule> filterRules(List<PolicyRule> allRules, SecurityAccessContext ctx)
    {
        List<PolicyRule> matchedRules = new ArrayList<PolicyRule>();
        for (PolicyRule rule : allRules)
        {
            try
            {
                if (rule.getTarget().getValue(ctx, Boolean.class))
                {
                    matchedRules.add(rule);
                }
            }
            catch (EvaluationException ex)
            {
                logger.error("An error occurred while evaluating policy rule.", ex);
            }
        }
        return matchedRules;
    }

    private boolean checkRules(List<PolicyRule> matchedRules, SecurityAccessContext cxt)
    {
        for (PolicyRule rule : matchedRules)
        {
            try
            {
                if (rule.getCondition().getValue(cxt, Boolean.class))
                {
                    return true;
                }
            }
            catch (EvaluationException ex)
            {
                logger.error("An error occurred while evaluating PolicyRule.", ex);
            }
        }
        return false;
    }
}
