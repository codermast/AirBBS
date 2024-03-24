package com.codermast.blog.entity;

import lombok.Data;

@Data
public class Result <T> {
    // 状态码
    String code;

    // 数据
    T data;

    // 信息
    String message;

    // 返回失败信息
    public static <T> Result<T> error(String message){
        Result<T> result = new Result<T>();

        result.code = "0";
        result.message = message;
        result.data = null;

        return result;
    }

    // 成功
    public static <T> Result<T> success(String message){
        return success(message,null);
    }

    // 成功
    public static <T> Result<T> success(T data){
        return success(null,data);
    }

    // 成功
    public static <T> Result<T> success(String message,T data){
        Result<T> result = new Result<T>();

        result.code = "1";
        result.message = message;
        result.data = data;
        return result;
    }
}
