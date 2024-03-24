package com.codermast.blog.service;


import com.baomidou.mybatisplus.extension.service.IService;
import com.codermast.blog.dto.UserDTO;
import com.codermast.blog.entity.Result;
import com.codermast.blog.entity.User;

public interface UserService extends IService<User> {
    User login(UserDTO userDTO);
}
