package com.promise.server.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.access.PermissionEvaluator;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;

import com.promise.common.security.AbacPermissionEvaluator;
import com.promise.common.security.RestAuthenticationEntryPoint;
import com.promise.common.security.provider.JwtAuthenticationProvider;

@Configuration
@ComponentScan(basePackages = {
        "com.promise.vm", "com.promise.common"
})
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true, securedEnabled = true)
public class RootConfigure extends WebSecurityConfigurerAdapter
{

    @Autowired
    private RestAuthenticationEntryPoint entryPoint;

    @Autowired
    private JwtAuthenticationProvider provider;

    @Override
    protected void configure(HttpSecurity http)
            throws Exception
    {
        http
                .csrf().disable()
                .exceptionHandling().authenticationEntryPoint(entryPoint)
                .and()
                .authorizeRequests().antMatchers("**/rest/v1/vm/**/**").authenticated()
                .and()
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS);
    }
    
    protected void configure(AuthenticationManagerBuilder auth) {
        auth.authenticationProvider(provider);
    }

    @Bean
    public PermissionEvaluator getAbacPermissionEvaluator()
    {
        return new AbacPermissionEvaluator();
    }
}
