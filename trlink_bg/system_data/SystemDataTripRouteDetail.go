package system_data

import (
    "github.com/woaijssss/godbx"
    
    
)

var SystemDataTripRouteDetailFields = struct {
   Id string
   RouteId string
   PoiId string
   RouteIndex string
   
}{
    "id",
    "route_id",
    "poi_id",
    "route_index",
    
}

var  SystemDataTripRouteDetailMeta = &godbx.TableMeta[SystemDataTripRouteDetail]{
    Table: "system_data_trip_route_detail",
    Columns: []string {
        "id",
        "route_id",
        "poi_id",
        "route_index",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataTripRouteDetail,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "route_id" == columnName {
            if point {
                 return &ins.RouteId
            }
            return ins.RouteId
        }
        if "poi_id" == columnName {
            if point {
                 return &ins.PoiId
            }
            return ins.PoiId
        }
        if "route_index" == columnName {
            if point {
                 return &ins.RouteIndex
            }
            return ins.RouteIndex
        }
        
        return nil
    },
}

var SystemDataTripRouteDetailDao godbx.QuickDao[SystemDataTripRouteDetail] = &struct {
	godbx.QuickDao[SystemDataTripRouteDetail]
}{
	godbx.NewBaseQuickDao(SystemDataTripRouteDetailMeta),
}

type SystemDataTripRouteDetail struct {
    Id int64
    RouteId int64
    PoiId int64
    RouteIndex int64
    
}
