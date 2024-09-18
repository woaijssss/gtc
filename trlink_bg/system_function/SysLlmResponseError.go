package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SysLlmResponseErrorFields = struct {
   Id string
   Question string
   Answer string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "question",
    "answer",
    "created_at",
    "modified_at",
    
}

var  SysLlmResponseErrorMeta = &godbx.TableMeta[SysLlmResponseError]{
    Table: "sys_llm_response_error",
    Columns: []string {
        "id",
        "question",
        "answer",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SysLlmResponseError,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "question" == columnName {
            if point {
                 return &ins.Question
            }
            return ins.Question
        }
        if "answer" == columnName {
            if point {
                 return &ins.Answer
            }
            return ins.Answer
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

var SysLlmResponseErrorDao godbx.QuickDao[SysLlmResponseError] = &struct {
	godbx.QuickDao[SysLlmResponseError]
}{
	godbx.NewBaseQuickDao(SysLlmResponseErrorMeta),
}

type SysLlmResponseError struct {
    Id int64
    Question string
    Answer string
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
