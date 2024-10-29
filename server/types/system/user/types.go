// Code generated by jzero. DO NOT EDIT.
package systemuser

import (
	"time"
)

var (
	_ = time.Now()
)

type AddRequest struct {
	Username   string `json:"username"`
	UserGender string `json:"userGender,default=1"`
	NickName   string `json:"nickName"`
	UserPhone  string `json:"userPhone"`
	UserEmail  string `json:"userEmail"`
	Status     string `json:"status,default=1"`
	Password   string `json:"password"`
}

type AddResponse struct {
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids"`
}

type DeleteResponse struct {
}

type EditRequest struct {
	Id         uint64 `json:"id"`
	Username   string `json:"username"`
	UserGender string `json:"userGender"`
	NickName   string `json:"nickName"`
	UserPhone  string `json:"userPhone"`
	UserEmail  string `json:"userEmail"`
	Status     string `json:"status"`
}

type EditResponse struct {
}

type ListRequest struct {
	PageRequest
	Username   string `form:"username,optional"`
	UserGender string `form:"userGender,optional"`
	NickName   string `form:"nickName,optional"`
	UserPhone  string `form:"userPhone,optional"`
	UserEmail  string `form:"userEmail,optional"`
	Status     string `form:"status,optional"`
}

type ListResponse struct {
	PageResponse
	Records []SystemUser `json:"records"`
}

type PageRequest struct {
	Current int `form:"current"`
	Size    int `form:"size"`
}

type PageResponse struct {
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
}

type SystemUser struct {
	Id         uint64   `json:"id"`
	Username   string   `json:"username"`
	UserGender string   `json:"userGender"`
	NickName   string   `json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	UserRoles  []string `json:"userRoles"`
	Status     string   `json:"status"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}
