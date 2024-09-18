package system_data

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SystemDataTripRoutePlanFields = struct {
   Id string
   Uuid string
   RouteId string
   StartCoordinate string
   EndCoordinate string
   WaysPoints string
   RoutePlanType string
   
}{
    "id",
    "uuid",
    "route_id",
    "start_coordinate",
    "end_coordinate",
    "ways_points",
    "route_plan_type",
    
}

var  SystemDataTripRoutePlanMeta = &godbx.TableMeta[SystemDataTripRoutePlan]{
    Table: "system_data_trip_route_plan",
    Columns: []string {
        "id",
        "uuid",
        "route_id",
        "start_coordinate",
        "end_coordinate",
        "ways_points",
        "route_plan_type",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataTripRoutePlan,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "uuid" == columnName {
            if point {
                 return &ins.Uuid
            }
            return ins.Uuid
        }
        if "route_id" == columnName {
            if point {
                 return &ins.RouteId
            }
            return ins.RouteId
        }
        if "start_coordinate" == columnName {
            if point {
                 return &ins.StartCoordinate
            }
            return ins.StartCoordinate
        }
        if "end_coordinate" == columnName {
            if point {
                 return &ins.EndCoordinate
            }
            return ins.EndCoordinate
        }
        if "ways_points" == columnName {
            if point {
                 return &ins.WaysPoints
            }
            return ins.WaysPoints
        }
        if "route_plan_type" == columnName {
            if point {
                 return &ins.RoutePlanType
            }
            return ins.RoutePlanType
        }
        
        return nil
    },
}

var SystemDataTripRoutePlanDao godbx.QuickDao[SystemDataTripRoutePlan] = &struct {
	godbx.QuickDao[SystemDataTripRoutePlan]
}{
	godbx.NewBaseQuickDao(SystemDataTripRoutePlanMeta),
}

type SystemDataTripRoutePlan struct {
    Id int64
    Uuid int64
    RouteId int64
    StartCoordinate string
    EndCoordinate string
    WaysPoints ttypes.NilableString
    RoutePlanType int32
    
}
