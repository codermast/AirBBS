package com.codermast.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.codermast.blog.dto.UserDTO;
import com.codermast.blog.entity.User;
import com.codermast.blog.mapper.UserMapper;
import com.codermast.blog.service.UserService;
import org.springframework.stereotype.Service;
import org.springframework.util.DigestUtils;

@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {


    @Override
    public User login(UserDTO userDTO) {
        // 控制层已经对 userDTO 是否为 null 已经进行校验
        LambdaQueryWrapper<User> queryWrapper = new LambdaQueryWrapper<>();

        queryWrapper.eq(User::getUsername, userDTO.getUsername());
        //queryWrapper.eq(User::getPassword, DigestUtils.md5Digest(userDTO.getPassword().getBytes()));
        queryWrapper.eq(User::getPassword, userDTO.getPassword());

        return this.getOne(queryWrapper);
    }
}
