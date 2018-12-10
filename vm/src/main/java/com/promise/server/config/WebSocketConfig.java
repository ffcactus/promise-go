package com.promise.server.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.socket.config.annotation.WebSocketConfigurer;
import org.springframework.web.socket.config.annotation.WebSocketHandlerRegistry;

import com.promise.server.ws.ServerWebSocketHandler;

public class WebSocketConfig implements WebSocketConfigurer
{
    @Autowired
    ServerWebSocketHandler handler;
    
    @Override
    public void registerWebSocketHandlers(WebSocketHandlerRegistry registry)
    {
        registry.addHandler(handler, "/ws");
        
    }

}
