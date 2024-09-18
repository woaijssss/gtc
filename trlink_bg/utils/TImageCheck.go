package utils

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TImageCheckFields = struct {
   Id string
   ImageType string
   TypeId string
   TraceId string
   Status string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "image_type",
    "type_id",
    "trace_id",
    "status",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TImageCheckMeta = &godbx.TableMeta[TImageCheck]{
    Table: "t_image_check",
    Columns: []string {
        "id",
        "image_type",
        "type_id",
        "trace_id",
        "status",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TImageCheck,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "image_type" == columnName {
            if point {
                 return &ins.ImageType
            }
            return ins.ImageType
        }
        if "type_id" == columnName {
            if point {
                 return &ins.TypeId
            }
            return ins.TypeId
        }
        if "trace_id" == columnName {
            if point {
                 return &ins.TraceId
            }
            return ins.TraceId
        }
        if "status" == columnName {
            if point {
                 return &ins.Status
            }
            return ins.Status
        }
        if "creator" == columnName {
            if point {
                 return &ins.Creator
            }
            return ins.Creator
        }
        if "updator" == columnName {
            if point {
                 return &ins.Updator
            }
            return ins.Updator
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

var TImageCheckDao godbx.QuickDao[TImageCheck] = &struct {
	godbx.QuickDao[TImageCheck]
}{
	godbx.NewBaseQuickDao(TImageCheckMeta),
}

type TImageCheck struct {
    Id int64
    ImageType string
    TypeId int64
    TraceId string
    Status int32
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
