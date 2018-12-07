package com.promise.common.security.policy;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

import javax.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.Resource;
import org.springframework.expression.Expression;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.module.SimpleModule;
import com.promise.common.security.PolicyRule;
import com.promise.common.security.SpelDeserializer;

@Component("promisePolicyDefinition")
public class PromisePolicyDefinition implements PolicyDefinition
{
    private static Logger logger = LoggerFactory.getLogger(PromisePolicyDefinition.class);

    @Value("classpath:default-policy.json")
    private Resource jsonDefinitionResource;

    private List<PolicyRule> rules;

    @PostConstruct
    private void init()
    {
        ObjectMapper mapper = new ObjectMapper();
        SimpleModule module = new SimpleModule();
        module.addDeserializer(Expression.class, new SpelDeserializer());
        mapper.registerModule(module);
        try
        {
            PolicyRule[] rulesArray = null;
            rulesArray = mapper.readValue(jsonDefinitionResource.getInputStream(), PolicyRule[].class);
            this.rules = (rulesArray != null ? Arrays.asList(rulesArray) : null);
            logger.info("[init] Policy loaded successfully.");
        }
        catch (JsonMappingException e)
        {
            logger.error("An error occurred while parsing the policy file.", e);
        }
        catch (IOException e)
        {
            logger.error("An error occurred while reading the policy file.", e);
        }
    }

    @Override
    public List<PolicyRule> getAllPolicyRules()
    {
        return rules;
    }

}
