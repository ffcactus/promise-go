package com.promise.server.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.access.PermissionEvaluator;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import com.promise.common.security.AbacPermissionEvaluator;
import com.promise.common.security.JwtAuthenticationTokenFilter;
import com.promise.common.security.RestAuthenticationEntryPoint;

@Configuration
@ComponentScan(basePackages = {
        "com.promise.vm", "com.promise.common"
})
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true, securedEnabled = true)
public class RootConfigure extends WebSecurityConfigurerAdapter
{

    @Autowired
    RestAuthenticationEntryPoint entryPoint;
    @Autowired
    JwtAuthenticationTokenFilter jwtAuthenticationFilter;

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
        http
                .addFilterBefore(jwtAuthenticationFilter, UsernamePasswordAuthenticationFilter.class);

    }

    @Bean
    public PermissionEvaluator getAbacPermissionEvaluator()
    {
        return new AbacPermissionEvaluator();
    }
}
