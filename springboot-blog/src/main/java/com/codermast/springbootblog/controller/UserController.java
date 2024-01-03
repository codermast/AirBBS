package com.codermast.springbootblog.controller;

import com.codermast.springbootblog.entity.R;
import com.codermast.springbootblog.entity.User;
import com.codermast.springbootblog.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/user")
public class UserController {

    @Autowired
    private UserService userService;

    // 用户查询
    @GetMapping("/{uid}")
    public R<User> getByUid(@PathVariable Long uid) {
        User user = userService.getById(uid);

        if (user == null) {
            return R.error("用户名不存在");
        }
        return R.success(user);
    }

    // 用户注册
    @PostMapping
    public R<User> save(@RequestBody User user) {
        boolean save = userService.save(user);
        if (save) {
            return R.success("注册成功！");
        } else {
            return R.error("注册失败！");
        }
    }

    // 用户删除
    @DeleteMapping
    public R<User> delete(@RequestBody User user) {
        boolean remove = userService.removeById(user);
        if (remove) {
            return R.success("删除成功！");
        } else {
            return R.error("删除失败！");
        }
    }

    // 用户修改
    @PutMapping
    public R<User> update(@RequestBody User user) {
        boolean save = userService.save(user);

        if (save) {
            return R.success(user);
        } else {
            return R.error("修改失败！");
        }
    }
}
