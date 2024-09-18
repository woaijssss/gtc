package community

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TCommunityPostCommentFields = struct {
   Id string
   PostId string
   UserNo string
   Image string
   NickName string
   Context string
   LikeNum string
   ParentId string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "post_id",
    "user_no",
    "image",
    "nick_name",
    "context",
    "like_num",
    "parent_id",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TCommunityPostCommentMeta = &godbx.TableMeta[TCommunityPostComment]{
    Table: "t_community_post_comment",
    Columns: []string {
        "id",
        "post_id",
        "user_no",
        "image",
        "nick_name",
        "context",
        "like_num",
        "parent_id",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TCommunityPostComment,point bool) any {
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
        if "image" == columnName {
            if point {
                 return &ins.Image
            }
            return ins.Image
        }
        if "nick_name" == columnName {
            if point {
                 return &ins.NickName
            }
            return ins.NickName
        }
        if "context" == columnName {
            if point {
                 return &ins.Context
            }
            return ins.Context
        }
        if "like_num" == columnName {
            if point {
                 return &ins.LikeNum
            }
            return ins.LikeNum
        }
        if "parent_id" == columnName {
            if point {
                 return &ins.ParentId
            }
            return ins.ParentId
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

var TCommunityPostCommentDao godbx.QuickDao[TCommunityPostComment] = &struct {
	godbx.QuickDao[TCommunityPostComment]
}{
	godbx.NewBaseQuickDao(TCommunityPostCommentMeta),
}

type TCommunityPostComment struct {
    Id int64
    PostId int64
    UserNo string
    Image ttypes.NilableString
    NickName ttypes.NilableString
    Context ttypes.NilableString
    LikeNum int32
    ParentId int32
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
