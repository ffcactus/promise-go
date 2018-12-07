package com.promise.common.security;

import org.springframework.expression.Expression;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class PolicyRule
{
    private String name;
    private String description;
    /*
     * Boolean SpEL expression. If evaluated to true, then this rule is applied to  the request access context.
     */ 
    private Expression target;
    
    /*
     * Boolean SpEL expression, if evaluated to true, then access granted.
     */
    private Expression condition;
}
