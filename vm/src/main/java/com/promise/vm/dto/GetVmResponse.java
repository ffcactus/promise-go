package com.promise.vm.dto;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@ToString
public class GetVmResponse
{
    public String uri;
    public String name;
}
