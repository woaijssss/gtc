package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SysMapPoiStubFields = struct {
   Id string
   Name string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "name",
    "created_at",
    "modified_at",
    
}

var  SysMapPoiStubMeta = &godbx.TableMeta[SysMapPoiStub]{
    Table: "sys_map_poi_stub",
    Columns: []string {
        "id",
        "name",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SysMapPoiStub,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
        }
        if "created_at" == columnName {
            if point {
                 return &ins.CreatedAt
            }
            return ins.CreatedAt
        }
        if "modified_at" == columnName {
            if point {
                 return &ins.ModifiedAt
            }
            return ins.ModifiedAt
        }
        
        return nil
    },
}

var SysMapPoiStubDao godbx.QuickDao[SysMapPoiStub] = &struct {
	godbx.QuickDao[SysMapPoiStub]
}{
	godbx.NewBaseQuickDao(SysMapPoiStubMeta),
}

type SysMapPoiStub struct {
    Id int64
    Name string
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
