package trip

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TTripFields = struct {
   Id string
   UserNo string
   Title string
   OriginArea string
   DestArea string
   Days string
   Budget string
   Hotel string
   Transport string
   People string
   StartDate string
   Ask string
   Answer string
   Type string
   SrcProvinceUuid string
   SrcCityUuid string
   DestinationProvinceUuid string
   DestinationCityUuid string
   Comment string
   Valid string
   IsShared string
   IsDelete string
   NextTimestamp string
   CreatedAt string
   ModifiedAt string
   EssentialList string
   
}{
    "id",
    "user_no",
    "title",
    "origin_area",
    "dest_area",
    "days",
    "budget",
    "hotel",
    "transport",
    "people",
    "start_date",
    "ask",
    "answer",
    "type",
    "src_province_uuid",
    "src_city_uuid",
    "destination_province_uuid",
    "destination_city_uuid",
    "comment",
    "valid",
    "is_shared",
    "is_delete",
    "next_timestamp",
    "created_at",
    "modified_at",
    "essential_list",
    
}

var  TTripMeta = &godbx.TableMeta[TTrip]{
    Table: "t_trip",
    Columns: []string {
        "id",
        "user_no",
        "title",
        "origin_area",
        "dest_area",
        "days",
        "budget",
        "hotel",
        "transport",
        "people",
        "start_date",
        "ask",
        "answer",
        "type",
        "src_province_uuid",
        "src_city_uuid",
        "destination_province_uuid",
        "destination_city_uuid",
        "comment",
        "valid",
        "is_shared",
        "is_delete",
        "next_timestamp",
        "created_at",
        "modified_at",
        "essential_list",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TTrip,point bool) any {
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
        if "title" == columnName {
            if point {
                 return &ins.Title
            }
            return ins.Title
        }
        if "origin_area" == columnName {
            if point {
                 return &ins.OriginArea
            }
            return ins.OriginArea
        }
        if "dest_area" == columnName {
            if point {
                 return &ins.DestArea
            }
            return ins.DestArea
        }
        if "days" == columnName {
            if point {
                 return &ins.Days
            }
            return ins.Days
        }
        if "budget" == columnName {
            if point {
                 return &ins.Budget
            }
            return ins.Budget
        }
        if "hotel" == columnName {
            if point {
                 return &ins.Hotel
            }
            return ins.Hotel
        }
        if "transport" == columnName {
            if point {
                 return &ins.Transport
            }
            return ins.Transport
        }
        if "people" == columnName {
            if point {
                 return &ins.People
            }
            return ins.People
        }
        if "start_date" == columnName {
            if point {
                 return &ins.StartDate
            }
            return ins.StartDate
        }
        if "ask" == columnName {
            if point {
                 return &ins.Ask
            }
            return ins.Ask
        }
        if "answer" == columnName {
            if point {
                 return &ins.Answer
            }
            return ins.Answer
        }
        if "type" == columnName {
            if point {
                 return &ins.Type
            }
            return ins.Type
        }
        if "src_province_uuid" == columnName {
            if point {
                 return &ins.SrcProvinceUuid
            }
            return ins.SrcProvinceUuid
        }
        if "src_city_uuid" == columnName {
            if point {
                 return &ins.SrcCityUuid
            }
            return ins.SrcCityUuid
        }
        if "destination_province_uuid" == columnName {
            if point {
                 return &ins.DestinationProvinceUuid
            }
            return ins.DestinationProvinceUuid
        }
        if "destination_city_uuid" == columnName {
            if point {
                 return &ins.DestinationCityUuid
            }
            return ins.DestinationCityUuid
        }
        if "comment" == columnName {
            if point {
                 return &ins.Comment
            }
            return ins.Comment
        }
        if "valid" == columnName {
            if point {
                 return &ins.Valid
            }
            return ins.Valid
        }
        if "is_shared" == columnName {
            if point {
                 return &ins.IsShared
            }
            return ins.IsShared
        }
        if "is_delete" == columnName {
            if point {
                 return &ins.IsDelete
            }
            return ins.IsDelete
        }
        if "next_timestamp" == columnName {
            if point {
                 return &ins.NextTimestamp
            }
            return ins.NextTimestamp
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
        if "essential_list" == columnName {
            if point {
                 return &ins.EssentialList
            }
            return ins.EssentialList
        }
        
        return nil
    },
}

var TTripDao godbx.QuickDao[TTrip] = &struct {
	godbx.QuickDao[TTrip]
}{
	godbx.NewBaseQuickDao(TTripMeta),
}

type TTrip struct {
    Id int64
    UserNo string
    Title string
    OriginArea string
    DestArea string
    Days int32
    Budget int32
    Hotel int32
    Transport int32
    People int32
    StartDate string
    Ask string
    Answer string
    Type int32
    SrcProvinceUuid int64
    SrcCityUuid int64
    DestinationProvinceUuid int64
    DestinationCityUuid int64
    Comment string
    Valid int32
    IsShared int32
    IsDelete int32
    NextTimestamp int64
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    EssentialList ttypes.NilableString
    
}
