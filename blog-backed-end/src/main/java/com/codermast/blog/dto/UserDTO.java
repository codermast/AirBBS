package com.codermast.blog.dto;

import lombok.Data;

import java.io.Serializable;

@Data
public class UserDTO implements Serializable {

    Long uid;

    String username;

    String password;

    String email;
}
