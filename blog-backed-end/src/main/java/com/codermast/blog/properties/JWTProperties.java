package com.codermast.blog.properties;


import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Data
@Component
@ConfigurationProperties("blog.jwt")
public class JWTProperties {
    private String adminSecretKey;
    private long adminTtl;
    private String adminTokenName;
}
