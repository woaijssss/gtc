package utils

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TAppDownloadFields = struct {
   Id string
   AppVersion string
   AppUrl string
   BrowseNum string
   DownloadNum string
   CreatedAt string
   
}{
    "id",
    "app_version",
    "app_url",
    "browse_num",
    "download_num",
    "created_at",
    
}

var  TAppDownloadMeta = &godbx.TableMeta[TAppDownload]{
    Table: "t_app_download",
    Columns: []string {
        "id",
        "app_version",
        "app_url",
        "browse_num",
        "download_num",
        "created_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TAppDownload,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "app_version" == columnName {
            if point {
                 return &ins.AppVersion
            }
            return ins.AppVersion
        }
        if "app_url" == columnName {
            if point {
                 return &ins.AppUrl
            }
            return ins.AppUrl
        }
        if "browse_num" == columnName {
            if point {
                 return &ins.BrowseNum
            }
            return ins.BrowseNum
        }
        if "download_num" == columnName {
            if point {
                 return &ins.DownloadNum
            }
            return ins.DownloadNum
        }
        if "created_at" == columnName {
            if point {
                 return &ins.CreatedAt
            }
            return ins.CreatedAt
        }
        
        return nil
    },
}

var TAppDownloadDao godbx.QuickDao[TAppDownload] = &struct {
	godbx.QuickDao[TAppDownload]
}{
	godbx.NewBaseQuickDao(TAppDownloadMeta),
}

type TAppDownload struct {
    Id int64
    AppVersion ttypes.NilableString
    AppUrl ttypes.NilableString
    BrowseNum int32
    DownloadNum int32
    CreatedAt ttypes.NormalDatetime
    
}
