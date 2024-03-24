package com.codermast.blog.vo;

import com.baomidou.mybatisplus.annotation.TableId;
import lombok.Data;

@Data
public class UserVO {
    @TableId
    private Long uid;

    private String username;

    private String email;
}
