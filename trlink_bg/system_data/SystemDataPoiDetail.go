package system_data

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SystemDataPoiDetailFields = struct {
   Id string
   PoiId string
   Uuid string
   Introduction string
   
}{
    "id",
    "poi_id",
    "uuid",
    "introduction",
    
}

var  SystemDataPoiDetailMeta = &godbx.TableMeta[SystemDataPoiDetail]{
    Table: "system_data_poi_detail",
    Columns: []string {
        "id",
        "poi_id",
        "uuid",
        "introduction",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataPoiDetail,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "poi_id" == columnName {
            if point {
                 return &ins.PoiId
            }
            return ins.PoiId
        }
        if "uuid" == columnName {
            if point {
                 return &ins.Uuid
            }
            return ins.Uuid
        }
        if "introduction" == columnName {
            if point {
                 return &ins.Introduction
            }
            return ins.Introduction
        }
        
        return nil
    },
}

var SystemDataPoiDetailDao godbx.QuickDao[SystemDataPoiDetail] = &struct {
	godbx.QuickDao[SystemDataPoiDetail]
}{
	godbx.NewBaseQuickDao(SystemDataPoiDetailMeta),
}

type SystemDataPoiDetail struct {
    Id int64
    PoiId int64
    Uuid string
    Introduction ttypes.NilableString
    
}
