package com.promise.task;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class TaskApplication implements CommandLineRunner 
{

    @Autowired
    TaskRepository repository;
    
    public static void main(String[] args) {
        SpringApplication.run(TaskApplication.class, args);
    }

    @Override
    public void run(String... args)
            throws Exception
    {
        repository.deleteAll();
        
        Task parentA = new Task("A", "ParentA", 100, null);
        repository.save(parentA);
        
        Task parentB = new Task("B", "ParentB", 50, null);
        Task childB1 = new Task("B1", "ChildB1", 50, parentB);
        Task childB2 = new Task("B2", "ChildB2", 40, parentB);        
        repository.save(parentB);
        repository.save(childB1);
        repository.save(childB2);
        
        System.out.println("--- All Tasks ---");
        repository.findAll().stream().forEach(System.out::println);
        System.out.println("--- Main Tasks ---");
        repository.findByParentIsNull().stream().forEach(System.out::println);
        System.out.println("--- Main Tasks Running ---");
        repository.findByParentIsNullAndPercentageLessThan(100).stream().forEach(System.out::println);
    }    
}
