package com.promise.aa;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.promise.aa.model.User;
import com.promise.aa.repository.UserRepository;

@SpringBootApplication
public class AaApplication implements CommandLineRunner
{
    @Autowired
    private UserRepository userRepository;

    @Value("${app.recreateDbOnStartup}")
    private boolean recreateDbOnStartup;

    public static void main(String args[])
    {
        SpringApplication.run(AaApplication.class, args);
    }

    @Override
    public void run(String... args)
            throws Exception
    {
        if (recreateDbOnStartup)
        {
            userRepository.deleteAll();
            String adminAuthorities = "ROLE_ADMIN";
            String managerAuthorities = "ROLE_MANAGER";
            String readerAuthorities = "ROLE_VIEWER";

            // admin in HP.
            userRepository.save(new User("id-admin1-hp-scope1", "admin1@hp.com", "password", "admin1@hp.com", "hp", "scope1", adminAuthorities));
            userRepository.save(new User("id-admin2-hp-scope2", "admin2@hp.com", "password", "admin2@hp.com", "hp", "scope2", adminAuthorities));

            // manager in HP.
            userRepository
                    .save(new User("id-manager1-hp-scope1", "manager1@hp.com", "password", "manager1@hp.com", "hp", "scope1", managerAuthorities));
            userRepository
                    .save(new User("id-manager2-hp-scope1", "manager2@hp.com", "password", "manager2@hp.com", "hp", "scope1", managerAuthorities));
            // reader in HP.
            userRepository.save(new User("id-reader1-hp-scope1", "reader1@hp.com", "password", "reader1@hp.com", "hp", "scope1", readerAuthorities));
            userRepository.save(new User("id-reader1-hp-scope2", "reader2@hp.com", "password", "reader2@hp.com", "hp", "scope2", readerAuthorities));            

            // admin in HW.
            userRepository.save(new User("id-admin1-hw-scope1", "admin1@hw.com", "password", "admin1@hw.com", "hw", "global", adminAuthorities));
            userRepository.save(new User("id-admin2-hw-scope2", "admin2@hw.com", "password", "admin2@hw.com", "hw", "global", adminAuthorities));
            // manager in HW.
            userRepository
                    .save(new User("id-manager1-hw-scope1", "manager1@hw.com", "password", "manager1@hw.com", "hw", "global", managerAuthorities));
            userRepository
                    .save(new User("id-manager2-hw-scope1", "manager2@hw.com", "password", "manager2@hw.com", "hw", "global", managerAuthorities));
            // reader in HW.
            userRepository.save(new User("id-reader1-hw-scope1", "reader1@hw.com", "password", "reader1@hw.com", "hw", "global", readerAuthorities));
            userRepository.save(new User("id-reader1-hw-scope2", "reader2@hw.com", "password", "reader2@hw.com", "hw", "global", readerAuthorities));
        }

    }
}
