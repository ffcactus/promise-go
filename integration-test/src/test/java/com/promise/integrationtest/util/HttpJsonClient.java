package com.promise.integrationtest.util;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.promise.common.PromiseToken;
import com.promise.common.response.PromiseGetResponse;

public class HttpJsonClient
{
    public static final int CONNECTION_TIMEOUT = 1000;
    public static final int READ_TIMEOUT = 1000;

    public static <R, D> ResponseEntity<PromiseGetResponse<D>> getWithType(
            String url,
            PromiseToken token,
            Class<D> responseClass)
    {
        HttpURLConnection c = null;
        try
        {
            final URL u = new URL(url);
            c = (HttpURLConnection) u.openConnection();
            c.setRequestMethod("GET");
            c.setConnectTimeout(CONNECTION_TIMEOUT);
            c.setReadTimeout(READ_TIMEOUT);
            c.setDoInput(true);
            c.setRequestProperty("Content-Type", "application/json");
            c.setRequestProperty("Accept", "application/json");
            if (token != null)
            {
                c.setRequestProperty("promise-token", token.getValue());
            }
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            InputStream inputStream;
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                    inputStream = c.getInputStream();
                    break;
                case HttpURLConnection.HTTP_NOT_FOUND:
                    inputStream = c.getErrorStream();
                    break;
                default:
                    return new ResponseEntity<>(null, HttpStatus.valueOf(status));
            }
            final BufferedReader br = new BufferedReader(new InputStreamReader(inputStream));
            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            final JavaType type = mapper.getTypeFactory()
                    .constructParametricType(PromiseGetResponse.class, responseClass);
            final PromiseGetResponse<D> response = mapper.readValue(sb.toString(), type);
            return new ResponseEntity<>(response, HttpStatus.valueOf(status));
        }
        catch (final MalformedURLException ex)
        {
            System.out.println(ex);
        }
        catch (final IOException ex)
        {
            System.out.println(ex);
        }
        finally
        {
            if (c != null)
            {
                try
                {
                    c.disconnect();
                }
                catch (final Exception ex)
                {
                    System.out.println(ex);
                }
            }
        }
        return null;
    }

    public static <R, T> ResponseEntity<T> post(
            String url,
            PromiseToken token,
            R request,
            Class<T> responseClass)
    {
        HttpURLConnection c = null;
        try
        {
            final URL u = new URL(url);
            c = (HttpURLConnection) u.openConnection();
            c.setRequestMethod("POST");
            c.setConnectTimeout(CONNECTION_TIMEOUT);
            c.setReadTimeout(READ_TIMEOUT);
            c.setDoInput(true);
            c.setDoOutput(true);
            c.setRequestProperty("Content-Type", "application/json");
            c.setRequestProperty("Accept", "application/json");
            if (token != null)
            {
                c.setRequestProperty("promise-token", token.getValue());
            }
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            final OutputStream os = c.getOutputStream();
            os.write(new ObjectMapper().writeValueAsBytes(request));
            os.flush();
            c.connect();
            final int status = c.getResponseCode();
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                case HttpURLConnection.HTTP_CREATED:
                    final BufferedReader br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    final StringBuilder sb = new StringBuilder();
                    String line;
                    while ((line = br.readLine()) != null)
                    {
                        sb.append(line + "\n");
                    }
                    br.close();
                    final ObjectMapper mapper = new ObjectMapper();
                    return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
                default:
                    return new ResponseEntity<>(null, HttpStatus.valueOf(status));
            }
        }
        catch (final MalformedURLException ex)
        {
            System.out.println(ex);
        }
        catch (final IOException ex)
        {
            System.out.println(ex);
        }
        finally
        {
            if (c != null)
            {
                try
                {
                    c.disconnect();
                }
                catch (final Exception ex)
                {
                    System.out.println(ex);
                }
            }
        }
        return null;
    }

    public static <T> ResponseEntity<T> get(String url, PromiseToken token, Class<T> responseClass)
    {
        HttpURLConnection c = null;
        try
        {
            final URL u = new URL(url);
            c = (HttpURLConnection) u.openConnection();
            c.setRequestMethod("GET");
            c.setConnectTimeout(CONNECTION_TIMEOUT);
            c.setReadTimeout(READ_TIMEOUT);
            c.setDoInput(true);
            c.setRequestProperty("Content-Type", "application/json");
            c.setRequestProperty("Accept", "application/json");
            if (token != null)
            {
                c.setRequestProperty("promise-token", token.getValue());
            }
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                    final BufferedReader br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    final StringBuilder sb = new StringBuilder();
                    String line;
                    while ((line = br.readLine()) != null)
                    {
                        sb.append(line + "\n");
                    }
                    br.close();
                    final ObjectMapper mapper = new ObjectMapper();
                    return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
                case HttpURLConnection.HTTP_NOT_FOUND:
                    return new ResponseEntity<>((T) null, HttpStatus.valueOf(status));
            }
        }
        catch (final MalformedURLException ex)
        {
            System.out.println(ex);
        }
        catch (final IOException ex)
        {
            System.out.println(ex);
        }
        finally
        {
            if (c != null)
            {
                try
                {
                    c.disconnect();
                }
                catch (final Exception ex)
                {
                    System.out.println(ex);
                }
            }
        }
        return null;
    }

    public static <T> ResponseEntity<T> delete(String url, PromiseToken token, Class<T> responseClass)
    {
        HttpURLConnection c = null;
        try
        {
            final URL u = new URL(url);
            c = (HttpURLConnection) u.openConnection();
            c.setRequestMethod("DELETE");
            c.setConnectTimeout(CONNECTION_TIMEOUT);
            c.setReadTimeout(READ_TIMEOUT);
            c.setDoInput(true);
            c.setDoOutput(true);
            c.setRequestProperty("Content-Type", "application/json");
            c.setRequestProperty("Accept", "application/json");
            if (token != null)
            {
                c.setRequestProperty("promise-token", token.getValue());
            }
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            switch (status)
            {
                case HttpURLConnection.HTTP_ACCEPTED:
                    final BufferedReader br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    final StringBuilder sb = new StringBuilder();
                    String line;
                    while ((line = br.readLine()) != null)
                    {
                        sb.append(line + "\n");
                    }
                    br.close();
                    final ObjectMapper mapper = new ObjectMapper();
                    return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
                case HttpURLConnection.HTTP_NOT_FOUND:
                    return new ResponseEntity<>(null, HttpStatus.valueOf(status));
            }
        }
        catch (final MalformedURLException ex)
        {
            System.out.println(ex);
        }
        catch (final IOException ex)
        {
            System.out.println(ex);
        }
        finally
        {
            if (c != null)
            {
                try
                {
                    c.disconnect();
                }
                catch (final Exception ex)
                {
                    System.out.println(ex);
                }
            }
        }
        return null;
    }
}
