package com.promise.server.ws;

import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

public class ServerWebSocketHandler implements WebSocketHandler
{

    @Override
    public void afterConnectionEstablished(WebSocketSession session)
            throws Exception
    {
        // TODO Auto-generated method stub

    }

    @Override
    public void handleMessage(WebSocketSession session, WebSocketMessage<?> message)
            throws Exception
    {
        // TODO Auto-generated method stub

    }

    @Override
    public void handleTransportError(WebSocketSession session, Throwable exception)
            throws Exception
    {
        // TODO Auto-generated method stub

    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus)
            throws Exception
    {
        // TODO Auto-generated method stub

    }

    @Override
    public boolean supportsPartialMessages()
    {
        // TODO Auto-generated method stub
        return false;
    }

}
