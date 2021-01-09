package controllers

import (
    "github.com/revel/revel"
    "github.com/hrmgo/app/models"
)

type Employee struct {
    *revel.Controller
}

type EmployeeCtrl struct {
    GorpController
}

// GET /employee/index
func (c Employee) Index() revel.Result {
    return c.Render()
}

// GET /employee/add
func (c Employee) Add() revel.Result {
    return c.Render()
}

// POST /employee/add
func (c Employee) SaveAdd() revel.Result {
    name := c.Params.Form.Get("name")
    regularizeDate := c.Params.Form.Get("regularize-date")
    birthday := c.Params.Form.Get("birthday")
    
    revel.AppLog.Info("Employee.SaveAdd: %s %s %s", name, regularizeDate, birthday)
    return c.Redirect(Employee.Index)
}