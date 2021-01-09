package models

import (
    "time"
)

type Employee struct {
    id int64
    name string
    joiningDate  time.Time
    birthday time.Time
    lastPaySlip time.Time
    totalPaySlip int64
}