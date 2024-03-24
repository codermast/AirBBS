package com.codermast.blog.interceptor;


import com.alibaba.fastjson2.JSON;
import com.codermast.blog.context.BaseContext;
import com.codermast.blog.entity.Result;
import com.codermast.blog.properties.JWTProperties;
import com.codermast.blog.utils.JWTUtils;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.JwsHeader;
import io.jsonwebtoken.JwtParser;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;

import java.io.IOException;
import java.io.PrintWriter;

@Slf4j
@Component
public class JWTTokenAdminInterceptor implements HandlerInterceptor {

    @Autowired
    private JWTProperties jwtProperties;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        // 判断当前拦截的是 Controller 方法还是其他资源
        if (!(handler instanceof HandlerMethod)) {
            // 当拦截到的不是动态方法，直接放行
            return true;
        }

        // 1. 从请求头中获取 token
        String token = request.getHeader(jwtProperties.getAdminTokenName());

        // 如果不存在，直接拦截
        if (token == null || token.isEmpty()) {
            setInterceptInfo(response);
            return false;
        }

        // 2. 判断 token 是否正确
        Jws<Claims> claimsJws = JWTUtils.parseClaims(token);

        // Token 解析失败拦截
        if (claimsJws == null){
            setInterceptInfo(response);
            return false;
        }
        // Token 正确
        JwsHeader header = claimsJws.getHeader();

        // TODO: 这里的 uid 最后要抽取出来
        Long uid = (Long) header.get("uid");

        BaseContext.setCurrentId(uid);
        log.info("当前用户id：" + uid);
        return true;
    }

    // 设置拦截返回信息
    public static void setInterceptInfo(HttpServletResponse response) throws IOException {
        // 设置响应信息
        response.setContentType("application/json;charset=UTF-8");
        PrintWriter writer = response.getWriter();
        writer.print(JSON.toJSON(Result.error("未登录！")));
        response.setStatus(401);
    }
}
