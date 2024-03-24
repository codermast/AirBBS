package com.codermast.blog.context;

public class BaseContext {

   private static final ThreadLocal<Long> threadLocal = new ThreadLocal<>();

    // 设置当前线程中的用户 id
    public static void setCurrentId(Long id){
        threadLocal.set(id);
    }

    // 获取当前线程中的用户 id
    public static Long getCurrentId(){
        return threadLocal.get();
    }

    // 移除当前线程中的用户 id
    public static void remove(){
        threadLocal.remove();
    }
}
