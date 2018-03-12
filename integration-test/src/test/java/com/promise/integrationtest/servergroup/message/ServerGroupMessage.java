package com.promise.integrationtest.servergroup.message;

public enum ServerGroupMessage {
	EXIST ("MessageIDServerGroupExist"),
	NOT_EXIST("MessageIDServerGroupNotExist");
	
	private String id;
	
	ServerGroupMessage(String id) {
		this.id = id;
	}
	
	public String getId() {
		return id;
	}
}
