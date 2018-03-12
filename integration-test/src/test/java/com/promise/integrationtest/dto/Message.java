package com.promise.integrationtest.dto;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Message {
	@JsonProperty(value = "ID", required = true)
	private String id;
	@JsonProperty(value = "Severity", required = true)
	private String severity;
	@JsonProperty(value = "Category", required = true)
	private String category;
	@JsonProperty(value = "CreateAt", required = true)
	private String createAt;
	@JsonProperty(value = "Description", required = true)
	private String description;
	@JsonProperty(value = "Arguments", required = false)
	private List<Argument> argument;
	@JsonProperty(value = "Supports", required = false)
	private List<Support> support;

	public String getId() {
		return id;
	}

	public void setId(String id) {
		this.id = id;
	}

	public String getSeverity() {
		return severity;
	}

	public void setSeverity(String severity) {
		this.severity = severity;
	}

	public String getCategory() {
		return category;
	}

	public void setCategory(String category) {
		this.category = category;
	}

	public String getCreateAt() {
		return createAt;
	}

	public void setCreateAt(String createAt) {
		this.createAt = createAt;
	}

	public String getDescription() {
		return description;
	}

	public void setDescription(String description) {
		this.description = description;
	}

	public List<Argument> getArgument() {
		return argument;
	}

	public void setArgument(List<Argument> argument) {
		this.argument = argument;
	}

	public List<Support> getSupport() {
		return support;
	}

	public void setSupport(List<Support> support) {
		this.support = support;
	}
}
