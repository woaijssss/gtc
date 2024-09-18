package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserFocusOnFields = struct {
   Id string
   UserNo string
   FocusedUserNo string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "user_no",
    "focused_user_no",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TUserFocusOnMeta = &godbx.TableMeta[TUserFocusOn]{
    Table: "t_user_focus_on",
    Columns: []string {
        "id",
        "user_no",
        "focused_user_no",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUserFocusOn,point bool) any {
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
        if "focused_user_no" == columnName {
            if point {
                 return &ins.FocusedUserNo
            }
            return ins.FocusedUserNo
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

var TUserFocusOnDao godbx.QuickDao[TUserFocusOn] = &struct {
	godbx.QuickDao[TUserFocusOn]
}{
	godbx.NewBaseQuickDao(TUserFocusOnMeta),
}

type TUserFocusOn struct {
    Id int64
    UserNo string
    FocusedUserNo string
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
