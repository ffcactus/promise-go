package com.promise.integrationtest.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Argument {
	@JsonProperty(value = "Type", required = true)
	private String type;
	@JsonProperty(value = "Name", required = true)
	private String name;
	@JsonProperty(value = "Value", required = true)
	private String value;

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

}
