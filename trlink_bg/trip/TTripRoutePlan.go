package trip

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TTripRoutePlanFields = struct {
   Id string
   TripId string
   StartCoordinate string
   EndCoordinate string
   WaysPoints string
   RoutePlanType string
   
}{
    "id",
    "trip_id",
    "start_coordinate",
    "end_coordinate",
    "ways_points",
    "route_plan_type",
    
}

var  TTripRoutePlanMeta = &godbx.TableMeta[TTripRoutePlan]{
    Table: "t_trip_route_plan",
    Columns: []string {
        "id",
        "trip_id",
        "start_coordinate",
        "end_coordinate",
        "ways_points",
        "route_plan_type",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TTripRoutePlan,point bool) any {
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

var TTripRoutePlanDao godbx.QuickDao[TTripRoutePlan] = &struct {
	godbx.QuickDao[TTripRoutePlan]
}{
	godbx.NewBaseQuickDao(TTripRoutePlanMeta),
}

type TTripRoutePlan struct {
    Id int64
    TripId int64
    StartCoordinate string
    EndCoordinate string
    WaysPoints ttypes.NilableString
    RoutePlanType int32
    
}
