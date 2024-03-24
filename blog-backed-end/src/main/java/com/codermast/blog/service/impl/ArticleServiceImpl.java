package com.codermast.blog.service.impl;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.codermast.blog.entity.Article;
import com.codermast.blog.mapper.ArticleMapper;
import com.codermast.blog.service.ArticleService;
import org.springframework.stereotype.Service;

@Service
public class ArticleServiceImpl  extends ServiceImpl<ArticleMapper, Article> implements ArticleService {
}
