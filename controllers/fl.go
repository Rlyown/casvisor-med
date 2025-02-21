// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"encoding/json"

	"github.com/beego/beego/utils/pagination"
	"github.com/casvisor/casvisor/object"
	"github.com/casvisor/casvisor/util"
)

// GetFLs
// @Title GetFLs
// @Tag FL API
// @Description get all FLs
// @Param   pageSize     query    string  true        "The size of each page"
// @Param   p     query    string  true        "The number of the page"
// @Success 200 {object} object.FL The Response object
// @router /get-fls [get]
func (c *ApiController) GetFLs() {
	owner := c.Input().Get("owner")
	limit := c.Input().Get("pageSize")
	page := c.Input().Get("p")
	field := c.Input().Get("field")
	value := c.Input().Get("value")
	sortField := c.Input().Get("sortField")
	sortOrder := c.Input().Get("sortOrder")

	if limit == "" || page == "" {
		fls, err := object.GetFederalLearnings()
		if err != nil {
			c.ResponseError(err.Error())
			return
		}
		c.ResponseOk(fls)
	} else {
		limit := util.ParseInt(limit)
		count, err := object.GetFederalLearningCount(owner, field, value)
		if err != nil {
			c.ResponseError(err.Error())
			return
		}
		paginator := pagination.SetPaginator(c.Ctx, limit, count)
		fls, err := object.GetPaginationFederalLearning(owner, paginator.Offset(), limit, field, value, sortField, sortOrder)
		if err != nil {
			c.ResponseError(err.Error())
			return
		}
		c.ResponseOk(fls, paginator.Nums())
	}
}

// GetFL
// @Title GetFL
// @Tag FL API
// @Description get FL
// @Param   id     query    string  true        "The id of the FL"
// @Success 200 {object} object.FL The Response object
// @router /get-fl [get]
func (c *ApiController) GetFL() {
	id := c.Input().Get("id")

	fl, err := object.GetFederalLearning(id)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.ResponseOk(fl)
}

// UpdateFL
// @Title UpdateFL
// @Tag FL API
// @Description update FL
// @Param   id     query    string  true        "The id of the FL"
// @Param   body    body   object.FL  true        "The details of the FL"
// @Success 200 {object} controllers.Response The Response object
// @router /update-fl [post]
func (c *ApiController) UpdateFL() {
	id := c.Input().Get("id")

	var fl object.FederalLearning
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &fl)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = wrapActionResponse(object.UpdateFederalLearning(id, &fl))
	c.ServeJSON()
}

// AddFL
// @Title AddFL
// @Tag FL API
// @Description add a FL
// @Param   body    body   object.FL  true        "The details of the FL"
// @Success 200 {object} controllers.Response The Response object
// @router /add-fl [post]
func (c *ApiController) AddFL() {
	var fl object.FederalLearning
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &fl)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = wrapActionResponse(object.AddFederalLearning(&fl))
	c.ServeJSON()
}

// DeleteFL
// @Title DeleteFL
// @Tag FL API
// @Description delete a FL
// @Param   body    body   object.FL  true        "The details of the FL"
// @Success 200 {object} controllers.Response The Response object
// @router /delete-fl [post]
func (c *ApiController) DeleteFL() {
	var fl object.FederalLearning
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &fl)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	c.Data["json"] = wrapActionResponse(object.DeleteFederalLearning(&fl))
	c.ServeJSON()
}
