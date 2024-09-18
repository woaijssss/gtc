package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserScoreFields = struct {
   Id string
   UserNo string
   Uri string
   Description string
   Score string
   CreatedAt string
   ModifiedAt string
   ScoreType string
   TypeId string
   
}{
    "id",
    "user_no",
    "uri",
    "description",
    "score",
    "created_at",
    "modified_at",
    "score_type",
    "type_id",
    
}

var  TUserScoreMeta = &godbx.TableMeta[TUserScore]{
    Table: "t_user_score",
    Columns: []string {
        "id",
        "user_no",
        "uri",
        "description",
        "score",
        "created_at",
        "modified_at",
        "score_type",
        "type_id",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUserScore,point bool) any {
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
        if "uri" == columnName {
            if point {
                 return &ins.Uri
            }
            return ins.Uri
        }
        if "description" == columnName {
            if point {
                 return &ins.Description
            }
            return ins.Description
        }
        if "score" == columnName {
            if point {
                 return &ins.Score
            }
            return ins.Score
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
        if "score_type" == columnName {
            if point {
                 return &ins.ScoreType
            }
            return ins.ScoreType
        }
        if "type_id" == columnName {
            if point {
                 return &ins.TypeId
            }
            return ins.TypeId
        }
        
        return nil
    },
}

var TUserScoreDao godbx.QuickDao[TUserScore] = &struct {
	godbx.QuickDao[TUserScore]
}{
	godbx.NewBaseQuickDao(TUserScoreMeta),
}

type TUserScore struct {
    Id int64
    UserNo string
    Uri string
    Description string
    Score int32
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    ScoreType int32
    TypeId int64
    
}
