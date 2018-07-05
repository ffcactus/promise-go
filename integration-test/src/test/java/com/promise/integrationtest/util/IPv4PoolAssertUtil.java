package com.promise.integrationtest.util;

import org.junit.Assert;
import org.springframework.http.HttpStatus;

import com.promise.integrationtest.base.MessageEnum;
import com.promise.integrationtest.idpool.dto.AllocateIPv4Request;
import com.promise.integrationtest.idpool.dto.AllocateIPv4Response;
import com.promise.integrationtest.idpool.dto.FreeIPv4Request;
import com.promise.integrationtest.idpool.dto.GetIPv4PoolResponse;
import com.promise.integrationtest.idpool.dto.PostIPv4PoolRequest;

public class IPv4PoolAssertUtil
{

    /*
     * Assert the IPv4 pool has been posted.
     */
    public static GetIPv4PoolResponse assertIPv4PoolPosted(PostIPv4PoolRequest request)
    {
        GetIPv4PoolResponse response = PromiseAssertUtil.assertPostResponse(
                "/promise/v1/id-pool/ipv4",
                request,
                GetIPv4PoolResponse.class);
        Assert.assertEquals(request.getName(), response.getName());
        Assert.assertEquals(request.getDescription(), response.getDescription());
        Assert.assertEquals(request.getSubnetMask(), response.getSubnetMask());
        Assert.assertEquals(request.getGateway(), response.getGateway());
        Assert.assertEquals(request.getDomain(), response.getDomain());
        PromiseAssertUtil.assertSameElements(request.getDnsServers(), response.getDnsServers());
        return response;
    }

    /**
     * Assert the address be allocated from pool.
     * 
     * @param id The pool ID.
     * @param key The key for the address.
     * @param address The address be allocated.
     * @param total The total address of the pool.
     * @param free The free address count of the pool.
     * @param allocatable The allocatable addresses count of the pool.
     * @return The address allocated.
     */
    public static void assertIPv4Allocate(String id, String key, String address, int total, int free, int allocatable)
    {
        final AllocateIPv4Request request = new AllocateIPv4Request();
        request.setKey(key);
        AllocateIPv4Response response = PromiseAssertUtil.assertActionResponse(
                "/promise/v1/id-pool/ipv4/" + id + "/action/allocate",
                request,
                AllocateIPv4Response.class);
        Assert.assertEquals(total, response.getPool().getTotal());
        Assert.assertEquals(free, response.getPool().getFree());
        Assert.assertEquals(allocatable, response.getPool().getAllocatable());
    }

    /**
     * Assert the address be free to pool.
     * 
     * @param id The pool ID.
     * @param address The address to free.
     * @param total The total address of the pool.
     * @param free The free address count of the pool.
     * @param allocatable The allocatable addresses count of the pool.
     */
    public static void assertIPv4Free(String id, String address, int total, int free, int allocatable)
    {
        final FreeIPv4Request request = new FreeIPv4Request();
        request.setAddress(address);
        GetIPv4PoolResponse response = PromiseAssertUtil.assertActionResponse(
                "/promise/v1/id-pool/ipv4/" + id + "/action/free",
                request,
                GetIPv4PoolResponse.class);
        Assert.assertEquals(total, response.getTotal());
        Assert.assertEquals(free, response.getFree());
        Assert.assertEquals(allocatable, response.getAllocatable());
    }

    /**
     * Assert the IPv4 Pool have no more address to allocate.
     * 
     * @param id
     */
    public static void assertIPv4PoolEmpty(String id)
    {
        final AllocateIPv4Request request = new AllocateIPv4Request();
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/" + id + "/action/allocate",
                HttpStatus.BAD_REQUEST,
                MessageEnum.IPv4PoolEmpty.getId(),
                request);

        final GetIPv4PoolResponse response = PromiseAssertUtil.assertGetResponse(
                "/promise/v1/id-pool/ipv4/" + id,
                GetIPv4PoolResponse.class);
        Assert.assertEquals(0, response.getAllocatable());
    }

    /**
     * Assert the message is right when you free an address that not belong to
     * this pool.
     * 
     * @param id The pool ID.
     */
    public static void assertIPv4PoolAddressNotBelong(String id, String address)
    {
        final FreeIPv4Request request = new FreeIPv4Request();
        request.setAddress(address);
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/" + id + "/action/free",
                HttpStatus.BAD_REQUEST,
                MessageEnum.IPv4PoolAddressNotExist.getId(),
                request);
    }

    /**
     * Assert the message is right when you free an address which is is not
     * allocated before.
     * 
     * @param id The pool ID.
     * @param address The address to free.
     */
    public static void assertIPv4AddressNotAllocated(String id, String address)
    {
        final FreeIPv4Request request = new FreeIPv4Request();
        request.setAddress(address);
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/" + id + "/action/free",
                HttpStatus.BAD_REQUEST,
                MessageEnum.IPv4PoolNotAllocatedError.getId(),
                request);
    }
}
