package com.promise.common;

import com.promise.common.model.PromiseErrorResponse;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@Data
@EqualsAndHashCode(callSuper=true)
@NoArgsConstructor
@AllArgsConstructor
public class PromiseRestException extends Throwable
{

    /**
     * 
     */
    private static final long serialVersionUID = -3967614962808489460L;
    
    private String id;
    
    public PromiseErrorResponse toPromiseErrorResponse() {
        return new PromiseErrorResponse(id, "", "");
    }
}
