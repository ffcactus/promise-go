package com.promise.task;

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
public class Task
{

    public Task()
    {

    }

    public Task(String id, String name, int percentage, Task parent)
    {
        this.id = id;
        this.name = name;
        this.percentage = percentage;
        this.parent = parent;
    }

    @Id
    @NotNull
    private String id;

    @NotNull
    @Getter
    @Setter
    private String name;

    @NotNull
    @Getter
    @Setter
    private int percentage;

    @NotNull
    @Getter
    @Setter
    private Task parent;
    
    @Override
    public String toString() {
        return String.format("{ID=%s, Name=%s, Percentage=%d, Parent=%s}", this.id, this.name, this.percentage, this.parent);
    }
}
