package com.codermast.blog.controller;

import com.codermast.blog.entity.Result;
import com.codermast.blog.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Slf4j
@RequestMapping("/user")
public class UserController {

    @PostMapping("/login")
    public Result login(User user){

    }
}
