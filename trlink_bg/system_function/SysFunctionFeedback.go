package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SysFunctionFeedbackFields = struct {
   Id string
   PoiId string
   UserNo string
   Images string
   Content string
   CreatedAt string
   ModifiedAt string
   Status string
   
}{
    "id",
    "poi_id",
    "user_no",
    "images",
    "content",
    "created_at",
    "modified_at",
    "status",
    
}

var  SysFunctionFeedbackMeta = &godbx.TableMeta[SysFunctionFeedback]{
    Table: "sys_function_feedback",
    Columns: []string {
        "id",
        "poi_id",
        "user_no",
        "images",
        "content",
        "created_at",
        "modified_at",
        "status",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SysFunctionFeedback,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "poi_id" == columnName {
            if point {
                 return &ins.PoiId
            }
            return ins.PoiId
        }
        if "user_no" == columnName {
            if point {
                 return &ins.UserNo
            }
            return ins.UserNo
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

var SysFunctionFeedbackDao godbx.QuickDao[SysFunctionFeedback] = &struct {
	godbx.QuickDao[SysFunctionFeedback]
}{
	godbx.NewBaseQuickDao(SysFunctionFeedbackMeta),
}

type SysFunctionFeedback struct {
    Id int64
    PoiId int64
    UserNo string
    Images ttypes.NilableString
    Content string
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    Status int32
    
}
