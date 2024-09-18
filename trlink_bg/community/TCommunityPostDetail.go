package community

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TCommunityPostDetailFields = struct {
   Id string
   PostId string
   Images string
   Video string
   Content string
   SrcPlacename string
   DestPlacename string
   StartTime string
   Crowd string
   PeopleCount string
   Budget string
   Gender string
   BudgetType string
   
}{
    "id",
    "post_id",
    "images",
    "video",
    "content",
    "src_placeName",
    "dest_placeName",
    "start_time",
    "crowd",
    "people_count",
    "budget",
    "gender",
    "budget_type",
    
}

var  TCommunityPostDetailMeta = &godbx.TableMeta[TCommunityPostDetail]{
    Table: "t_community_post_detail",
    Columns: []string {
        "id",
        "post_id",
        "images",
        "video",
        "content",
        "src_placeName",
        "dest_placeName",
        "start_time",
        "crowd",
        "people_count",
        "budget",
        "gender",
        "budget_type",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TCommunityPostDetail,point bool) any {
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
        if "images" == columnName {
            if point {
                 return &ins.Images
            }
            return ins.Images
        }
        if "video" == columnName {
            if point {
                 return &ins.Video
            }
            return ins.Video
        }
        if "content" == columnName {
            if point {
                 return &ins.Content
            }
            return ins.Content
        }
        if "src_placeName" == columnName {
            if point {
                 return &ins.SrcPlacename
            }
            return ins.SrcPlacename
        }
        if "dest_placeName" == columnName {
            if point {
                 return &ins.DestPlacename
            }
            return ins.DestPlacename
        }
        if "start_time" == columnName {
            if point {
                 return &ins.StartTime
            }
            return ins.StartTime
        }
        if "crowd" == columnName {
            if point {
                 return &ins.Crowd
            }
            return ins.Crowd
        }
        if "people_count" == columnName {
            if point {
                 return &ins.PeopleCount
            }
            return ins.PeopleCount
        }
        if "budget" == columnName {
            if point {
                 return &ins.Budget
            }
            return ins.Budget
        }
        if "gender" == columnName {
            if point {
                 return &ins.Gender
            }
            return ins.Gender
        }
        if "budget_type" == columnName {
            if point {
                 return &ins.BudgetType
            }
            return ins.BudgetType
        }
        
        return nil
    },
}

var TCommunityPostDetailDao godbx.QuickDao[TCommunityPostDetail] = &struct {
	godbx.QuickDao[TCommunityPostDetail]
}{
	godbx.NewBaseQuickDao(TCommunityPostDetailMeta),
}

type TCommunityPostDetail struct {
    Id int64
    PostId int64
    Images ttypes.NilableString
    Video ttypes.NilableString
    Content ttypes.NilableString
    SrcPlacename string
    DestPlacename string
    StartTime string
    Crowd string
    PeopleCount int32
    Budget int32
    Gender int32
    BudgetType int32
    
}
