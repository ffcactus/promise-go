package com.promise.server.ws;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import com.promise.common.security.manager.JwtAuthenticationSuccessHandler;

@Component
public class ServerWebSocketHandler implements WebSocketHandler
{

    private static final Logger logger = LoggerFactory.getLogger(JwtAuthenticationSuccessHandler.class);
    
    @Override
    public void afterConnectionEstablished(WebSocketSession session)
            throws Exception
    {
        logger.info("afterConnectionEstablished");

    }

    @Override
    public void handleMessage(WebSocketSession session, WebSocketMessage<?> message)
            throws Exception
    {
        logger.info("handleMessage");

    }

    @Override
    public void handleTransportError(WebSocketSession session, Throwable exception)
            throws Exception
    {
        logger.info("handleTransportError");

    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus)
            throws Exception
    {
        logger.info("afterConnectionClosed");

    }

    @Override
    public boolean supportsPartialMessages()
    {
        logger.info("supportsPartialMessages");
        return false;
    }

}
