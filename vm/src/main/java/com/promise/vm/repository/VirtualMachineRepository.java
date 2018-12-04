package com.promise.vm.repository;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.promise.vm.model.VirtualMachine;

public interface VirtualMachineRepository extends MongoRepository<VirtualMachine, Long>
{
    public List<VirtualMachine> findByName(String name);
}
