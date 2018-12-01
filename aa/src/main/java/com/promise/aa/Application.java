package com.promise.aa;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

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
    }
}
