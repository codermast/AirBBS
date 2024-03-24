package com.codermast.blog.utils;

import com.codermast.blog.properties.JWTProperties;
import io.jsonwebtoken.*;
import io.jsonwebtoken.security.Keys;
import io.jsonwebtoken.security.MacAlgorithm;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.crypto.SecretKey;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

// 关于 JWT 相关的工具类

public class JWTUtils {

    /*
     * 私钥：1. 生成签名的时候使用的秘钥secret，一般可以从本地配置文件中读取，切记这个秘钥不能外露，
     *      2. 只在服务端使用，在任何场景都不应该流露出去。
     *      3. 一旦客户端得知这个secret, 那就意味着客户端是可以自我签发jwt了。
     *      4. 应该大于等于 256位(长度32及以上的字符串)，并且是随机的字符串
     */
    private static final String sign = "sdkjfhsdkjfhalkssdkjfhalksjdhflkasjhdsdkjfhalksjdhflkasjhdjdhflkasjhdalksjdhflkasjhdfklasdsdkjfhalksjdhflkasjhdfklasdsdkjfhalksjdhflkasjhdfklasdsdkjfhalksjdhflkasjhdfklasd";

    // 过期时间，单位毫秒
    private static final long ttlMillis = 1000 * 60 * 60L;

    // 秘钥实例
    private static final SecretKey KEY = Keys.hmacShaKeyFor(sign.getBytes());

    // 创建 token
    public static String createJWT(Map<String, Object> claims) {
        // 指定签名的时候使用的签名算法，也就是header那部分
        MacAlgorithm hs256 = Jwts.SIG.HS256;

        return Jwts.builder()   // 获取 JWT 构建器
                .header()       // 设置头部
                .add("typ", "JWT")  // 添加头部信息
                .add("alg", "HS256")// 添加头部信息
                .and()  // 提交
                .claims(claims) // 添加自定义信息
                .expiration(new Date(System.currentTimeMillis() + ttlMillis))   // 设置过期时间
                .subject("codermast-admin") // 设置主题
                .signWith(KEY, hs256)          // 设置秘钥
                .id(UUID.randomUUID().toString())   // 设置 id
                .issuer("codermast")    // 设置签发者
                .compact(); // 打包
    }

    // 解析 token
    public static Jws<Claims> parseClaims(String token) {
        JwtParser build = Jwts.parser().verifyWith(KEY).build();
        try {
            return build.parseSignedClaims(token);
        } catch (JwtException | IllegalArgumentException e) {
            return null;
        }
    }

    public static void main(String[] args) {
        Map<String, Object> map = new HashMap<>();

        map.put("username", "codermast");
        map.put("uid", 10);
        String token = createJWT(map);
        System.out.println(token);

        Jws<Claims> claimsJws = parseClaims(token);

        System.out.println(parseClaims(token));
    }
}
