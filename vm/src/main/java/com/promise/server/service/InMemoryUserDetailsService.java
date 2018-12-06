package com.promise.server.service;

import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
public class InMemoryUserDetailsService implements UserDetailsService
{

    @Override
    public UserDetails loadUserByUsername(String username)
            throws UsernameNotFoundException
    {
        // TODO Auto-generated method stub
        return null;
    }
//    private Map<String, ProjectSecurityUser> users = new HashMap<>();
//    
//    public InMemoryUserDetailsService() {
//    }
//    
//    @PostConstruct
//    private void init() {
//        this.users = new HashMap<>();
//        users.put("admin", new ProjectSecurityUser("admin", "password", UserRole.ADMIN));       
//        users.put("pm1", new ProjectSecurityUser("pm1", "password", UserRole.PM));
//        users.put("pm2", new ProjectSecurityUser("pm2", "password", UserRole.PM));
//        users.put("dev1", new ProjectSecurityUser("dev1", "password", UserRole.DEVELOPER));
//        users.put("dev2", new ProjectSecurityUser("dev2", "password", UserRole.DEVELOPER));
//        users.put("test1", new ProjectSecurityUser("test1", "password", UserRole.TESTER));
//        users.put("test2",new ProjectSecurityUser("test2", "password", UserRole.TESTER));
//    }
}
