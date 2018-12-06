package com.promise.server.dto;

import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@ToString
public class GetServerResponse
{
    public String uri;
    public String name;
}
