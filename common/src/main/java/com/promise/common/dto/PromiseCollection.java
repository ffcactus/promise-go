package com.promise.common.dto;

import java.util.List;

import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
public class PromiseCollection<T>
{
    private long total;
    private long skip;
    private long top;
    private List<T> members;
}
