package community

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TCommunityPostCommentLikeFields = struct {
   Id string
   UserNo string
   CommentId string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "user_no",
    "comment_id",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TCommunityPostCommentLikeMeta = &godbx.TableMeta[TCommunityPostCommentLike]{
    Table: "t_community_post_comment_like",
    Columns: []string {
        "id",
        "user_no",
        "comment_id",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TCommunityPostCommentLike,point bool) any {
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
        if "comment_id" == columnName {
            if point {
                 return &ins.CommentId
            }
            return ins.CommentId
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

var TCommunityPostCommentLikeDao godbx.QuickDao[TCommunityPostCommentLike] = &struct {
	godbx.QuickDao[TCommunityPostCommentLike]
}{
	godbx.NewBaseQuickDao(TCommunityPostCommentLikeMeta),
}

type TCommunityPostCommentLike struct {
    Id int64
    UserNo string
    CommentId string
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
