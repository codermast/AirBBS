package com.codermast.blog.entity;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;

@Data
@TableName("articles")
public class Article {

    @TableId
    private Long cid;

    // 文章标题
    private String title;

    // 文章内容
    private String content;

    // 作者 ID
    private Long authorId;
}
