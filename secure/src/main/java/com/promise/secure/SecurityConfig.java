package com.promise.secure;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;

@Configuration
@EnableWebSecurity
@ComponentScan("com.promise")
public class SecurityConfig extends WebSecurityConfigurerAdapter
{
    @Autowired
    private MyFirstAuthenticationProvider provider1;
    @Autowired
    private MySecondAuthenticationProvider provider2;

//    protected void configure(HttpSecurity security)
//            throws Exception
//    {
//    }

    protected void configure(AuthenticationManagerBuilder auth)
            throws Exception
    {
        auth.authenticationProvider(provider1).authenticationProvider(provider2);
    }

}
