package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserAuthorizationFields = struct {
   Id string
   UserNo string
   IdentityType string
   Identifier string
   Credential string
   IsDel string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "user_no",
    "identity_type",
    "identifier",
    "credential",
    "is_del",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TUserAuthorizationMeta = &godbx.TableMeta[TUserAuthorization]{
    Table: "t_user_authorization",
    Columns: []string {
        "id",
        "user_no",
        "identity_type",
        "identifier",
        "credential",
        "is_del",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUserAuthorization,point bool) any {
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
        if "identity_type" == columnName {
            if point {
                 return &ins.IdentityType
            }
            return ins.IdentityType
        }
        if "identifier" == columnName {
            if point {
                 return &ins.Identifier
            }
            return ins.Identifier
        }
        if "credential" == columnName {
            if point {
                 return &ins.Credential
            }
            return ins.Credential
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
        
        return nil
    },
}

var TUserAuthorizationDao godbx.QuickDao[TUserAuthorization] = &struct {
	godbx.QuickDao[TUserAuthorization]
}{
	godbx.NewBaseQuickDao(TUserAuthorizationMeta),
}

type TUserAuthorization struct {
    Id int64
    UserNo string
    IdentityType int32
    Identifier string
    Credential ttypes.NilableString
    IsDel int8
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
