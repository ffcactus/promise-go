package com.promise.vm;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.promise.vm.model.User;
import com.promise.vm.repository.UserRepository;
import com.promise.vm.repository.VirtualMachineRepository;

@SpringBootApplication
public class VmApplication implements CommandLineRunner 
{
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private VirtualMachineRepository vmRepository;
    
    public static void main(String[] args) {
        SpringApplication.run(VmApplication.class, args);
    }
    
    @Override
    public void run(String... args) throws Exception {
        System.out.println("-------------------");
        userRepository.deleteAll();
        vmRepository.deleteAll();
        
        userRepository.save(new User("1", "BaiBin"));
        userRepository.save(new User("2", "YeZhiFan"));
        userRepository.save(new User("3", "BaiXiaoBin"));
        
        for (User user : userRepository.findAll()) {
            System.out.println(user);
        }
        System.out.println(userRepository.findByNameContains("Bai"));
    }
}
