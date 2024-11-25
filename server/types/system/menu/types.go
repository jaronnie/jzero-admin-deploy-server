// Code generated by jzero. DO NOT EDIT.
package systemmenu

import (
	"time"
)

var (
	_ = time.Now()
)

type AddRequest struct {
	ActiveMenu      string   `json:"activeMenu,optional"`
	MenuType        string   `json:"menuType"`
	MenuName        string   `json:"menuName"`
	RouteName       string   `json:"routeName"`
	RoutePath       string   `json:"routePath"`
	Component       string   `json:"component"`
	Icon            string   `json:"icon"`
	IconType        string   `json:"iconType"`
	ParentId        uint64   `json:"parentId"`
	Status          string   `json:"status"`
	KeepAlive       bool     `json:"keepAlive"`
	Constant        bool     `json:"constant"`
	Order           uint64   `json:"order"`
	HideInMenu      bool     `json:"hideInMenu"`
	Href            string   `json:"href,optional"`
	MultiTab        bool     `json:"multiTab"`
	FixedIndexInTab bool     `json:"fixedIndexInTab,optional"`
	Query           []Query  `json:"query"`
	Buttons         []Button `json:"buttons"`
	I18nKey         string   `json:"i18nKey"`
}

type AddResponse struct {
}

type Button struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids"`
}

type DeleteResponse struct {
}

type EditRequest struct {
	Id              uint64   `json:"id"`
	ActiveMenu      string   `json:"activeMenu,optional"`
	MenuType        string   `json:"menuType"`
	MenuName        string   `json:"menuName"`
	RouteName       string   `json:"routeName"`
	RoutePath       string   `json:"routePath"`
	Component       string   `json:"component"`
	Icon            string   `json:"icon"`
	IconType        string   `json:"iconType"`
	ParentId        uint64   `json:"parentId"`
	Status          string   `json:"status"`
	KeepAlive       bool     `json:"keepAlive"`
	Constant        bool     `json:"constant"`
	Order           uint64   `json:"order"`
	HideInMenu      bool     `json:"hideInMenu"`
	Href            string   `json:"href,optional"`
	MutiTab         bool     `json:"mutiTab"`
	FixedIndexInTab bool     `json:"fixedIndexInTab,optional"`
	Query           []Query  `json:"query"`
	Buttons         []Button `json:"buttons"`
	I18nKey         string   `json:"i18nKey"`
}

type EditResponse struct {
}

type GetAllPagesRequest struct {
}

type ListRequest struct {
	PageRequest
}

type ListResponse struct {
	PageResponse
	Records []SystemMenu `json:"records"`
}

type PageRequest struct {
	Current int `form:"current,default=1,optional"`
	Size    int `form:"size,default=10,optional"`
}

type PageResponse struct {
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
}

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SystemMenu struct {
	Id         uint64       `json:"id"`
	ParentId   uint64       `json:"parentId"`
	MenuType   string       `json:"menuType"`
	MenuName   string       `json:"menuName"`
	RouteName  string       `json:"routeName"`
	RoutePath  string       `json:"routePath"`
	Component  string       `json:"component"`
	Icon       string       `json:"icon"`
	IconType   string       `json:"iconType"`
	Buttons    []Button     `json:"buttons"`
	Order      uint64       `json:"order"`
	I18nKey    string       `json:"i18nKey"`
	Status     string       `json:"status"`
	KeepAlive  bool         `json:"keepAlive"`
	Constant   bool         `json:"constant"`
	HideInMenu bool         `json:"hideInMenu"`
	MultiTab   bool         `json:"mutiTab"`
	ActiveMenu string       `json:"activeMenu"`
	Children   []SystemMenu `json:"children"`
}

type TreeRequest struct {
}

type TreeResponse struct {
	Id       uint64         `json:"id"`
	Label    string         `json:"label"`
	PId      uint64         `json:"pId"`
	Order    uint64         `json:"-"`
	Children []TreeResponse `json:"children"`
}
