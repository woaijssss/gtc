package user

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TUserFields = struct {
   Id string
   NickName string
   Avatar string
   Phone string
   UserNo string
   WxSessionKey string
   Gender string
   AccessToken string
   Score string
   IsSystem string
   CreatedAt string
   ModifiedAt string
   AvatarStatus string
   Birthday string
   SelfIntroduction string
   LocationProvinceUuid string
   LocationCityUuid string
   HomeProvinceUuid string
   HomeCityUuid string
   Zodiac string
   Constellation string
   Mbti string
   Income string
   Education string
   Occupation string
   Emotion string
   Hobby string
   
}{
    "id",
    "nick_name",
    "avatar",
    "phone",
    "user_no",
    "wx_session_key",
    "gender",
    "access_token",
    "score",
    "is_system",
    "created_at",
    "modified_at",
    "avatar_status",
    "birthday",
    "self_introduction",
    "location_province_uuid",
    "location_city_uuid",
    "home_province_uuid",
    "home_city_uuid",
    "zodiac",
    "constellation",
    "mbti",
    "income",
    "education",
    "occupation",
    "emotion",
    "hobby",
    
}

var  TUserMeta = &godbx.TableMeta[TUser]{
    Table: "t_user",
    Columns: []string {
        "id",
        "nick_name",
        "avatar",
        "phone",
        "user_no",
        "wx_session_key",
        "gender",
        "access_token",
        "score",
        "is_system",
        "created_at",
        "modified_at",
        "avatar_status",
        "birthday",
        "self_introduction",
        "location_province_uuid",
        "location_city_uuid",
        "home_province_uuid",
        "home_city_uuid",
        "zodiac",
        "constellation",
        "mbti",
        "income",
        "education",
        "occupation",
        "emotion",
        "hobby",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TUser,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "nick_name" == columnName {
            if point {
                 return &ins.NickName
            }
            return ins.NickName
        }
        if "avatar" == columnName {
            if point {
                 return &ins.Avatar
            }
            return ins.Avatar
        }
        if "phone" == columnName {
            if point {
                 return &ins.Phone
            }
            return ins.Phone
        }
        if "user_no" == columnName {
            if point {
                 return &ins.UserNo
            }
            return ins.UserNo
        }
        if "wx_session_key" == columnName {
            if point {
                 return &ins.WxSessionKey
            }
            return ins.WxSessionKey
        }
        if "gender" == columnName {
            if point {
                 return &ins.Gender
            }
            return ins.Gender
        }
        if "access_token" == columnName {
            if point {
                 return &ins.AccessToken
            }
            return ins.AccessToken
        }
        if "score" == columnName {
            if point {
                 return &ins.Score
            }
            return ins.Score
        }
        if "is_system" == columnName {
            if point {
                 return &ins.IsSystem
            }
            return ins.IsSystem
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
        if "avatar_status" == columnName {
            if point {
                 return &ins.AvatarStatus
            }
            return ins.AvatarStatus
        }
        if "birthday" == columnName {
            if point {
                 return &ins.Birthday
            }
            return ins.Birthday
        }
        if "self_introduction" == columnName {
            if point {
                 return &ins.SelfIntroduction
            }
            return ins.SelfIntroduction
        }
        if "location_province_uuid" == columnName {
            if point {
                 return &ins.LocationProvinceUuid
            }
            return ins.LocationProvinceUuid
        }
        if "location_city_uuid" == columnName {
            if point {
                 return &ins.LocationCityUuid
            }
            return ins.LocationCityUuid
        }
        if "home_province_uuid" == columnName {
            if point {
                 return &ins.HomeProvinceUuid
            }
            return ins.HomeProvinceUuid
        }
        if "home_city_uuid" == columnName {
            if point {
                 return &ins.HomeCityUuid
            }
            return ins.HomeCityUuid
        }
        if "zodiac" == columnName {
            if point {
                 return &ins.Zodiac
            }
            return ins.Zodiac
        }
        if "constellation" == columnName {
            if point {
                 return &ins.Constellation
            }
            return ins.Constellation
        }
        if "mbti" == columnName {
            if point {
                 return &ins.Mbti
            }
            return ins.Mbti
        }
        if "income" == columnName {
            if point {
                 return &ins.Income
            }
            return ins.Income
        }
        if "education" == columnName {
            if point {
                 return &ins.Education
            }
            return ins.Education
        }
        if "occupation" == columnName {
            if point {
                 return &ins.Occupation
            }
            return ins.Occupation
        }
        if "emotion" == columnName {
            if point {
                 return &ins.Emotion
            }
            return ins.Emotion
        }
        if "hobby" == columnName {
            if point {
                 return &ins.Hobby
            }
            return ins.Hobby
        }
        
        return nil
    },
}

var TUserDao godbx.QuickDao[TUser] = &struct {
	godbx.QuickDao[TUser]
}{
	godbx.NewBaseQuickDao(TUserMeta),
}

type TUser struct {
    Id int64
    NickName ttypes.NilableString
    Avatar ttypes.NilableString
    Phone ttypes.NilableString
    UserNo ttypes.NilableString
    WxSessionKey ttypes.NilableString
    Gender string
    AccessToken ttypes.NilableString
    Score int32
    IsSystem int32
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    AvatarStatus int32
    Birthday ttypes.NilableDate
    SelfIntroduction ttypes.NilableString
    LocationProvinceUuid int64
    LocationCityUuid int64
    HomeProvinceUuid int64
    HomeCityUuid int64
    Zodiac int32
    Constellation int32
    Mbti int32
    Income int32
    Education int32
    Occupation string
    Emotion int32
    Hobby string
    
}
