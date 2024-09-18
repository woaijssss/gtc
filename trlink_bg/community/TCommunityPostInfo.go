package community

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TCommunityPostInfoFields = struct {
   Id string
   UserNo string
   TripId string
   PoiUuid string
   LocationText string
   Location string
   Province string
   ProvinceUuid string
   CityUuid string
   AdUuid string
   Title string
   Image string
   PostType string
   IsValid string
   CreatedAt string
   ModifiedAt string
   LikeNum string
   CollectNum string
   CommentNum string
   Status string
   AuditDes string
   FromType string
   
}{
    "id",
    "user_no",
    "trip_id",
    "poi_uuid",
    "location_text",
    "location",
    "province",
    "province_uuid",
    "city_uuid",
    "ad_uuid",
    "title",
    "image",
    "post_type",
    "is_valid",
    "created_at",
    "modified_at",
    "like_num",
    "collect_num",
    "comment_num",
    "status",
    "audit_des",
    "from_type",
    
}

var  TCommunityPostInfoMeta = &godbx.TableMeta[TCommunityPostInfo]{
    Table: "t_community_post_info",
    Columns: []string {
        "id",
        "user_no",
        "trip_id",
        "poi_uuid",
        "location_text",
        "location",
        "province",
        "province_uuid",
        "city_uuid",
        "ad_uuid",
        "title",
        "image",
        "post_type",
        "is_valid",
        "created_at",
        "modified_at",
        "like_num",
        "collect_num",
        "comment_num",
        "status",
        "audit_des",
        "from_type",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TCommunityPostInfo,point bool) any {
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
        if "trip_id" == columnName {
            if point {
                 return &ins.TripId
            }
            return ins.TripId
        }
        if "poi_uuid" == columnName {
            if point {
                 return &ins.PoiUuid
            }
            return ins.PoiUuid
        }
        if "location_text" == columnName {
            if point {
                 return &ins.LocationText
            }
            return ins.LocationText
        }
        if "location" == columnName {
            if point {
                 return &ins.Location
            }
            return ins.Location
        }
        if "province" == columnName {
            if point {
                 return &ins.Province
            }
            return ins.Province
        }
        if "province_uuid" == columnName {
            if point {
                 return &ins.ProvinceUuid
            }
            return ins.ProvinceUuid
        }
        if "city_uuid" == columnName {
            if point {
                 return &ins.CityUuid
            }
            return ins.CityUuid
        }
        if "ad_uuid" == columnName {
            if point {
                 return &ins.AdUuid
            }
            return ins.AdUuid
        }
        if "title" == columnName {
            if point {
                 return &ins.Title
            }
            return ins.Title
        }
        if "image" == columnName {
            if point {
                 return &ins.Image
            }
            return ins.Image
        }
        if "post_type" == columnName {
            if point {
                 return &ins.PostType
            }
            return ins.PostType
        }
        if "is_valid" == columnName {
            if point {
                 return &ins.IsValid
            }
            return ins.IsValid
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
        if "like_num" == columnName {
            if point {
                 return &ins.LikeNum
            }
            return ins.LikeNum
        }
        if "collect_num" == columnName {
            if point {
                 return &ins.CollectNum
            }
            return ins.CollectNum
        }
        if "comment_num" == columnName {
            if point {
                 return &ins.CommentNum
            }
            return ins.CommentNum
        }
        if "status" == columnName {
            if point {
                 return &ins.Status
            }
            return ins.Status
        }
        if "audit_des" == columnName {
            if point {
                 return &ins.AuditDes
            }
            return ins.AuditDes
        }
        if "from_type" == columnName {
            if point {
                 return &ins.FromType
            }
            return ins.FromType
        }
        
        return nil
    },
}

var TCommunityPostInfoDao godbx.QuickDao[TCommunityPostInfo] = &struct {
	godbx.QuickDao[TCommunityPostInfo]
}{
	godbx.NewBaseQuickDao(TCommunityPostInfoMeta),
}

type TCommunityPostInfo struct {
    Id int64
    UserNo string
    TripId int64
    PoiUuid string
    LocationText string
    Location string
    Province string
    ProvinceUuid int64
    CityUuid int64
    AdUuid int64
    Title string
    Image string
    PostType int32
    IsValid int32
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    LikeNum int32
    CollectNum int32
    CommentNum int32
    Status int32
    AuditDes ttypes.NilableString
    FromType int32
    
}
