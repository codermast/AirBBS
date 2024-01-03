package com.codermast.springbootblog.service.impl;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.codermast.springbootblog.entity.User;
import com.codermast.springbootblog.mapper.UserMapper;
import com.codermast.springbootblog.service.UserService;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {
}
