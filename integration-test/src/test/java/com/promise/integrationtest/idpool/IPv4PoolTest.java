package com.promise.integrationtest.idpool;

import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.AfterClass;
import org.junit.Assert;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.promise.integrationtest.base.MessageEnum;
import com.promise.integrationtest.base.PromiseIntegrationTest;
import com.promise.integrationtest.dto.DeleteResourceResponse;
import com.promise.integrationtest.idpool.dto.AllocateIPv4Request;
import com.promise.integrationtest.idpool.dto.FreeIPv4Request;
import com.promise.integrationtest.idpool.dto.GetIPv4PoolResponse;
import com.promise.integrationtest.idpool.dto.IPv4PoolCollectionMemberResponse;
import com.promise.integrationtest.idpool.dto.IPv4RangeRequest;
import com.promise.integrationtest.idpool.dto.PostIPv4PoolRequest;
import com.promise.integrationtest.util.IPv4PoolAssertUtil;
import com.promise.integrationtest.util.PromiseAssertUtil;
import com.promise.integrationtest.util.RestClient;

public class IPv4PoolTest extends PromiseIntegrationTest
{
    @BeforeClass
    public static void setUpBeforeClass()
            throws Exception
    {

    }

    @AfterClass
    public static void tearDownAfterClass()
            throws Exception
    {
        // Remove all IPv4 pool.
        final ResponseEntity<DeleteResourceResponse> response1 = RestClient.delete(
                "/promise/v1/id-pool/ipv4",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response1.getStatusCode());
    }

    @Before
    public void setUp()
            throws Exception
    {
        // Remove all IPv4 pool.
        final ResponseEntity<DeleteResourceResponse> response1 = RestClient.delete(
                "/promise/v1/id-pool/ipv4",
                DeleteResourceResponse.class);
        Assert.assertEquals(HttpStatus.ACCEPTED, response1.getStatusCode());
    }

    @Test
    public void testHappyPath()
            throws UnsupportedEncodingException
    {
        final IPv4RangeRequest range1 = new IPv4RangeRequest("0.0.0.0", "0.0.0.255");
        final IPv4RangeRequest range2 = new IPv4RangeRequest("0.0.1.0", "0.0.1.255");

        final List<IPv4RangeRequest> ranges1 = new ArrayList<>();

        ranges1.add(range1);
        ranges1.add(range2);

        final PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();

        request1.setName("pool1");
        request1.setDescription("description.");
        request1.setRanges(ranges1);
        request1.setSubnetMask("0.0.0.0");
        request1.setGateway("0.0.0.1");
        request1.setDomain("domain");
        final String[] dns = {
                "0.0.0.2",
                "0.0.0.3"
        };
        request1.setDnsServers(Arrays.asList(dns));

        final GetIPv4PoolResponse response1 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);

        final GetIPv4PoolResponse response2 = PromiseAssertUtil.assertGetResponse(
                response1.getUri(),
                GetIPv4PoolResponse.class);
        Assert.assertEquals("pool1", response2.getName());

