package main

import (
	"net/url"
	"time"
)

type User struct {
	ID          *int64       `json:"id,string,omitempty"`
	Username    *string      `json:"username,omitempty"`
	Email       *string      `json:"email,omitempty"`
	Password    *string      `json:"password,omitempty"`
	PassConfirm *string      `json:"passconfirm,omitempty"`
	PassHash    *string      `json:"passhash,omitempty"`
	RoleID      *int         `json:"role_id,string,omitempty"`
	Firstname   *string      `json:"firstname,omitempty"`
	Lastname    *string      `json:"lastname,omitempty"`
	Company     *Company     `json:"company,omitempty"`
	Parent      *User        `json:"parent,omitempty"`
	Children    *[]*User     `json:"children,omitempty"`
	JobPosition *JobPosition `json:"job_position,omitempty"`
	Sex         *string      `json:"sex,omitempty"`
	Birthdate   *time.Time   `json:"birthdate,omitempty"`
	Active      *bool        `json:"active,omitempty"`
}

type Company struct {
	ID        *int64   `json:"id,string,omitempty"`
	Title     *string  `json:"title,omitempty"`
	Parent    *Company `json:"parent,omitempty"`
	Contact   *User    `json:"contact,omitempty"`
	Domain    *url.URL `json:"domain,omitempty"`
	Published *bool    `json:"published,omitempty"`
}

type JobPosition struct {
	ID        *int64  `json:"id,string,omitempty"`
	Title     *string `json:"title,omitempty"`
	Published *bool   `json:"published,omitempty"`
}
