package com.codermast.blog.controller;

import com.codermast.blog.entity.Article;
import com.codermast.blog.entity.Result;
import com.codermast.blog.service.ArticleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/article")
public class ArticleController {

    @Autowired
    private ArticleService articleService;

    // 文章创建
    @PostMapping
    public Result<Article> save(@RequestBody Article article) {
        boolean save = articleService.save(article);

        if (save) {
            return Result.success(article);
        } else {
            return Result.error("添加失败");
        }
    }

    // 文章修改
    @PutMapping
    public Result<Article> modify(@RequestBody Article article) {
        boolean update = articleService.updateById(article);

        if (update) {
            return Result.success(article);
        } else {
            return Result.error("修改失败");
        }
    }

    // 文章删除
    @DeleteMapping
    public Result<Article> delete(@RequestBody Article article) {
        boolean remove = articleService.removeById(article);

        if (remove) {
            return Result.success("删除成功");
        } else {
            return Result.error("删除失败");
        }
    }

    // 根据文章id查文章
    @GetMapping("/{cid}")
    public Result<Article> getByCid(@PathVariable Long cid) {
        Article article = articleService.getById(cid);

        if (article != null){
            return Result.success(article);
        }else {
            return Result.error("文章不存在");
        }
    }

    // 查询所有文章
    @GetMapping("/all")
    public Result<List<Article>> getAll(){
        List<Article> list = articleService.list();
        return Result.success(list);
    }

    //// 根据关键字查文章
    //@GetMapping("/")
}
