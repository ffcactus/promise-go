package com.promise.task;

import java.util.List;
import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;


public interface TaskRepository extends MongoRepository<Task, String>
{
    public Optional<Task> findById(String id);
    public List<Task> findByParentIsNull();
    public List<Task> findByParentIsNullAndPercentageLessThan(int lessThan);    
}
