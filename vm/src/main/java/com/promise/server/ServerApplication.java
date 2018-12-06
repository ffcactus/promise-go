package com.promise.server;

import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.promise.server.model.Server;
import com.promise.server.repository.ServerRepository;

@SpringBootApplication
public class ServerApplication implements CommandLineRunner
{
    @Autowired
    private ServerRepository repository;
    
    @Value("${app.recreateDbOnStartup}")
    private boolean recreateDbOnStartup;
    
    public static void main(String[] args) {
        SpringApplication.run(ServerApplication.class, args);
    }
    
    @Override
    public void run(String... args)
            throws Exception
    {
        if (recreateDbOnStartup)
        {
            // HP
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "DL980 - Scope1 - " + i, "RACK", "HP", "Scope1"));
            }
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "BL980 - Scope1 - " + i, "BLADE", "HP", "Scope1"));
            }            
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "DL980 - Scope2 - " + i, "RACK", "HP", "Scope2"));
            }
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "BL980 - Scope2 - " + i, "BLADE", "HP", "Scope2"));
            }            
            // HW
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "RH2288 - Scope1 - " + i, "RACK", "HW", "Scope1"));
            }
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "CH121 - Scope1 - " + i, "BLADE", "HW", "Scope1"));
            }            
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "RH2288 - Scope2 - " + i, "RACK", "HW", "Scope2"));
            }
            for (int i = 0; i < 5000; i++) {
                repository.save(new Server(UUID.randomUUID().toString(), "CH121 - Scope2 - " + i, "BLADE", "HW", "Scope2"));
            }                            
        }
    }
    
}
