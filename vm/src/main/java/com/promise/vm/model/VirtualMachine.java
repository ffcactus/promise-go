package com.promise.vm.model;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class VirtualMachine
{
    @Id
    public String id;
    public String name;
    
    public String toString() {
        return String.format("VirtualMachine[name=%s, id=%l]", name, id);
    }
}
