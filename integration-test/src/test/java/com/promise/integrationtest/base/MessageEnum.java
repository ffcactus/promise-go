package com.promise.integrationtest.base;

public enum MessageEnum
{
    InternalError ("Promise.Message.InternalError"),
    NotExist ("Promise.Message.NotExist"),
    Duplicate ("Promise.Message.Duplicate"),
    InvalidRequest ("Promise.Message.InvalidRequest"),
    UnknownPropertyValue ("Promise.Message.UnknownPropertyValue"),
    Timeout ("Promise.Message.Timeout"),
    UnknownFilterName ("Promise.Message.UnknownFilterName"),

    ServerGroupDeleteDefault ("Server.Message.ServerGroupDeleteDefault"),

    ServerServerGroupDeleteDefault ("Server.Message.ServerServerGroupDeleteDefault"),

    IPv4PoolEmpty ("IPv4.Message.IPv4PoolEmpty"),
    IPv4PoolAddressNotExist ("IPv4.Message.AddressNotExist"),
    IPv4PoolFormatError ("IPv4.Message.IPv4FormatError"),
    IPv4PoolRangeEndAddressError ("IPv4.Message.IPv4RangeEndAddressError"),
    IPv4PoolRangeSizeError ("IPv4.Message.IPv4RangeSizeError"),
    IPv4PoolRangeCountError ("IPv4.Message.IPv4RangeCountError"),
    IPv4PoolNotAllocatedError ("IPv4.Message.IPv4NotAllocatedError"),

    TaskNoStep ("Task.Message.NoStep"),

    ;

    private String id;

    MessageEnum(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
