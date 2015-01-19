package models

import (
    "github.com/go-xorm/xorm"
    "log"
)

var engine *xorm.Engine

func init() {
    var err error
    if engine, err = xorm.NewEngine(driverName, *sourceName); err != nil {
        log.Fatal(err)
    }
    if err := engine.Sync(File{}, Tag{}, Dataset{}); err != nil {
        log.Fatal(err)
    }
}

func GetEngine() *xorm.Engine {
    return engine
}
