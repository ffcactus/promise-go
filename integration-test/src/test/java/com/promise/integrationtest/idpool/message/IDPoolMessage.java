package com.promise.integrationtest.idpool.message;

public enum IDPoolMessage
{
    IPv4PoolEmpty ("IPPool.Message.IPv4PoolEmpty"),
    IPv4PoolAddressNotExist ("IPPool.Message.AddressNotExist"),
    IPv4PoolFormatError ("IPPool.Message.IPv4FormatError"),
    IPv4PoolRangeEndAddressError ("IPPool.Message.IPv4RangeEndAddressError"),
    IPv4PoolRangeSizeError ("IPPool.Message.IPv4RangeSizeError"),
    IPv4PoolRangeCountError ("IPPool.Message.IPv4RangeCountError"),
    IPv4PoolNotAllocatedError ("IPPool.Message.IPv4NotAllocatedError");
    
    private String id;

    IDPoolMessage(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
