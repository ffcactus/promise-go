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
            repository.deleteAll();
            // HP
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hp-rack-scope1-" + i, "DL980 - Scope1 - " + i, "hp", "Rack", "Scope1"));
            }
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hp-blade-scope1-" + i, "BL980 - Scope1 - " + i, "hp", "Blade", "Scope1"));
            }            
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hp-rack-scope2-" + i, "DL980 - Scope2 - " + i, "hp", "Rack", "Scope2"));
            }
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hp-blade-scope2-" + i, "BL980 - Scope2 - " + i, "hp", "Blade", "Scope2"));
            }            
            // HW
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hw-rack-scope1-" + i, "RH2288 - Scope1 - " + i, "hw", "Rack", "Scope1"));
            }
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hw-blade-scope1-" + i, "CH121 - Scope1 - " + i, "hw", "Blade", "Scope1"));
            }            
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hw-rack-scope2-" + i, "RH2288 - Scope2 - " + i, "hw", "Rack", "Scope2"));
            }
            for (int i = 0; i < 5; i++) {
                repository.save(new Server("hw-blade-scope2-" + i, "CH121 - Scope2 - " + i, "hw", "Blade", "Scope2"));
            }                            
        }
    }
    
}
