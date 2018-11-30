package com.promise.task;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TaskService
{
    @Autowired
    TaskRepository repository;
    
    public List<Task> getTasks() {
        return repository.findAll();
        
    }
    
    public Optional<Task> getTaskById(String id) {
        return repository.findById(id);
    }
    
    public List<Task> getParentTasks() {
        return repository.findByParentIsNull();
    }
    
    public List<Task> getParentTaskRunning() {
        return repository.findByParentIsNullAndPercentageLessThan(100);
    }

    public Task save(Task task)
    {
        return repository.save(task);
    }
}