        PromiseAssertUtil.assertDeleteResource(response1.getUri());
    }

    /**
     * When you post a IPv4 pool that exist, it will fail.
     */
    @Test
    public void testPostExist()
    {
        final IPv4RangeRequest range1 = new IPv4RangeRequest("0.0.0.0", "0.0.0.255");
        final IPv4RangeRequest range2 = new IPv4RangeRequest("0.0.1.0", "0.0.1.255");

        final List<IPv4RangeRequest> ranges1 = new ArrayList<>();

        ranges1.add(range1);
        ranges1.add(range2);

        final PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();

        request1.setName("pool1");
        request1.setDescription("description.");
        request1.setRanges(ranges1);
        request1.setSubnetMask("0.0.0.0");
        request1.setGateway("0.0.0.1");
        request1.setDomain("domain");
        final String[] dns = {
                "0.0.0.2",
                "0.0.0.3"
        };
        request1.setDnsServers(Arrays.asList(dns));

        IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);
        PromiseAssertUtil
                .assertPostMessage("/promise/v1/id-pool/ipv4", MessageEnum.Duplicate.getId(), request1);
    }

    /**
     * When you operate on a pool that is not exist, you will fail.
     */
    @Test
    public void testPoolNotExist()
    {
        PromiseAssertUtil.assertDeleteMessage(
                "/promise/v1/id-pool/ipv4/i_am_not_exist",
                MessageEnum.NotExist.getId());
        PromiseAssertUtil.assertGetMessage(
                "/promise/v1/id-pool/ipv4/i_am_not_exist",
                MessageEnum.NotExist.getId());
        final AllocateIPv4Request request1 = new AllocateIPv4Request();
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/i_am_not_exist/action/allocate",
                MessageEnum.NotExist.getId(),
                request1);
        final FreeIPv4Request request2 = new FreeIPv4Request();
        request2.setAddress("0.0.0.0");
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/i_am_not_exist/action/free",
                MessageEnum.NotExist.getId(),
                request2);
    }

    /**
     * You can get collection.
     *
     * @throws UnsupportedEncodingException
     */
    @Test
    public void testGetCollection()
            throws UnsupportedEncodingException
    {
        final IPv4RangeRequest range = new IPv4RangeRequest("0.0.0.0", "0.0.0.1");
        final List<IPv4RangeRequest> ranges = new ArrayList<>();

        ranges.add(range);
        final PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();
        final PostIPv4PoolRequest request2 = new PostIPv4PoolRequest();
        final PostIPv4PoolRequest request3 = new PostIPv4PoolRequest();
        request1.setName("pool1");
        request1.setRanges(ranges);
        request2.setName("pool2");
        request2.setRanges(ranges);
        request3.setName("pool3");
        request3.setRanges(ranges);

        final GetIPv4PoolResponse response1 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);
        final GetIPv4PoolResponse response2 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request2);
        final GetIPv4PoolResponse response3 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request3);

        final List<IPv4PoolCollectionMemberResponse> members1 = PromiseAssertUtil
                .assertGetCollection("/promise/v1/id-pool/ipv4", 3, 3, IPv4PoolCollectionMemberResponse.class);
        Assert.assertTrue(members1.contains(response1));
        Assert.assertTrue(members1.contains(response2));
        Assert.assertTrue(members1.contains(response3));

        PromiseAssertUtil.assertDeleteResource(response1.getUri());
        final List<IPv4PoolCollectionMemberResponse> members2 = PromiseAssertUtil
                .assertGetCollection("/promise/v1/id-pool/ipv4", 2, 2, IPv4PoolCollectionMemberResponse.class);
        Assert.assertTrue(members2.contains(response2));
        Assert.assertTrue(members2.contains(response3));

        // Test start and count parameters.
        PromiseAssertUtil.assertGetColletcionWithStartCount("/promise/v1/id-pool/ipv4", 2);
    }

    /**
     * You can allocate from pool.
     */
    @Test
    public void testAllocate()
    {
        final IPv4RangeRequest range1 = new IPv4RangeRequest("0.0.0.1", "0.0.0.1");
        final IPv4RangeRequest range2 = new IPv4RangeRequest("0.0.0.2", "0.0.0.2");
        final List<IPv4RangeRequest> ranges1 = new ArrayList<>();
        ranges1.add(range1);
        ranges1.add(range2);
        final PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();
        request1.setName("pool1");
        request1.setDescription("description.");
        request1.setRanges(ranges1);
        request1.setSubnetMask("0.0.0.0");
        request1.setGateway("0.0.0.1");
        request1.setDomain("domain");
        final String[] dns = {
                "0.0.0.2",
                "0.0.0.3"
        };
        request1.setDnsServers(Arrays.asList(dns));
        final GetIPv4PoolResponse response1 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);
        IPv4PoolAssertUtil.assertIPv4Allocate(response1.getId(), null, "0.0.0.1", 2, 1, 1);
        IPv4PoolAssertUtil.assertIPv4Allocate(response1.getId(), "key", "0.0.0.2", 2, 0, 0);
        IPv4PoolAssertUtil.assertIPv4PoolEmpty(response1.getId());
    }

    /**
     * You can free address to the pool.
     */
    @Test
    public void testFree()
    {
        // Create a pool with 2 IPs.
        final IPv4RangeRequest range1 = new IPv4RangeRequest("0.0.0.1", "0.0.0.1");
        final IPv4RangeRequest range2 = new IPv4RangeRequest("0.0.0.2", "0.0.0.2");
        final List<IPv4RangeRequest> ranges1 = new ArrayList<>();
        ranges1.add(range1);
        ranges1.add(range2);
        final PostIPv4PoolRequest request1 = new PostIPv4PoolRequest();
        request1.setName("pool1");
        request1.setDescription("description.");
        request1.setRanges(ranges1);
        request1.setSubnetMask("0.0.0.0");
        request1.setGateway("0.0.0.1");
        request1.setDomain("domain");
        final String[] dns = {
                "0.0.0.2",
                "0.0.0.3"
        };
        request1.setDnsServers(Arrays.asList(dns));
        final GetIPv4PoolResponse response1 = IPv4PoolAssertUtil.assertIPv4PoolPosted(request1);
        // Allocate 2 IPs
        IPv4PoolAssertUtil.assertIPv4Allocate(response1.getId(), null, "0.0.0.1", 2, 1, 1);
        IPv4PoolAssertUtil.assertIPv4Allocate(response1.getId(), "key", "0.0.0.2", 2, 0, 0);
        // Now do the free.
        IPv4PoolAssertUtil.assertIPv4Free(response1.getId(), "0.0.0.1", 2, 1, 1);
        // Free the same key should show error message.
        IPv4PoolAssertUtil.assertIPv4AddressNotAllocated(response1.getId(), "0.0.0.1");
        // Since one of the IPs has no key, so after all the IP return to the pool, one of the IP should marked as free.
        IPv4PoolAssertUtil.assertIPv4Free(response1.getId(), "0.0.0.2", 2, 1, 2);
        // Free a IP which doesn't belong to the pool should show error message.
        IPv4PoolAssertUtil.assertIPv4PoolAddressNotBelong(response1.getId(), "1.1.1.1");
    }

    /**
     * System will check post pool request.
     */
    @Test
    public void testPostPoolRequestValidation()
    {
        final IPv4RangeRequest range1 = new IPv4RangeRequest("a.a.a.a", "0.0.0.1");
        final IPv4RangeRequest range2 = new IPv4RangeRequest("0.0.0.1", "0.0.0.0");
        final IPv4RangeRequest range3 = new IPv4RangeRequest("0.0.0.1", "0.0.1.1");
        final List<IPv4RangeRequest> ranges1 = new ArrayList<>();
        final List<IPv4RangeRequest> ranges2 = new ArrayList<>();
        final List<IPv4RangeRequest> ranges3 = new ArrayList<>();
        ranges1.add(range1);
        ranges2.add(range2);
        ranges3.add(range3);

        final PostIPv4PoolRequest request = new PostIPv4PoolRequest();
        request.setName("pool1");
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4",
                MessageEnum.IPv4PoolRangeCountError.getId(),
                request);

        request.setRanges(ranges1);
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4",
                MessageEnum.IPv4PoolFormatError.getId(),
                request);
        request.setRanges(ranges2);
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4",
                MessageEnum.IPv4PoolRangeEndAddressError.getId(),
                request);
        request.setRanges(ranges3);
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4",
                MessageEnum.IPv4PoolRangeSizeError.getId(),
                request);
    }

    /**
     * System will check action of freeing IPv4 address.
     */
    @Test
    public void testFreeIPv4RequestValidation()
    {
        final FreeIPv4Request request = new FreeIPv4Request();
        request.setAddress("a.a.a.a");
        PromiseAssertUtil.assertPostMessage(
                "/promise/v1/id-pool/ipv4/any" + "/action/free",
                MessageEnum.IPv4PoolFormatError.getId(),
                request);
    }
}
