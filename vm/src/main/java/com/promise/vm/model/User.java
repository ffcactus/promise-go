package com.promise.vm.model;

import javax.validation.constraints.NotNull;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class User
{
    @Id
    @NotNull
    private String id;

    @NotNull
    @Getter
    @Setter
    private String name;

    @Override
    public String toString()
    {
        return String.format("User[name=-%s, id=%s", name, id);
    }
    
    public User(String id, String name) {
        this.id = id;
        this.name = name;
    }
}
