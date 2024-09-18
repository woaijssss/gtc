package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SysFunctionSupplementFields = struct {
   Id string
   Uuid string
   Name string
   UserNo string
   Longitude string
   Latitude string
   Images string
   Content string
   CreatedAt string
   ModifiedAt string
   Status string
   
}{
    "id",
    "uuid",
    "name",
    "user_no",
    "longitude",
    "latitude",
    "images",
    "content",
    "created_at",
    "modified_at",
    "status",
    
}

var  SysFunctionSupplementMeta = &godbx.TableMeta[SysFunctionSupplement]{
    Table: "sys_function_supplement",
    Columns: []string {
        "id",
        "uuid",
        "name",
        "user_no",
        "longitude",
        "latitude",
        "images",
        "content",
        "created_at",
        "modified_at",
        "status",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SysFunctionSupplement,point bool) any {
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
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
        }
        if "user_no" == columnName {
            if point {
                 return &ins.UserNo
            }
            return ins.UserNo
        }
        if "longitude" == columnName {
            if point {
                 return &ins.Longitude
            }
            return ins.Longitude
        }
        if "latitude" == columnName {
            if point {
                 return &ins.Latitude
            }
            return ins.Latitude
        }
        if "images" == columnName {
            if point {
                 return &ins.Images
            }
            return ins.Images
        }
        if "content" == columnName {
            if point {
                 return &ins.Content
            }
            return ins.Content
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
        if "status" == columnName {
            if point {
                 return &ins.Status
            }
            return ins.Status
        }
        
        return nil
    },
}

var SysFunctionSupplementDao godbx.QuickDao[SysFunctionSupplement] = &struct {
	godbx.QuickDao[SysFunctionSupplement]
}{
	godbx.NewBaseQuickDao(SysFunctionSupplementMeta),
}

type SysFunctionSupplement struct {
    Id int64
    Uuid string
    Name string
    UserNo string
    Longitude float64
    Latitude float64
    Images ttypes.NilableString
    Content string
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    Status int32
    
}
