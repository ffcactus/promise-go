package com.promise.integrationtest.util;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.promise.integrationtest.Environment;

public class RestClient
{
    public static final int CONNECTION_TIMEOUT = 5000;
    public static final int READ_TIMEOUT = 5000;

    /**
     * The general HTTP GET method.
     *
     */
    public static <T> ResponseEntity<T> get(String url, Class<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
            }

            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
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

    /**
     * The general HTTP GET method.
     *
     */
    public static <T> ResponseEntity<T> get(String url, TypeReference<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
            }

            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            final T v = mapper.readValue(sb.toString(), responseClass);
            return new ResponseEntity<>(v, HttpStatus.valueOf(status));

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

    public static <R, T> ResponseEntity<T> post(String url, R request, Class<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            final OutputStream os = c.getOutputStream();
            os.write(new ObjectMapper().writeValueAsBytes(request));
            os.flush();
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_OK:
                case HttpURLConnection.HTTP_CREATED:
                case HttpURLConnection.HTTP_ACCEPTED:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
            }
            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
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

    public static <R, T> ResponseEntity<T> post(String url, R request, TypeReference<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            final OutputStream os = c.getOutputStream();
            os.write(new ObjectMapper().writeValueAsBytes(request));
            os.flush();
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_ACCEPTED:
                case HttpURLConnection.HTTP_OK:
                case HttpURLConnection.HTTP_CREATED:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
            }
            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            final T v = mapper.readValue(sb.toString(), responseClass);
            return new ResponseEntity<>(v, HttpStatus.valueOf(status));
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

    public static <T> ResponseEntity<T> delete(String url, Class<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                case HttpURLConnection.HTTP_ACCEPTED:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
            }
            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();
            final ObjectMapper mapper = new ObjectMapper();
            return new ResponseEntity<>(mapper.readValue(sb.toString(), responseClass), HttpStatus.valueOf(status));
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

    public static <T> ResponseEntity<T> delete(String url, TypeReference<T> responseClass)
    {
        url = "http://" + Environment.getHostname() + url;
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
            c.setUseCaches(false);
            c.setAllowUserInteraction(false);
            c.connect();
            final int status = c.getResponseCode();
            BufferedReader br;
            switch (status)
            {
                case HttpURLConnection.HTTP_BAD_REQUEST:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
                case HttpURLConnection.HTTP_ACCEPTED:
                    br = new BufferedReader(new InputStreamReader(c.getInputStream()));
                    break;
                default:
                    br = new BufferedReader(new InputStreamReader(c.getErrorStream()));
                    break;
            }
            final StringBuilder sb = new StringBuilder();
            String line;
            while ((line = br.readLine()) != null)
            {
                sb.append(line + "\n");
            }
            br.close();

            final ObjectMapper mapper = new ObjectMapper();
            final T v = mapper.readValue(sb.toString(), responseClass);
            return new ResponseEntity<>(v, HttpStatus.valueOf(status));

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
