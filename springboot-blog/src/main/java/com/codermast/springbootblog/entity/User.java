package com.codermast.springbootblog.entity;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;

@Data
@TableName("users")
public class User {

    @TableId
    private Long uid;

    private String username;

    private String password;

    private String mail;
}
