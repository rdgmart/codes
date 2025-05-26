package br.com.rdgmartins.codes.quarkus.social.dto;

import io.quarkus.runtime.util.StringUtil;

import java.text.SimpleDateFormat;
import java.util.Date;

public class UserCreateRequest {
    private String name;
    private String email;
    private String birthday;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getBirthday() {
        return birthday;
    }

    public void setBirthday(String birthday) {
        this.birthday = birthday;
    }

    public Date getBirthdayDate() {
        try {
            if(!StringUtil.isNullOrEmpty(getBirthday())){
                SimpleDateFormat sdf = new SimpleDateFormat("dd/MM/yyyy");
                return sdf.parse(getBirthday());
            }
        } catch (Exception e) {}
        return null;
    }
}
