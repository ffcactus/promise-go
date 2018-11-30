package com.promise.vm.model;

import javax.validation.constraints.NotNull;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class User
{
    @Id
    @NotNull
    private String id;

    @NotNull
    private String name;


}
