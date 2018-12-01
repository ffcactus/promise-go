package com.promise.aa.authorization;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;

/**
 * This component is similar to <tt>PermissionEvaluator</tt> but can be called
 * at any point in your code,
 * given that the <tt>SecurityContext</tt> is available and filled with current,
 * authenticated user information.
 * <p>
 * This component is used when the data needed to make the access decision is
 * not available to
 * <tt>@PreAuthorize</tt> and <tt>@PostAuthorize</tt> annotations.
 * <p>
 * For example, when updating an entity:
 * <li><tt>@PreAuthorize</tt> will have access only to the method's parameters,
 * which are the updated entity's information,
 * while the access decision needs the information of the existing entity.</li>
 * <li><tt>@PostAuthorize</tt> will be called after the update is done, which is
 * too late for an access decision to be taken.</li>
 *
 */
public class ContextAwarePolicyEnforcement
{
    @Autowired
    protected PolicyEnforcement policy;

    public void checkPermission(Object resource, String permission)
    {

        //Getting the subject
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();

        //Getting the environment
        Map<String, Object> environment = new HashMap<>();
        environment.put("time", new Date());
        if (!policy.check(auth.getPrincipal(), resource, permission, environment))
        {
            throw new AccessDeniedException("Access is denied");
        }
    }
}
