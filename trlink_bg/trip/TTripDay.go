package trip

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TTripDayFields = struct {
   Id string
   TripId string
   PoiId string
   PoiUuid string
   Name string
   DayIndex string
   Title string
   Note string
   Catering string
   Accommodation string
   Transport string
   MapType string
   Longitude string
   Latitude string
   Distance string
   Image string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "trip_id",
    "poi_id",
    "poi_uuid",
    "name",
    "day_index",
    "title",
    "note",
    "catering",
    "accommodation",
    "transport",
    "map_type",
    "longitude",
    "latitude",
    "distance",
    "image",
    "created_at",
    "modified_at",
    
}

var  TTripDayMeta = &godbx.TableMeta[TTripDay]{
    Table: "t_trip_day",
    Columns: []string {
        "id",
        "trip_id",
        "poi_id",
        "poi_uuid",
        "name",
        "day_index",
        "title",
        "note",
        "catering",
        "accommodation",
        "transport",
        "map_type",
        "longitude",
        "latitude",
        "distance",
        "image",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TTripDay,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "trip_id" == columnName {
            if point {
                 return &ins.TripId
            }
            return ins.TripId
        }
        if "poi_id" == columnName {
            if point {
                 return &ins.PoiId
            }
            return ins.PoiId
        }
        if "poi_uuid" == columnName {
            if point {
                 return &ins.PoiUuid
            }
            return ins.PoiUuid
        }
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
        }
        if "day_index" == columnName {
            if point {
                 return &ins.DayIndex
            }
            return ins.DayIndex
        }
        if "title" == columnName {
            if point {
                 return &ins.Title
            }
            return ins.Title
        }
        if "note" == columnName {
            if point {
                 return &ins.Note
            }
            return ins.Note
        }
        if "catering" == columnName {
            if point {
                 return &ins.Catering
            }
            return ins.Catering
        }
        if "accommodation" == columnName {
            if point {
                 return &ins.Accommodation
            }
            return ins.Accommodation
        }
        if "transport" == columnName {
            if point {
                 return &ins.Transport
            }
            return ins.Transport
        }
        if "map_type" == columnName {
            if point {
                 return &ins.MapType
            }
            return ins.MapType
        }
        if "longitude" == columnName {
            if point {
                 return &ins.Longitude
            }
            return ins.Longitude
        }
        if "latitude" == columnName {
            if point {
                 return &ins.Latitude
            }
            return ins.Latitude
        }
        if "distance" == columnName {
            if point {
                 return &ins.Distance
            }
            return ins.Distance
        }
        if "image" == columnName {
            if point {
                 return &ins.Image
            }
            return ins.Image
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

var TTripDayDao godbx.QuickDao[TTripDay] = &struct {
	godbx.QuickDao[TTripDay]
}{
	godbx.NewBaseQuickDao(TTripDayMeta),
}

type TTripDay struct {
    Id int64
    TripId int64
    PoiId int64
    PoiUuid string
    Name string
    DayIndex int32
    Title string
    Note string
    Catering string
    Accommodation string
    Transport string
    MapType int32
    Longitude float64
    Latitude float64
    Distance float64
    Image string
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
