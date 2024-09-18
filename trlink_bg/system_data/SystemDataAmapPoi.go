package system_data

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SystemDataAmapPoiFields = struct {
   Id string
   Uuid string
   ParentId string
   Name string
   PoiType string
   IsMark string
   Longitude string
   Latitude string
   Images string
   ApiVersion string
   Typecode string
   Pcode string
   Adcode string
   Citycode string
   ProvinceUuid string
   CityUuid string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "uuid",
    "parent_id",
    "name",
    "poi_type",
    "is_mark",
    "longitude",
    "latitude",
    "images",
    "api_version",
    "typecode",
    "pcode",
    "adcode",
    "citycode",
    "province_uuid",
    "city_uuid",
    "created_at",
    "modified_at",
    
}

var  SystemDataAmapPoiMeta = &godbx.TableMeta[SystemDataAmapPoi]{
    Table: "system_data_amap_poi",
    Columns: []string {
        "id",
        "uuid",
        "parent_id",
        "name",
        "poi_type",
        "is_mark",
        "longitude",
        "latitude",
        "images",
        "api_version",
        "typecode",
        "pcode",
        "adcode",
        "citycode",
        "province_uuid",
        "city_uuid",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataAmapPoi,point bool) any {
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
        if "parent_id" == columnName {
            if point {
                 return &ins.ParentId
            }
            return ins.ParentId
        }
        if "name" == columnName {
            if point {
                 return &ins.Name
            }
            return ins.Name
        }
        if "poi_type" == columnName {
            if point {
                 return &ins.PoiType
            }
            return ins.PoiType
        }
        if "is_mark" == columnName {
            if point {
                 return &ins.IsMark
            }
            return ins.IsMark
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
        if "images" == columnName {
            if point {
                 return &ins.Images
            }
            return ins.Images
        }
        if "api_version" == columnName {
            if point {
                 return &ins.ApiVersion
            }
            return ins.ApiVersion
        }
        if "typecode" == columnName {
            if point {
                 return &ins.Typecode
            }
            return ins.Typecode
        }
        if "pcode" == columnName {
            if point {
                 return &ins.Pcode
            }
            return ins.Pcode
        }
        if "adcode" == columnName {
            if point {
                 return &ins.Adcode
            }
            return ins.Adcode
        }
        if "citycode" == columnName {
            if point {
                 return &ins.Citycode
            }
            return ins.Citycode
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

var SystemDataAmapPoiDao godbx.QuickDao[SystemDataAmapPoi] = &struct {
	godbx.QuickDao[SystemDataAmapPoi]
}{
	godbx.NewBaseQuickDao(SystemDataAmapPoiMeta),
}

type SystemDataAmapPoi struct {
    Id int64
    Uuid string
    ParentId string
    Name string
    PoiType int32
    IsMark int32
    Longitude float64
    Latitude float64
    Images ttypes.NilableString
    ApiVersion string
    Typecode int64
    Pcode int64
    Adcode int64
    Citycode int64
    ProvinceUuid int64
    CityUuid int64
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
