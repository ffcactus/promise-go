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

    ServerGroupDeleteDefault ("ServerGroup.Message.DeleteDefault"),

    ServerServerGroupDeleteDefault ("ServerServerGroup.Message.DeleteDefault"),

    IPv4PoolEmpty ("IPv4.Message.PoolEmpty"),
    IPv4PoolAddressNotExist ("IPv4.Message.AddressNotExist"),
    IPv4PoolFormatError ("IPv4.Message.FormatError"),
    IPv4PoolRangeEndAddressError ("IPv4.Message.RangeEndAddressError"),
    IPv4PoolRangeSizeError ("IPv4.Message.RangeSizeError"),
    IPv4PoolRangeCountError ("IPv4.Message.RangeCountError"),
    IPv4PoolNotAllocatedError ("IPv4.Message.NotAllocatedError"),

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
