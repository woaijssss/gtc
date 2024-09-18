package services

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var SystemDataServiceHealthcareFields = struct {
   Id string
   Uuid string
   ParentId string
   Name string
   Longitude string
   Latitude string
   Address string
   Typecode string
   Pcode string
   Adcode string
   Citycode string
   MainImage string
   BusinessArea string
   Opentime string
   OpentimeWeek string
   Tel string
   Feature string
   Rating string
   Cost string
   Alias string
   Keytag string
   Rectag string
   Images string
   IsMark string
   ProvinceUuid string
   CityUuid string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "uuid",
    "parent_id",
    "name",
    "longitude",
    "latitude",
    "address",
    "typecode",
    "pcode",
    "adcode",
    "citycode",
    "main_image",
    "business_area",
    "opentime",
    "opentime_week",
    "tel",
    "feature",
    "rating",
    "cost",
    "alias",
    "keytag",
    "rectag",
    "images",
    "is_mark",
    "province_uuid",
    "city_uuid",
    "created_at",
    "modified_at",
    
}

var  SystemDataServiceHealthcareMeta = &godbx.TableMeta[SystemDataServiceHealthcare]{
    Table: "system_data_service_healthcare",
    Columns: []string {
        "id",
        "uuid",
        "parent_id",
        "name",
        "longitude",
        "latitude",
        "address",
        "typecode",
        "pcode",
        "adcode",
        "citycode",
        "main_image",
        "business_area",
        "opentime",
        "opentime_week",
        "tel",
        "feature",
        "rating",
        "cost",
        "alias",
        "keytag",
        "rectag",
        "images",
        "is_mark",
        "province_uuid",
        "city_uuid",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *SystemDataServiceHealthcare,point bool) any {
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
        if "address" == columnName {
            if point {
                 return &ins.Address
            }
            return ins.Address
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
        if "main_image" == columnName {
            if point {
                 return &ins.MainImage
            }
            return ins.MainImage
        }
        if "business_area" == columnName {
            if point {
                 return &ins.BusinessArea
            }
            return ins.BusinessArea
        }
        if "opentime" == columnName {
            if point {
                 return &ins.Opentime
            }
            return ins.Opentime
        }
        if "opentime_week" == columnName {
            if point {
                 return &ins.OpentimeWeek
            }
            return ins.OpentimeWeek
        }
        if "tel" == columnName {
            if point {
                 return &ins.Tel
            }
            return ins.Tel
        }
        if "feature" == columnName {
            if point {
                 return &ins.Feature
            }
            return ins.Feature
        }
        if "rating" == columnName {
            if point {
                 return &ins.Rating
            }
            return ins.Rating
        }
        if "cost" == columnName {
            if point {
                 return &ins.Cost
            }
            return ins.Cost
        }
        if "alias" == columnName {
            if point {
                 return &ins.Alias
            }
            return ins.Alias
        }
        if "keytag" == columnName {
            if point {
                 return &ins.Keytag
            }
            return ins.Keytag
        }
        if "rectag" == columnName {
            if point {
                 return &ins.Rectag
            }
            return ins.Rectag
        }
        if "images" == columnName {
            if point {
                 return &ins.Images
            }
            return ins.Images
        }
        if "is_mark" == columnName {
            if point {
                 return &ins.IsMark
            }
            return ins.IsMark
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

var SystemDataServiceHealthcareDao godbx.QuickDao[SystemDataServiceHealthcare] = &struct {
	godbx.QuickDao[SystemDataServiceHealthcare]
}{
	godbx.NewBaseQuickDao(SystemDataServiceHealthcareMeta),
}

type SystemDataServiceHealthcare struct {
    Id int64
    Uuid string
    ParentId string
    Name string
    Longitude float64
    Latitude float64
    Address string
    Typecode string
    Pcode int64
    Adcode int64
    Citycode int64
    MainImage string
    BusinessArea string
    Opentime string
    OpentimeWeek string
    Tel string
    Feature string
    Rating float64
    Cost int32
    Alias string
    Keytag string
    Rectag string
    Images ttypes.NilableString
    IsMark int32
    ProvinceUuid int64
    CityUuid int64
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
