package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SysFunctionSystemFeedbackFields = struct {
   Id string
   UserNo string
   Images string
   Content string
   FeedbackType string
   CreatedAt string
   ModifiedAt string
   Status string
   
}{
    "id",
    "user_no",
    "images",
    "content",
    "feedback_type",
    "created_at",
    "modified_at",
    "status",
    
}

var  SysFunctionSystemFeedbackMeta = &godbx.TableMeta[SysFunctionSystemFeedback]{
    Table: "sys_function_system_feedback",
    Columns: []string {
        "id",
        "user_no",
        "images",
        "content",
        "feedback_type",
        "created_at",
        "modified_at",
        "status",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SysFunctionSystemFeedback,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
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
        if "feedback_type" == columnName {
            if point {
                 return &ins.FeedbackType
            }
            return ins.FeedbackType
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

var SysFunctionSystemFeedbackDao godbx.QuickDao[SysFunctionSystemFeedback] = &struct {
	godbx.QuickDao[SysFunctionSystemFeedback]
}{
	godbx.NewBaseQuickDao(SysFunctionSystemFeedbackMeta),
}

type SysFunctionSystemFeedback struct {
    Id int64
    UserNo string
    Images ttypes.NilableString
    Content string
    FeedbackType int32
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    Status int32
    
}
