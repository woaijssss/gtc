package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserLoginTraceFields = struct {
   Id string
   UserNo string
   NickName string
   LoginType string
   Ip string
   IpRegion string
   UserAgent string
   Address string
   IsSuccess string
   ModifiedAt string
   
}{
    "id",
    "user_no",
    "nick_name",
    "login_type",
    "ip",
    "ip_region",
    "user_agent",
    "address",
    "is_success",
    "modified_at",
    
}

var  TUserLoginTraceMeta = &godbx.TableMeta[TUserLoginTrace]{
    Table: "t_user_login_trace",
    Columns: []string {
        "id",
        "user_no",
        "nick_name",
        "login_type",
        "ip",
        "ip_region",
        "user_agent",
        "address",
        "is_success",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUserLoginTrace,point bool) any {
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
        if "nick_name" == columnName {
            if point {
                 return &ins.NickName
            }
            return ins.NickName
        }
        if "login_type" == columnName {
            if point {
                 return &ins.LoginType
            }
            return ins.LoginType
        }
        if "ip" == columnName {
            if point {
                 return &ins.Ip
            }
            return ins.Ip
        }
        if "ip_region" == columnName {
            if point {
                 return &ins.IpRegion
            }
            return ins.IpRegion
        }
        if "user_agent" == columnName {
            if point {
                 return &ins.UserAgent
            }
            return ins.UserAgent
        }
        if "address" == columnName {
            if point {
                 return &ins.Address
            }
            return ins.Address
        }
        if "is_success" == columnName {
            if point {
                 return &ins.IsSuccess
            }
            return ins.IsSuccess
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

var TUserLoginTraceDao godbx.QuickDao[TUserLoginTrace] = &struct {
	godbx.QuickDao[TUserLoginTrace]
}{
	godbx.NewBaseQuickDao(TUserLoginTraceMeta),
}

type TUserLoginTrace struct {
    Id int64
    UserNo string
    NickName ttypes.NilableString
    LoginType int32
    Ip ttypes.NilableString
    IpRegion ttypes.NilableString
    UserAgent ttypes.NilableString
    Address ttypes.NilableString
    IsSuccess int32
    ModifiedAt ttypes.NormalDatetime
    
}
