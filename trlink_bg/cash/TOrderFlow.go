package cash

import (
    "github.com/woaijssss/godbx"
    "github.com/woaijssss/godbx/ttypes"
    
)

var TOrderFlowFields = struct {
   Id string
   OrderNo string
   OutTradeNo string
   PayStatus string
   NotifyStatus string
   RefundNo string
   PaymentType string
   TotalAmount string
   FeeAmount string
   Param string
   Reason string
   UserNo string
   Ip string
   IpRegion string
   PayTime string
   Remarks string
   Description string
   IsDel string
   Creator string
   Updator string
   CreatedAt string
   ModifiedAt string
   
}{
    "id",
    "order_no",
    "out_trade_no",
    "pay_status",
    "notify_status",
    "refund_no",
    "payment_type",
    "total_amount",
    "fee_amount",
    "param",
    "reason",
    "user_no",
    "ip",
    "ip_region",
    "pay_time",
    "remarks",
    "description",
    "is_del",
    "creator",
    "updator",
    "created_at",
    "modified_at",
    
}

var  TOrderFlowMeta = &godbx.TableMeta[TOrderFlow]{
    Table: "t_order_flow",
    Columns: []string {
        "id",
        "order_no",
        "out_trade_no",
        "pay_status",
        "notify_status",
        "refund_no",
        "payment_type",
        "total_amount",
        "fee_amount",
        "param",
        "reason",
        "user_no",
        "ip",
        "ip_region",
        "pay_time",
        "remarks",
        "description",
        "is_del",
        "creator",
        "updator",
        "created_at",
        "modified_at",
        
    },
    AutoColumn: "id",
    LookupFieldFunc: func(columnName string,ins *TOrderFlow,point bool) any {
        if "id" == columnName {
            if point {
                 return &ins.Id
            }
            return ins.Id
        }
        if "order_no" == columnName {
            if point {
                 return &ins.OrderNo
            }
            return ins.OrderNo
        }
        if "out_trade_no" == columnName {
            if point {
                 return &ins.OutTradeNo
            }
            return ins.OutTradeNo
        }
        if "pay_status" == columnName {
            if point {
                 return &ins.PayStatus
            }
            return ins.PayStatus
        }
        if "notify_status" == columnName {
            if point {
                 return &ins.NotifyStatus
            }
            return ins.NotifyStatus
        }
        if "refund_no" == columnName {
            if point {
                 return &ins.RefundNo
            }
            return ins.RefundNo
        }
        if "payment_type" == columnName {
            if point {
                 return &ins.PaymentType
            }
            return ins.PaymentType
        }
        if "total_amount" == columnName {
            if point {
                 return &ins.TotalAmount
            }
            return ins.TotalAmount
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
        if "pay_time" == columnName {
            if point {
                 return &ins.PayTime
            }
            return ins.PayTime
        }
        if "remarks" == columnName {
            if point {
                 return &ins.Remarks
            }
            return ins.Remarks
        }
        if "description" == columnName {
            if point {
                 return &ins.Description
            }
            return ins.Description
        }
        if "is_del" == columnName {
            if point {
                 return &ins.IsDel
            }
            return ins.IsDel
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

var TOrderFlowDao godbx.QuickDao[TOrderFlow] = &struct {
	godbx.QuickDao[TOrderFlow]
}{
	godbx.NewBaseQuickDao(TOrderFlowMeta),
}

type TOrderFlow struct {
    Id int64
    OrderNo string
    OutTradeNo ttypes.NilableString
    PayStatus int8
    NotifyStatus int8
    RefundNo ttypes.NilableString
    PaymentType string
    TotalAmount float64
    FeeAmount float64
    Param ttypes.NilableString
    Reason ttypes.NilableString
    UserNo string
    Ip ttypes.NilableString
    IpRegion ttypes.NilableString
    PayTime ttypes.NilableDatetime
    Remarks ttypes.NilableString
    Description ttypes.NilableString
    IsDel int8
    Creator ttypes.NilableString
    Updator ttypes.NilableString
    CreatedAt ttypes.NormalDatetime
    ModifiedAt ttypes.NormalDatetime
    
}
