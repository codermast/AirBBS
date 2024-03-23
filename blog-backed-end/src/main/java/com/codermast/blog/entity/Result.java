package com.codermast.blog.entity;

public class Result {
    // 状态码
    String code;

    // 数据
    Object data;

    // 信息
    String message;

    // 返回失败信息
    public Result error(String message){
        Result result = new Result();

        result.code = "0";
        result.message = message;
        result.data = null;

        return result;
    }

    // 成功
    public Result success(String message){
        return this.success(message,null);
    }

    // 成功
    public Result success(String message,Object data){
        Result result = new Result();

        result.code = "1";
        result.message = message;
        result.data = data;
        return result;
    }
}
