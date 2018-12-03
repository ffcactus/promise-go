package com.promise.aa;

import java.util.Arrays;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.promise.aa.model.User;
import com.promise.aa.repository.UserRepository;

@SpringBootApplication
public class Application implements CommandLineRunner
{
    @Autowired
    private UserRepository userRepository;
    
    public static void main(String args[])
    {
        SpringApplication.run(Application.class, args);
    }

    @Override
    public void run(String... args)
            throws Exception
    {
        userRepository.deleteAll();
        User admin = new User("admin", "admin", "password", "token", "Admin", "admin@email.com");
        User pm1 = new User("pm1", "pm1", "password", "token", "Manager", "manager1@email.com");
        User developer1 = new User("user1", "user1", "password", "token", "Developer", "user1@email.com");
        User developer2 = new User("user2", "user2", "password", "token", "Developer", "user2@email.com");
        
        userRepository.saveAll(Arrays.asList(admin, pm1, developer1, developer2));
    }
}
