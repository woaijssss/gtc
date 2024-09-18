package system_data

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SystemDataRegionLevelFields = struct {
   Id string
   Uuid string
   ParentId string
   Name string
   ShortName string
   Code string
   Citycode string
   Level string
   IsDirect string
   Image string
   
}{
    "id",
    "uuid",
    "parent_id",
    "name",
    "short_name",
    "code",
    "citycode",
    "level",
    "is_direct",
    "image",
    
}

var  SystemDataRegionLevelMeta = &godbx.TableMeta[SystemDataRegionLevel]{
    Table: "system_data_region_level",
    Columns: []string {
        "id",
        "uuid",
        "parent_id",
        "name",
        "short_name",
        "code",
        "citycode",
        "level",
        "is_direct",
        "image",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataRegionLevel,point bool) any {
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
        if "parent_id" == columnName {
            if point {
                 return &ins.ParentId
            }
            return ins.ParentId
        }
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
        }
        if "short_name" == columnName {
            if point {
                 return &ins.ShortName
            }
            return ins.ShortName
        }
        if "code" == columnName {
            if point {
                 return &ins.Code
            }
            return ins.Code
        }
        if "citycode" == columnName {
            if point {
                 return &ins.Citycode
            }
            return ins.Citycode
        }
        if "level" == columnName {
            if point {
                 return &ins.Level
            }
            return ins.Level
        }
        if "is_direct" == columnName {
            if point {
                 return &ins.IsDirect
            }
            return ins.IsDirect
        }
        if "image" == columnName {
            if point {
                 return &ins.Image
            }
            return ins.Image
        }
        
        return nil
    },
}

var SystemDataRegionLevelDao godbx.QuickDao[SystemDataRegionLevel] = &struct {
	godbx.QuickDao[SystemDataRegionLevel]
}{
	godbx.NewBaseQuickDao(SystemDataRegionLevelMeta),
}

type SystemDataRegionLevel struct {
    Id int64
    Uuid int64
    ParentId int64
    Name string
    ShortName string
    Code int64
    Citycode int64
    Level int32
    IsDirect int32
    Image ttypes.NilableString
    
}
