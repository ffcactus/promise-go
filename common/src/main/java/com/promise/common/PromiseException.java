package com.promise.common;

import com.promise.common.model.PromiseError;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@Data
@EqualsAndHashCode(callSuper=true)
@NoArgsConstructor
@AllArgsConstructor
public class PromiseException extends Throwable
{

    /**
     * 
     */
    private static final long serialVersionUID = -3967614962808489460L;
    
    private String id;
    
    public PromiseError toPromiseErrorResponse() {
        return new PromiseError(id, "", "");
    }
}
