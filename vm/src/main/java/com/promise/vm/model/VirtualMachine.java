package com.promise.vm.model;

import org.springframework.data.annotation.Id;

import com.promise.common.PromiseCategory;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class VirtualMachine
{
    @Id
    public String id;
    public PromiseCategory category;
    public String name;
    public String userUri;
}
