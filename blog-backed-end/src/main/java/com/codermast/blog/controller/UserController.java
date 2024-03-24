package com.codermast.blog.controller;

import com.codermast.blog.context.BaseContext;
import com.codermast.blog.dto.UserDTO;
import com.codermast.blog.entity.Result;
import com.codermast.blog.entity.User;
import com.codermast.blog.service.UserService;
import com.codermast.blog.utils.JWTUtils;
import com.codermast.blog.vo.UserLoginVO;
import com.codermast.blog.vo.UserVO;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Slf4j
@CrossOrigin
@RestController
@RequestMapping("/user")
public class UserController {

    @Autowired
    UserService userService;

    // 用户登录
    @PostMapping("/login")
    public Result<UserLoginVO> login(UserDTO userDTO) {

        // 判断登录信息是否合法
        if (userDTO == null){
            return Result.error("请求非法！");
        }

        if (userDTO.getUsername() == null || userDTO.getUsername().isEmpty()){
            return Result.error("用户名为空！");
        }

        if (userDTO.getPassword() == null || userDTO.getPassword().isEmpty()){
            return Result.error("密码为空！");
        }

        // 根据登录信息获取用户对象
        User user = userService.login(userDTO);

        if (user == null) {
            return Result.error("用户名或密码错误！");
        }

        // 此时验证通过

        Map<String,Object> claims = new HashMap<>();
        // TODO:将写死的 uid 抽取出来
        claims.put("uid",user.getUid());
        String token = JWTUtils.createJWT(claims);
        UserLoginVO userLoginVO = UserLoginVO.builder()
                .uid(user.getUid())
                .username(user.getUsername())
                .email(user.getEmail())
                .token(token)
                .build();

        BeanUtils.copyProperties(user,userLoginVO);

        return Result.success(userLoginVO);
    }

    // 用户注册或修改
    @RequestMapping(method = {RequestMethod.POST,RequestMethod.PUT})
    public Result<UserVO> save(@RequestBody UserDTO userDTO) {
        User user = new User();

        BeanUtils.copyProperties(userDTO, user);

        boolean save = userService.saveOrUpdate(user);

        UserVO userVO = new UserVO();
        BeanUtils.copyProperties(user,userVO);

        if (save) {
            return Result.success(userVO);
        } else {
            return Result.error("操作失败！");
        }
    }

    // 退出登录
    @PostMapping("/logout")
    public Result<UserVO> logout() {
        BaseContext.remove();
        return Result.success("退出成功！");
    }

    // 获取所有用户信息
    @GetMapping("/all")
    public Result<List<UserDTO>> all(){
        List<User> list = userService.list();
        List<UserDTO> ret = new ArrayList<>();
        for (User user : list) {
            UserDTO userDTO = new UserDTO();
            BeanUtils.copyProperties(user,userDTO);
            ret.add(userDTO);
        }
        return Result.success(ret);
    }

    // 根据 uid 获取用户信息
    @GetMapping("/{uid}")
    public Result<UserDTO> getByUid(@PathVariable Long uid){
        User userServiceById = userService.getById(uid);

        UserDTO userDTO = new UserDTO();
        BeanUtils.copyProperties(userServiceById,userDTO);
        return Result.success(userDTO);
    }
}
