package com.promise.integrationtest.base;

public enum ErrorResponseEnum
{
    InternalError ("Promise.ErrorResponse.InternalError"),
    NotExist ("Promise.ErrorResponse.NotExist"),
    Duplicate ("Promise.ErrorResponse.Duplicate"),
    InvalidRequest ("Promise.ErrorResponse.InvalidRequest"),
    UnknownPropertyValue ("Promise.ErrorResponse.UnknownPropertyValue"),
    Timeout ("Promise.ErrorResponse.Timeout"),
    UnknownFilterName ("Promise.ErrorResponse.UnknownFilterName"),

    ServerGroupDeleteDefault ("ServerGroup.ErrorResponse.DeleteDefault"),

    ServerServerGroupDeleteDefault ("ServerServerGroup.ErrorResponse.DeleteDefault"),

    IPv4PoolEmpty ("IPv4.ErrorResponse.PoolEmpty"),
    IPv4PoolAddressNotExist ("IPv4.ErrorResponse.AddressNotExist"),
    IPv4PoolFormatError ("IPv4.ErrorResponse.FormatError"),
    IPv4PoolRangeEndAddressError ("IPv4.ErrorResponse.RangeEndAddressError"),
    IPv4PoolRangeSizeError ("IPv4.ErrorResponse.RangeSizeError"),
    IPv4PoolRangeCountError ("IPv4.ErrorResponse.RangeCountError"),
    IPv4PoolNotAllocatedError ("IPv4.ErrorResponse.NotAllocatedError"),

    TaskNoStep ("Task.ErrorResponse.NoStep"),

    ;

    private String id;

    ErrorResponseEnum(String id)
    {
        this.id = id;
    }

    public String getId()
    {
        return id;
    }
}
