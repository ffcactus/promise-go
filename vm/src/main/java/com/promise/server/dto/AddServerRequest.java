package com.promise.server.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AddServerRequest
{
    public String name;
    public String type;
    public String partition;
    public String scope;
}
