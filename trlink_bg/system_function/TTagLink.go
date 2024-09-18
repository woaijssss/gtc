package system_function

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TTagLinkFields = struct {
   Id string
   Title string
   Sort string
   ImageUrl string
   JumpUrl string
   JumpType string
   JumpValue string
   Status string
   IsTabbar string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "title",
    "sort",
    "image_url",
    "jump_url",
    "jump_type",
    "jump_value",
    "status",
    "is_tabbar",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TTagLinkMeta = &godbx.TableMeta[TTagLink]{
    Table: "t_tag_link",
    Columns: []string {
        "id",
        "title",
        "sort",
        "image_url",
        "jump_url",
        "jump_type",
        "jump_value",
        "status",
        "is_tabbar",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TTagLink,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "title" == columnName {
            if point {
                 return &ins.Title
            }
            return ins.Title
        }
        if "sort" == columnName {
            if point {
                 return &ins.Sort
            }
            return ins.Sort
        }
        if "image_url" == columnName {
            if point {
                 return &ins.ImageUrl
            }
            return ins.ImageUrl
        }
        if "jump_url" == columnName {
            if point {
                 return &ins.JumpUrl
            }
            return ins.JumpUrl
        }
        if "jump_type" == columnName {
            if point {
                 return &ins.JumpType
            }
            return ins.JumpType
        }
        if "jump_value" == columnName {
            if point {
                 return &ins.JumpValue
            }
            return ins.JumpValue
        }
        if "status" == columnName {
            if point {
                 return &ins.Status
            }
            return ins.Status
        }
        if "is_tabbar" == columnName {
            if point {
                 return &ins.IsTabbar
            }
            return ins.IsTabbar
        }
        if "creator" == columnName {
            if point {
                 return &ins.Creator
            }
            return ins.Creator
        }
        if "updator" == columnName {
            if point {
                 return &ins.Updator
            }
            return ins.Updator
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

var TTagLinkDao godbx.QuickDao[TTagLink] = &struct {
	godbx.QuickDao[TTagLink]
}{
	godbx.NewBaseQuickDao(TTagLinkMeta),
}

type TTagLink struct {
    Id int64
    Title ttypes.NilableString
    Sort int32
    ImageUrl ttypes.NilableString
    JumpUrl ttypes.NilableString
    JumpType ttypes.NilableString
    JumpValue ttypes.NilableString
    Status int32
    IsTabbar int32
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
