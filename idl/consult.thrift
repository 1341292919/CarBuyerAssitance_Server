namespace go consult
include "./model.thrift"

struct ConsultRequest{
    required string budget_range,
    required string preferred_type,
    required string use_casecase,
    required string fuel_type,
    required string brand_preference,
}
struct ConsultResponse{
    1: required model.BaseResp base,
    2: required model.ConsultResult data,
}
struct QueryConsultRequest{
    1:required i64 consult_id,
}
struct QueryConsultResponse{
        1: required model.BaseResp base,
        2: required model.Consultation data,
}
struct QueryUserScoreRequest{

}
struct QueryUserScoreResponse{
            1: required model.BaseResp base,
            2: required model.PointList data,
}
struct QueryGiftRequest{

}
struct QueryGiftResponse{
                1: required model.BaseResp base,
                2: required model.GiftList data,
}
struct BuyGiftRequest{
    1:required i64 gift_id,
}
struct BuyGiftResponse{
          1: required model.BaseResp base,
          2: required model.Order data,
}
struct QueryOrderRequest{
    1: required string user_id,
}
struct QueryOrderResponse{
    1:required model.BaseResp base,
              2: required model.OrderList data,
}
service ConsultService{
    ConsultResponse Consult(1:ConsultRequest req)(api.get ="/api/consult/purchase"),
    QueryConsultResponse QueryConsult(1:QueryConsultRequest req)(api.get ="/api/consult/query"),
    QueryUserScoreResponse QueryUserScore(1:QueryUserScoreRequest req)(api.get ="/api/score/user/query"),
    QueryGiftResponse QueryGift(1:QueryGiftRequest req)(api.get="/api/score/gift/query"),
    BuyGiftResponse BuyGift(1:BuyGiftRequest req)(api.post = "/api/score/gift/purchase"),
    QueryOrderResponse QueryOrder(1:QueryOrderRequest req)(api.get="/api/score/order/query"),
}
