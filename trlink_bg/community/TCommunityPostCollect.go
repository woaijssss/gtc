package community

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TCommunityPostCollectFields = struct {
   Id string
   PostId string
   UserNo string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "post_id",
    "user_no",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TCommunityPostCollectMeta = &godbx.TableMeta[TCommunityPostCollect]{
    Table: "t_community_post_collect",
    Columns: []string {
        "id",
        "post_id",
        "user_no",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TCommunityPostCollect,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "post_id" == columnName {
            if point {
                 return &ins.PostId
            }
            return ins.PostId
        }
        if "user_no" == columnName {
            if point {
                 return &ins.UserNo
            }
            return ins.UserNo
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

var TCommunityPostCollectDao godbx.QuickDao[TCommunityPostCollect] = &struct {
	godbx.QuickDao[TCommunityPostCollect]
}{
	godbx.NewBaseQuickDao(TCommunityPostCollectMeta),
}

type TCommunityPostCollect struct {
    Id int64
    PostId int64
    UserNo string
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
