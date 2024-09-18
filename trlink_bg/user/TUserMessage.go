package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserMessageFields = struct {
   Id string
   UserNo string
   Content string
   ReadStatus string
   IsDel string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   Title string
   MsgType string
   
}{
    "id",
    "user_no",
    "content",
    "read_status",
    "is_del",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    "title",
    "msg_type",
    
}

var  TUserMessageMeta = &godbx.TableMeta[TUserMessage]{
    Table: "t_user_message",
    Columns: []string {
        "id",
        "user_no",
        "content",
        "read_status",
        "is_del",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        "title",
        "msg_type",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUserMessage,point bool) any {
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
        if "content" == columnName {
            if point {
                 return &ins.Content
            }
            return ins.Content
        }
        if "read_status" == columnName {
            if point {
                 return &ins.ReadStatus
            }
            return ins.ReadStatus
        }
        if "is_del" == columnName {
            if point {
                 return &ins.IsDel
            }
            return ins.IsDel
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
        if "title" == columnName {
            if point {
                 return &ins.Title
            }
            return ins.Title
        }
        if "msg_type" == columnName {
            if point {
                 return &ins.MsgType
            }
            return ins.MsgType
        }
        
        return nil
    },
}

var TUserMessageDao godbx.QuickDao[TUserMessage] = &struct {
	godbx.QuickDao[TUserMessage]
}{
	godbx.NewBaseQuickDao(TUserMessageMeta),
}

type TUserMessage struct {
    Id int64
    UserNo string
    Content ttypes.NilableString
    ReadStatus int32
    IsDel int32
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    Title ttypes.NilableString
    MsgType int32
    
}
