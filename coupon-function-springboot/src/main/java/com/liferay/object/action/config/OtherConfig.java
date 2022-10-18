package com.liferay.object.action.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder;

import com.fasterxml.jackson.databind.ObjectMapper;

@Configuration
public class OtherConfig {

    @Bean
    public ObjectMapper customJson() {
        return new Jackson2ObjectMapperBuilder()
            .indentOutput(true)
            .build();
    }

}