package com.promise.mongo;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.core.MongoClientFactoryBean;

@Configuration
public class AppConfig
{
    @Bean
    public  MongoClientFactoryBean mongo()
    {
        MongoClientFactoryBean mongo = new MongoClientFactoryBean();
        mongo.setHost("100.101.64.180");
        return mongo;
    }

}
