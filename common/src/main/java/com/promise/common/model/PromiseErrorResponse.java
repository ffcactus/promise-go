package com.promise.common.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class PromiseErrorResponse
{
    private String id;
    private String summary;
    private String solution;
}
