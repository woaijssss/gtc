package system_function

import (
    "github.com/woaijssss/godbx"
    
    
)

var SystemConfigFields = struct {
   Id string
   Uuid string
   Env string
   
}{
    "id",
    "uuid",
    "env",
    
}

var  SystemConfigMeta = &godbx.TableMeta[SystemConfig]{
    Table: "system_config",
    Columns: []string {
        "id",
        "uuid",
        "env",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemConfig,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "uuid" == columnName {
            if point {
                 return &ins.Uuid
            }
            return ins.Uuid
        }
        if "env" == columnName {
            if point {
                 return &ins.Env
            }
            return ins.Env
        }
        
        return nil
    },
}

var SystemConfigDao godbx.QuickDao[SystemConfig] = &struct {
	godbx.QuickDao[SystemConfig]
}{
	godbx.NewBaseQuickDao(SystemConfigMeta),
}

type SystemConfig struct {
    Id int64
    Uuid string
    Env int32
    
}
