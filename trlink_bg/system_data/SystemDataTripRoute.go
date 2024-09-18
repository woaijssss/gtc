package system_data

import (
    "github.com/woaijssss/godbx"
    
    
)

var SystemDataTripRouteFields = struct {
   Id string
   Name string
   ProvinceUuid string
   CityUuid string
   
}{
    "id",
    "name",
    "province_uuid",
    "city_uuid",
    
}

var  SystemDataTripRouteMeta = &godbx.TableMeta[SystemDataTripRoute]{
    Table: "system_data_trip_route",
    Columns: []string {
        "id",
        "name",
        "province_uuid",
        "city_uuid",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataTripRoute,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
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
        
        return nil
    },
}

var SystemDataTripRouteDao godbx.QuickDao[SystemDataTripRoute] = &struct {
	godbx.QuickDao[SystemDataTripRoute]
}{
	godbx.NewBaseQuickDao(SystemDataTripRouteMeta),
}

type SystemDataTripRoute struct {
    Id int64
    Name string
    ProvinceUuid int64
    CityUuid int64
    
}
