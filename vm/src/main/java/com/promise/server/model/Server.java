package com.promise.server.model;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class Server
{
    @Id
    public String id;
    public String name;
    public String partition;
    public String scope;
    public String type;
}
