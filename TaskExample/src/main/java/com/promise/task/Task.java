package com.promise.task;

import javax.validation.constraints.NotNull;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import lombok.ToString;

@NoArgsConstructor
@AllArgsConstructor
@ToString
public class Task
{
    @Id
    @NotNull
    private String id;

    @NotNull
    private String name;

    @NotNull
    private int percentage;

    @NotNull
    private Task parent;
}
