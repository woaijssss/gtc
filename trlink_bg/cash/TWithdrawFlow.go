package cash

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TWithdrawFlowFields = struct {
   Id string
   WithdrawCode string
   OutTradeNo string
   Status string
   WithdrawType string
   Amount string
   FeeAmount string
   Param string
   Reason string
   UserNo string
   UserName string
   Ip string
   IpRegion string
   Md5key string
   Remark string
   Description string
   PaymentTime string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   BatchWithdrawCode string
   BatchOutTradeNo string
   AmountType string
   Score string
   
}{
    "id",
    "withdraw_code",
    "out_trade_no",
    "status",
    "withdraw_type",
    "amount",
    "fee_amount",
    "param",
    "reason",
    "user_no",
    "user_name",
    "ip",
    "ip_region",
    "md5key",
    "remark",
    "description",
    "payment_time",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    "batch_withdraw_code",
    "batch_out_trade_no",
    "amount_type",
    "score",
    
}

var  TWithdrawFlowMeta = &godbx.TableMeta[TWithdrawFlow]{
    Table: "t_withdraw_flow",
    Columns: []string {
        "id",
        "withdraw_code",
        "out_trade_no",
        "status",
        "withdraw_type",
        "amount",
        "fee_amount",
        "param",
        "reason",
        "user_no",
        "user_name",
        "ip",
        "ip_region",
        "md5key",
        "remark",
        "description",
        "payment_time",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        "batch_withdraw_code",
        "batch_out_trade_no",
        "amount_type",
        "score",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TWithdrawFlow,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "withdraw_code" == columnName {
            if point {
                 return &ins.WithdrawCode
            }
            return ins.WithdrawCode
        }
        if "out_trade_no" == columnName {
            if point {
                 return &ins.OutTradeNo
            }
            return ins.OutTradeNo
        }
        if "status" == columnName {
            if point {
                 return &ins.Status
            }
            return ins.Status
        }
        if "withdraw_type" == columnName {
            if point {
                 return &ins.WithdrawType
            }
            return ins.WithdrawType
        }
        if "amount" == columnName {
            if point {
                 return &ins.Amount
            }
            return ins.Amount
        }
        if "fee_amount" == columnName {
            if point {
                 return &ins.FeeAmount
            }
            return ins.FeeAmount
        }
        if "param" == columnName {
            if point {
                 return &ins.Param
            }
            return ins.Param
        }
        if "reason" == columnName {
            if point {
                 return &ins.Reason
            }
            return ins.Reason
        }
        if "user_no" == columnName {
            if point {
                 return &ins.UserNo
            }
            return ins.UserNo
        }
        if "user_name" == columnName {
            if point {
                 return &ins.UserName
            }
            return ins.UserName
        }
        if "ip" == columnName {
            if point {
                 return &ins.Ip
            }
            return ins.Ip
        }
        if "ip_region" == columnName {
            if point {
                 return &ins.IpRegion
            }
            return ins.IpRegion
        }
        if "md5key" == columnName {
            if point {
                 return &ins.Md5key
            }
            return ins.Md5key
        }
        if "remark" == columnName {
            if point {
                 return &ins.Remark
            }
            return ins.Remark
        }
        if "description" == columnName {
            if point {
                 return &ins.Description
            }
            return ins.Description
        }
        if "payment_time" == columnName {
            if point {
                 return &ins.PaymentTime
            }
            return ins.PaymentTime
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
        if "batch_withdraw_code" == columnName {
            if point {
                 return &ins.BatchWithdrawCode
            }
            return ins.BatchWithdrawCode
        }
        if "batch_out_trade_no" == columnName {
            if point {
                 return &ins.BatchOutTradeNo
            }
            return ins.BatchOutTradeNo
        }
        if "amount_type" == columnName {
            if point {
                 return &ins.AmountType
            }
            return ins.AmountType
        }
        if "score" == columnName {
            if point {
                 return &ins.Score
            }
            return ins.Score
        }
        
        return nil
    },
}

var TWithdrawFlowDao godbx.QuickDao[TWithdrawFlow] = &struct {
	godbx.QuickDao[TWithdrawFlow]
}{
	godbx.NewBaseQuickDao(TWithdrawFlowMeta),
}

type TWithdrawFlow struct {
    Id int64
    WithdrawCode string
    OutTradeNo ttypes.NilableString
    Status int8
    WithdrawType string
    Amount float64
    FeeAmount float64
    Param ttypes.NilableString
    Reason ttypes.NilableString
    UserNo string
    UserName ttypes.NilableString
    Ip ttypes.NilableString
    IpRegion ttypes.NilableString
    Md5key ttypes.NilableString
    Remark ttypes.NilableString
    Description ttypes.NilableString
    PaymentTime ttypes.NilableString
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    BatchWithdrawCode ttypes.NilableString
    BatchOutTradeNo ttypes.NilableString
    AmountType int8
    Score int32
    
}
