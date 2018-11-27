package com.promise.mongo;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.WriteResultChecking;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;

import com.mongodb.MongoClient;
import com.promise.mongo.model.Person;

@SpringBootApplication
public class MongoApp implements CommandLineRunner
{

    private static final Log log = LogFactory.getLog(MongoApp.class);

    @Autowired
    private MongoClient client;  

    public static void main(String[] args)
            throws Exception
    {
        SpringApplication.run(MongoApp.class, args);
    }

    @Override
    public void run(String... args)
            throws Exception
    {
        
        MongoTemplate t = new MongoTemplate(client, "database");
        t.setWriteResultChecking(WriteResultChecking.EXCEPTION);
        
        MongoOperations mongoOps = t;
        mongoOps.insert(new Person("Joe", 34));

        log.info(mongoOps.findOne(new Query(Criteria.where("name").is("Joe")), Person.class));

        mongoOps.dropCollection("person");
    }
}
