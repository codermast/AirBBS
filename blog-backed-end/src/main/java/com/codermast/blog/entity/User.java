package com.codermast.blog.entity;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;

@TableName("users")
public class User {

    @TableId
    Long uid;

    String username;

    String password;

    String email;
}
