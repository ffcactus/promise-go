package com.promise.task;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class Controller
{
    @Autowired
    private TaskService service;
    
    @PostMapping("/tasks")
    public Task saveTask(@RequestBody Task task) {
        return service.save(task);
    }
    
    @GetMapping("/tasks")
    public List<Task> getTasks() {
        return service.getTasks();
    }

    @GetMapping("/tasks/{id}")
    public Optional<Task> getTaskById(@PathVariable String id) {
        return service.getTaskById(id);
    }
}
