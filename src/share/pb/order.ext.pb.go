// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.ext.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WantTicketReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	FilmId int64 `protobuf:"varint,2,opt,name=filmId" json:"filmId,omitempty"`
	Type   int64 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
}

func (m *WantTicketReq) Reset()                    { *m = WantTicketReq{} }
func (m *WantTicketReq) String() string            { return proto.CompactTextString(m) }
func (*WantTicketReq) ProtoMessage()               {}
func (*WantTicketReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *WantTicketReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WantTicketReq) GetFilmId() int64 {
	if m != nil {
		return m.FilmId
	}
	return 0
}

func (m *WantTicketReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type WantTicketRsp struct {
}

func (m *WantTicketRsp) Reset()                    { *m = WantTicketRsp{} }
func (m *WantTicketRsp) String() string            { return proto.CompactTextString(m) }
func (*WantTicketRsp) ProtoMessage()               {}
func (*WantTicketRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type TicketReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	FilmId int64 `protobuf:"varint,2,opt,name=filmId" json:"filmId,omitempty"`
	MhId   int64 `protobuf:"varint,3,opt,name=mhId" json:"mhId,omitempty"`
	SitId  int64 `protobuf:"varint,4,opt,name=sitId" json:"sitId,omitempty"`
}

func (m *TicketReq) Reset()                    { *m = TicketReq{} }
func (m *TicketReq) String() string            { return proto.CompactTextString(m) }
func (*TicketReq) ProtoMessage()               {}
func (*TicketReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TicketReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TicketReq) GetFilmId() int64 {
	if m != nil {
		return m.FilmId
	}
	return 0
}

func (m *TicketReq) GetMhId() int64 {
	if m != nil {
		return m.MhId
	}
	return 0
}

func (m *TicketReq) GetSitId() int64 {
	if m != nil {
		return m.SitId
	}
	return 0
}

type TicketRsp struct {
}

func (m *TicketRsp) Reset()                    { *m = TicketRsp{} }
func (m *TicketRsp) String() string            { return proto.CompactTextString(m) }
func (*TicketRsp) ProtoMessage()               {}
func (*TicketRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type Order struct {
	OrderNum   string `protobuf:"bytes,1,opt,name=orderNum" json:"orderNum,omitempty"`
	OrderPrice string `protobuf:"bytes,2,opt,name=orderPrice" json:"orderPrice,omitempty"`
	CreateAt   string `protobuf:"bytes,3,opt,name=createAt" json:"createAt,omitempty"`
	PayAt      string `protobuf:"bytes,4,opt,name=payAt" json:"payAt,omitempty"`
	MhId       int64  `protobuf:"varint,5,opt,name=mhId" json:"mhId,omitempty"`
	UserId     int64  `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	MovieId    int64  `protobuf:"varint,7,opt,name=movieId" json:"movieId,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Order) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *Order) GetOrderPrice() string {
	if m != nil {
		return m.OrderPrice
	}
	return ""
}

func (m *Order) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *Order) GetPayAt() string {
	if m != nil {
		return m.PayAt
	}
	return ""
}

func (m *Order) GetMhId() int64 {
	if m != nil {
		return m.MhId
	}
	return 0
}

func (m *Order) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type MovieTicket struct {
	FilmName  string `protobuf:"bytes,1,opt,name=filmName" json:"filmName,omitempty"`
	StartTime string `protobuf:"bytes,2,opt,name=startTime" json:"startTime,omitempty"`
	Cinema    string `protobuf:"bytes,3,opt,name=cinema" json:"cinema,omitempty"`
	OrderNum  string `protobuf:"bytes,4,opt,name=orderNum" json:"orderNum,omitempty"`
}

func (m *MovieTicket) Reset()                    { *m = MovieTicket{} }
func (m *MovieTicket) String() string            { return proto.CompactTextString(m) }
func (*MovieTicket) ProtoMessage()               {}
func (*MovieTicket) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *MovieTicket) GetFilmName() string {
	if m != nil {
		return m.FilmName
	}
	return ""
}

func (m *MovieTicket) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *MovieTicket) GetCinema() string {
	if m != nil {
		return m.Cinema
	}
	return ""
}

func (m *MovieTicket) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

type LookOrdersReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *LookOrdersReq) Reset()                    { *m = LookOrdersReq{} }
func (m *LookOrdersReq) String() string            { return proto.CompactTextString(m) }
func (*LookOrdersReq) ProtoMessage()               {}
func (*LookOrdersReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *LookOrdersReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type LookOrdersRsp struct {
	MovieTickets []*MovieTicket `protobuf:"bytes,1,rep,name=movieTickets" json:"movieTickets,omitempty"`
}

func (m *LookOrdersRsp) Reset()                    { *m = LookOrdersRsp{} }
func (m *LookOrdersRsp) String() string            { return proto.CompactTextString(m) }
func (*LookOrdersRsp) ProtoMessage()               {}
func (*LookOrdersRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *LookOrdersRsp) GetMovieTickets() []*MovieTicket {
	if m != nil {
		return m.MovieTickets
	}
	return nil
}

type PayOrderReq struct {
	OrderNum string `protobuf:"bytes,1,opt,name=orderNum" json:"orderNum,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *PayOrderReq) Reset()                    { *m = PayOrderReq{} }
func (m *PayOrderReq) String() string            { return proto.CompactTextString(m) }
func (*PayOrderReq) ProtoMessage()               {}
func (*PayOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *PayOrderReq) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *PayOrderReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type PayOrderRsp struct {
}

func (m *PayOrderRsp) Reset()                    { *m = PayOrderRsp{} }
func (m *PayOrderRsp) String() string            { return proto.CompactTextString(m) }
func (*PayOrderRsp) ProtoMessage()               {}
func (*PayOrderRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

type UndoOrderReq struct {
	OrderId int64 `protobuf:"varint,1,opt,name=orderId" json:"orderId,omitempty"`
}

func (m *UndoOrderReq) Reset()                    { *m = UndoOrderReq{} }
func (m *UndoOrderReq) String() string            { return proto.CompactTextString(m) }
func (*UndoOrderReq) ProtoMessage()               {}
func (*UndoOrderReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *UndoOrderReq) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type UndoOrderRsp struct {
}

func (m *UndoOrderRsp) Reset()                    { *m = UndoOrderRsp{} }
func (m *UndoOrderRsp) String() string            { return proto.CompactTextString(m) }
func (*UndoOrderRsp) ProtoMessage()               {}
func (*UndoOrderRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

type GetOrderMessageReq struct {
	OrderNum string `protobuf:"bytes,1,opt,name=orderNum" json:"orderNum,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
}

func (m *GetOrderMessageReq) Reset()                    { *m = GetOrderMessageReq{} }
func (m *GetOrderMessageReq) String() string            { return proto.CompactTextString(m) }
func (*GetOrderMessageReq) ProtoMessage()               {}
func (*GetOrderMessageReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *GetOrderMessageReq) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *GetOrderMessageReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetOrderMessageRsp struct {
	TicketDetail *TicketDetail `protobuf:"bytes,1,opt,name=ticketDetail" json:"ticketDetail,omitempty"`
}

func (m *GetOrderMessageRsp) Reset()                    { *m = GetOrderMessageRsp{} }
func (m *GetOrderMessageRsp) String() string            { return proto.CompactTextString(m) }
func (*GetOrderMessageRsp) ProtoMessage()               {}
func (*GetOrderMessageRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *GetOrderMessageRsp) GetTicketDetail() *TicketDetail {
	if m != nil {
		return m.TicketDetail
	}
	return nil
}

type TicketDetail struct {
	FilmName      string  `protobuf:"bytes,1,opt,name=filmName" json:"filmName,omitempty"`
	FilmImg       string  `protobuf:"bytes,2,opt,name=filmImg" json:"filmImg,omitempty"`
	StartTime     string  `protobuf:"bytes,3,opt,name=startTime" json:"startTime,omitempty"`
	EndTime       string  `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`
	CinemaName    string  `protobuf:"bytes,5,opt,name=cinemaName" json:"cinemaName,omitempty"`
	MhName        string  `protobuf:"bytes,6,opt,name=mhName" json:"mhName,omitempty"`
	Seat          string  `protobuf:"bytes,7,opt,name=seat" json:"seat,omitempty"`
	OrderNum      string  `protobuf:"bytes,8,opt,name=orderNum" json:"orderNum,omitempty"`
	CinemaAddress string  `protobuf:"bytes,9,opt,name=cinemaAddress" json:"cinemaAddress,omitempty"`
	Price         float32 `protobuf:"fixed32,10,opt,name=price" json:"price,omitempty"`
	CreateAt      string  `protobuf:"bytes,11,opt,name=createAt" json:"createAt,omitempty"`
	Phone         int64   `protobuf:"varint,12,opt,name=phone" json:"phone,omitempty"`
	CinemaPhone   int64   `protobuf:"varint,13,opt,name=cinemaPhone" json:"cinemaPhone,omitempty"`
}

func (m *TicketDetail) Reset()                    { *m = TicketDetail{} }
func (m *TicketDetail) String() string            { return proto.CompactTextString(m) }
func (*TicketDetail) ProtoMessage()               {}
func (*TicketDetail) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *TicketDetail) GetFilmName() string {
	if m != nil {
		return m.FilmName
	}
	return ""
}

func (m *TicketDetail) GetFilmImg() string {
	if m != nil {
		return m.FilmImg
	}
	return ""
}

func (m *TicketDetail) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *TicketDetail) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *TicketDetail) GetCinemaName() string {
	if m != nil {
		return m.CinemaName
	}
	return ""
}

func (m *TicketDetail) GetMhName() string {
	if m != nil {
		return m.MhName
	}
	return ""
}

func (m *TicketDetail) GetSeat() string {
	if m != nil {
		return m.Seat
	}
	return ""
}

func (m *TicketDetail) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *TicketDetail) GetCinemaAddress() string {
	if m != nil {
		return m.CinemaAddress
	}
	return ""
}

func (m *TicketDetail) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TicketDetail) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *TicketDetail) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *TicketDetail) GetCinemaPhone() int64 {
	if m != nil {
		return m.CinemaPhone
	}
	return 0
}

type LookAlreadyOrdersReq struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *LookAlreadyOrdersReq) Reset()                    { *m = LookAlreadyOrdersReq{} }
func (m *LookAlreadyOrdersReq) String() string            { return proto.CompactTextString(m) }
func (*LookAlreadyOrdersReq) ProtoMessage()               {}
func (*LookAlreadyOrdersReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *LookAlreadyOrdersReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// messages:{"totalMovie":2,"oneNoComment":2,"movies":[{"filmImg":"http://img5.mtime.cn/mt/2017/10/23/101938.17733324_1280X720X2.jpg","filmName":"悟空传","time":"2017-07-13","director":"果子酱","actorNames":["彭于晏","欧豪"],"orderNum":"201701021003123"},{"filmImg":"http://img5.mtime.cn/mt/2017/10/23/101938.17733324_1280X720X2.jpg","filmName":"悟空传","time":"2017-07-13","director":"果子酱","actorNames":["彭于晏","欧豪"],"orderNum":"201701021003123"}]}
type LookAlreadyOrdersRsp struct {
	TotalMovie   int64           `protobuf:"varint,1,opt,name=totalMovie" json:"totalMovie,omitempty"`
	OneNoComment int64           `protobuf:"varint,2,opt,name=oneNoComment" json:"oneNoComment,omitempty"`
	Movies       []*AlreadyMovie `protobuf:"bytes,3,rep,name=movies" json:"movies,omitempty"`
}

func (m *LookAlreadyOrdersRsp) Reset()                    { *m = LookAlreadyOrdersRsp{} }
func (m *LookAlreadyOrdersRsp) String() string            { return proto.CompactTextString(m) }
func (*LookAlreadyOrdersRsp) ProtoMessage()               {}
func (*LookAlreadyOrdersRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *LookAlreadyOrdersRsp) GetTotalMovie() int64 {
	if m != nil {
		return m.TotalMovie
	}
	return 0
}

func (m *LookAlreadyOrdersRsp) GetOneNoComment() int64 {
	if m != nil {
		return m.OneNoComment
	}
	return 0
}

func (m *LookAlreadyOrdersRsp) GetMovies() []*AlreadyMovie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type AlreadyMovie struct {
	FilmImg    string   `protobuf:"bytes,1,opt,name=filmImg" json:"filmImg,omitempty"`
	FilmName   string   `protobuf:"bytes,2,opt,name=filmName" json:"filmName,omitempty"`
	Time       string   `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	Director   string   `protobuf:"bytes,4,opt,name=director" json:"director,omitempty"`
	ActorNames []string `protobuf:"bytes,5,rep,name=actorNames" json:"actorNames,omitempty"`
	OrderNum   string   `protobuf:"bytes,6,opt,name=orderNum" json:"orderNum,omitempty"`
}

func (m *AlreadyMovie) Reset()                    { *m = AlreadyMovie{} }
func (m *AlreadyMovie) String() string            { return proto.CompactTextString(m) }
func (*AlreadyMovie) ProtoMessage()               {}
func (*AlreadyMovie) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *AlreadyMovie) GetFilmImg() string {
	if m != nil {
		return m.FilmImg
	}
	return ""
}

func (m *AlreadyMovie) GetFilmName() string {
	if m != nil {
		return m.FilmName
	}
	return ""
}

func (m *AlreadyMovie) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *AlreadyMovie) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *AlreadyMovie) GetActorNames() []string {
	if m != nil {
		return m.ActorNames
	}
	return nil
}

func (m *AlreadyMovie) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

type OrderCommentReq struct {
	UserId         int64  `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score          int64  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	CommentContent string `protobuf:"bytes,3,opt,name=commentContent" json:"commentContent,omitempty"`
	OrderNum       string `protobuf:"bytes,4,opt,name=orderNum" json:"orderNum,omitempty"`
	MovieId        int64  `protobuf:"varint,5,opt,name=movieId" json:"movieId,omitempty"`
}

func (m *OrderCommentReq) Reset()                    { *m = OrderCommentReq{} }
func (m *OrderCommentReq) String() string            { return proto.CompactTextString(m) }
func (*OrderCommentReq) ProtoMessage()               {}
func (*OrderCommentReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

func (m *OrderCommentReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *OrderCommentReq) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *OrderCommentReq) GetCommentContent() string {
	if m != nil {
		return m.CommentContent
	}
	return ""
}

func (m *OrderCommentReq) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *OrderCommentReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

type OrderCommentRsp struct {
}

func (m *OrderCommentRsp) Reset()                    { *m = OrderCommentRsp{} }
func (m *OrderCommentRsp) String() string            { return proto.CompactTextString(m) }
func (*OrderCommentRsp) ProtoMessage()               {}
func (*OrderCommentRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func init() {
	proto.RegisterType((*WantTicketReq)(nil), "pb.WantTicketReq")
	proto.RegisterType((*WantTicketRsp)(nil), "pb.WantTicketRsp")
	proto.RegisterType((*TicketReq)(nil), "pb.TicketReq")
	proto.RegisterType((*TicketRsp)(nil), "pb.TicketRsp")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*MovieTicket)(nil), "pb.MovieTicket")
	proto.RegisterType((*LookOrdersReq)(nil), "pb.LookOrdersReq")
	proto.RegisterType((*LookOrdersRsp)(nil), "pb.LookOrdersRsp")
	proto.RegisterType((*PayOrderReq)(nil), "pb.PayOrderReq")
	proto.RegisterType((*PayOrderRsp)(nil), "pb.PayOrderRsp")
	proto.RegisterType((*UndoOrderReq)(nil), "pb.UndoOrderReq")
	proto.RegisterType((*UndoOrderRsp)(nil), "pb.UndoOrderRsp")
	proto.RegisterType((*GetOrderMessageReq)(nil), "pb.GetOrderMessageReq")
	proto.RegisterType((*GetOrderMessageRsp)(nil), "pb.GetOrderMessageRsp")
	proto.RegisterType((*TicketDetail)(nil), "pb.TicketDetail")
	proto.RegisterType((*LookAlreadyOrdersReq)(nil), "pb.LookAlreadyOrdersReq")
	proto.RegisterType((*LookAlreadyOrdersRsp)(nil), "pb.LookAlreadyOrdersRsp")
	proto.RegisterType((*AlreadyMovie)(nil), "pb.AlreadyMovie")
	proto.RegisterType((*OrderCommentReq)(nil), "pb.OrderCommentReq")
	proto.RegisterType((*OrderCommentRsp)(nil), "pb.OrderCommentRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderServiceExt service

type OrderServiceExtClient interface {
	// 想看
	WantTicket(ctx context.Context, in *WantTicketReq, opts ...client.CallOption) (*WantTicketRsp, error)
	// 下单
	Ticket(ctx context.Context, in *TicketReq, opts ...client.CallOption) (*TicketRsp, error)
	// 支付
	PayOrder(ctx context.Context, in *PayOrderReq, opts ...client.CallOption) (*PayOrderRsp, error)
	// 取消订单
	UndoOrder(ctx context.Context, in *UndoOrderReq, opts ...client.CallOption) (*UndoOrderRsp, error)
	// 根据订单编号获取电影票具体信息
	GetOrderMessage(ctx context.Context, in *GetOrderMessageReq, opts ...client.CallOption) (*GetOrderMessageRsp, error)
	// 查看所有电影票
	LookOrders(ctx context.Context, in *LookOrdersReq, opts ...client.CallOption) (*LookOrdersRsp, error)
	// 查看所有看过的电影票
	LookAlreadyOrders(ctx context.Context, in *LookAlreadyOrdersReq, opts ...client.CallOption) (*LookAlreadyOrdersRsp, error)
	// 进行评分
	OrderComment(ctx context.Context, in *OrderCommentReq, opts ...client.CallOption) (*OrderCommentRsp, error)
}

type orderServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewOrderServiceExtClient(serviceName string, c client.Client) OrderServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &orderServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *orderServiceExtClient) WantTicket(ctx context.Context, in *WantTicketReq, opts ...client.CallOption) (*WantTicketRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.WantTicket", in)
	out := new(WantTicketRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) Ticket(ctx context.Context, in *TicketReq, opts ...client.CallOption) (*TicketRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.Ticket", in)
	out := new(TicketRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) PayOrder(ctx context.Context, in *PayOrderReq, opts ...client.CallOption) (*PayOrderRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.PayOrder", in)
	out := new(PayOrderRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) UndoOrder(ctx context.Context, in *UndoOrderReq, opts ...client.CallOption) (*UndoOrderRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.UndoOrder", in)
	out := new(UndoOrderRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) GetOrderMessage(ctx context.Context, in *GetOrderMessageReq, opts ...client.CallOption) (*GetOrderMessageRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.GetOrderMessage", in)
	out := new(GetOrderMessageRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) LookOrders(ctx context.Context, in *LookOrdersReq, opts ...client.CallOption) (*LookOrdersRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.LookOrders", in)
	out := new(LookOrdersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) LookAlreadyOrders(ctx context.Context, in *LookAlreadyOrdersReq, opts ...client.CallOption) (*LookAlreadyOrdersRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.LookAlreadyOrders", in)
	out := new(LookAlreadyOrdersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceExtClient) OrderComment(ctx context.Context, in *OrderCommentReq, opts ...client.CallOption) (*OrderCommentRsp, error) {
	req := c.c.NewRequest(c.serviceName, "OrderServiceExt.OrderComment", in)
	out := new(OrderCommentRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderServiceExt service

type OrderServiceExtHandler interface {
	// 想看
	WantTicket(context.Context, *WantTicketReq, *WantTicketRsp) error
	// 下单
	Ticket(context.Context, *TicketReq, *TicketRsp) error
	// 支付
	PayOrder(context.Context, *PayOrderReq, *PayOrderRsp) error
	// 取消订单
	UndoOrder(context.Context, *UndoOrderReq, *UndoOrderRsp) error
	// 根据订单编号获取电影票具体信息
	GetOrderMessage(context.Context, *GetOrderMessageReq, *GetOrderMessageRsp) error
	// 查看所有电影票
	LookOrders(context.Context, *LookOrdersReq, *LookOrdersRsp) error
	// 查看所有看过的电影票
	LookAlreadyOrders(context.Context, *LookAlreadyOrdersReq, *LookAlreadyOrdersRsp) error
	// 进行评分
	OrderComment(context.Context, *OrderCommentReq, *OrderCommentRsp) error
}

func RegisterOrderServiceExtHandler(s server.Server, hdlr OrderServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&OrderServiceExt{hdlr}, opts...))
}

type OrderServiceExt struct {
	OrderServiceExtHandler
}

func (h *OrderServiceExt) WantTicket(ctx context.Context, in *WantTicketReq, out *WantTicketRsp) error {
	return h.OrderServiceExtHandler.WantTicket(ctx, in, out)
}

func (h *OrderServiceExt) Ticket(ctx context.Context, in *TicketReq, out *TicketRsp) error {
	return h.OrderServiceExtHandler.Ticket(ctx, in, out)
}

func (h *OrderServiceExt) PayOrder(ctx context.Context, in *PayOrderReq, out *PayOrderRsp) error {
	return h.OrderServiceExtHandler.PayOrder(ctx, in, out)
}

func (h *OrderServiceExt) UndoOrder(ctx context.Context, in *UndoOrderReq, out *UndoOrderRsp) error {
	return h.OrderServiceExtHandler.UndoOrder(ctx, in, out)
}

func (h *OrderServiceExt) GetOrderMessage(ctx context.Context, in *GetOrderMessageReq, out *GetOrderMessageRsp) error {
	return h.OrderServiceExtHandler.GetOrderMessage(ctx, in, out)
}

func (h *OrderServiceExt) LookOrders(ctx context.Context, in *LookOrdersReq, out *LookOrdersRsp) error {
	return h.OrderServiceExtHandler.LookOrders(ctx, in, out)
}

func (h *OrderServiceExt) LookAlreadyOrders(ctx context.Context, in *LookAlreadyOrdersReq, out *LookAlreadyOrdersRsp) error {
	return h.OrderServiceExtHandler.LookAlreadyOrders(ctx, in, out)
}

func (h *OrderServiceExt) OrderComment(ctx context.Context, in *OrderCommentReq, out *OrderCommentRsp) error {
	return h.OrderServiceExtHandler.OrderComment(ctx, in, out)
}

func init() { proto.RegisterFile("order.ext.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x93, 0xd4, 0x13, 0xa7, 0xa1, 0x4b, 0x55, 0x59, 0x16, 0x42, 0xd1, 0x0a, 0x41,
	0x4e, 0x91, 0x68, 0x7b, 0xe2, 0x16, 0xb5, 0x08, 0x82, 0x68, 0xa9, 0xdc, 0x22, 0xce, 0xae, 0xbd,
	0xb4, 0x56, 0xe3, 0x1f, 0xbc, 0xdb, 0xaa, 0x11, 0x57, 0x5e, 0x83, 0x77, 0xe0, 0xce, 0x7b, 0xf0,
	0x24, 0x3c, 0x00, 0xda, 0x59, 0x3b, 0xde, 0x75, 0x4a, 0x90, 0xe0, 0xb6, 0xdf, 0xb7, 0xb3, 0xb3,
	0x33, 0xb3, 0xdf, 0x8c, 0x0d, 0xc3, 0xac, 0x88, 0x58, 0x31, 0x61, 0x77, 0x62, 0x92, 0x17, 0x99,
	0xc8, 0x48, 0x2b, 0xbf, 0xa0, 0x67, 0x30, 0xf8, 0x18, 0xa4, 0xe2, 0x3c, 0x0e, 0xaf, 0x99, 0xf0,
	0xd9, 0x67, 0xb2, 0x0b, 0xdd, 0x1b, 0xce, 0x8a, 0x59, 0xe4, 0x5a, 0x23, 0x6b, 0xdc, 0xf6, 0x4b,
	0x24, 0xf9, 0x4f, 0xf1, 0x3c, 0x99, 0x45, 0x6e, 0x4b, 0xf1, 0x0a, 0x11, 0x02, 0x1b, 0x62, 0x91,
	0x33, 0xb7, 0x8d, 0x2c, 0xae, 0xe9, 0xd0, 0x70, 0xca, 0x73, 0xca, 0xc0, 0xfe, 0xaf, 0x1b, 0x92,
	0xab, 0x59, 0x54, 0xdd, 0x20, 0xd7, 0x64, 0x07, 0x3a, 0x3c, 0x16, 0xb3, 0xc8, 0xdd, 0x40, 0x52,
	0x01, 0xda, 0x5f, 0x5e, 0xc3, 0x73, 0xfa, 0xc3, 0x82, 0xce, 0x7b, 0x99, 0x31, 0xf1, 0x60, 0x13,
	0x53, 0x3f, 0xb9, 0x49, 0xf0, 0x4a, 0xdb, 0x5f, 0x62, 0xf2, 0x04, 0x00, 0xd7, 0xa7, 0x45, 0x1c,
	0x32, 0xbc, 0xd8, 0xf6, 0x35, 0x46, 0x9e, 0x0d, 0x0b, 0x16, 0x08, 0x36, 0x15, 0x18, 0x80, 0xed,
	0x2f, 0xb1, 0x0c, 0x22, 0x0f, 0x16, 0x53, 0x81, 0x41, 0xd8, 0xbe, 0x02, 0xcb, 0x70, 0x3b, 0x5a,
	0xb8, 0x75, 0xca, 0x5d, 0x23, 0x65, 0x17, 0x7a, 0x49, 0x76, 0x1b, 0xb3, 0x59, 0xe4, 0xf6, 0x70,
	0xa3, 0x82, 0xf4, 0x0b, 0xf4, 0x8f, 0xe5, 0x52, 0xe5, 0x23, 0xc3, 0x90, 0xd5, 0x38, 0x09, 0x12,
	0x56, 0xa5, 0x50, 0x61, 0xf2, 0x18, 0x6c, 0x2e, 0x82, 0x42, 0x9c, 0xc7, 0x49, 0x95, 0x41, 0x4d,
	0xc8, 0xab, 0xc3, 0x38, 0x65, 0x49, 0x50, 0x86, 0x5f, 0x22, 0xa3, 0x28, 0x1b, 0x66, 0x51, 0xe8,
	0x73, 0x18, 0xbc, 0xcb, 0xb2, 0x6b, 0xac, 0x1e, 0x5f, 0xf3, 0x64, 0xf4, 0xc8, 0x30, 0xe4, 0x39,
	0xd9, 0x07, 0x27, 0xa9, 0xc3, 0xe6, 0xae, 0x35, 0x6a, 0x8f, 0xfb, 0x7b, 0xc3, 0x49, 0x7e, 0x31,
	0xd1, 0xd2, 0xf1, 0x0d, 0x23, 0x3a, 0x85, 0xfe, 0x69, 0xb0, 0x40, 0x27, 0xf2, 0xb2, 0x75, 0xcf,
	0x55, 0x07, 0xd2, 0x32, 0x02, 0x19, 0x68, 0x2e, 0x78, 0x4e, 0xc7, 0xe0, 0x7c, 0x48, 0xa3, 0x6c,
	0xe9, 0xd2, 0x85, 0x1e, 0xba, 0x58, 0x26, 0x50, 0x41, 0xba, 0xa5, 0x5b, 0xf2, 0x9c, 0xbe, 0x01,
	0xf2, 0x9a, 0x09, 0x84, 0xc7, 0x8c, 0xf3, 0xe0, 0x92, 0xfd, 0x6b, 0x48, 0x6f, 0x57, 0x3d, 0xf1,
	0x9c, 0x1c, 0x80, 0x23, 0x30, 0xed, 0x23, 0x26, 0x82, 0x78, 0x8e, 0xde, 0xfa, 0x7b, 0x0f, 0x65,
	0x81, 0xce, 0x35, 0xde, 0x37, 0xac, 0xe8, 0xaf, 0x16, 0x38, 0xfa, 0xf6, 0x5a, 0x3d, 0xb8, 0xd0,
	0xc3, 0xce, 0x49, 0x2e, 0x4b, 0x35, 0x54, 0xd0, 0x54, 0x4a, 0xbb, 0xa9, 0x14, 0x17, 0x7a, 0x2c,
	0x8d, 0x70, 0x4f, 0x09, 0xa2, 0x82, 0xb2, 0x49, 0x94, 0x6a, 0xf0, 0xbe, 0x8e, 0x6a, 0x92, 0x9a,
	0x91, 0x25, 0x48, 0xae, 0x70, 0xaf, 0xab, 0x34, 0xa6, 0x90, 0x6c, 0x05, 0xce, 0x02, 0x81, 0xda,
	0xb6, 0x7d, 0x5c, 0x1b, 0xa5, 0xdc, 0x6c, 0x94, 0xf2, 0x29, 0x0c, 0x94, 0xd7, 0x69, 0x14, 0x15,
	0x8c, 0x73, 0xd7, 0x46, 0x03, 0x93, 0xc4, 0xb6, 0xc3, 0x6e, 0x85, 0x91, 0x35, 0x6e, 0xf9, 0x0a,
	0x18, 0x8d, 0xda, 0xbf, 0xa7, 0x51, 0xaf, 0xb2, 0x94, 0xb9, 0x8e, 0x9a, 0x16, 0x08, 0xc8, 0x08,
	0xfa, 0xca, 0xf1, 0x29, 0xee, 0x0d, 0x70, 0x4f, 0xa7, 0xe8, 0x04, 0x76, 0xa4, 0xbc, 0xa7, 0xf3,
	0x82, 0x05, 0xd1, 0xe2, 0xef, 0xed, 0xf0, 0xd5, 0xba, 0xef, 0x00, 0xcf, 0x65, 0x01, 0x45, 0x26,
	0x82, 0x39, 0xf6, 0x40, 0x79, 0x48, 0x63, 0x08, 0x05, 0x27, 0x4b, 0xd9, 0x49, 0x76, 0x98, 0x25,
	0x09, 0x4b, 0x45, 0xa9, 0x24, 0x83, 0x23, 0x63, 0xe8, 0x62, 0xd7, 0x70, 0xb7, 0x8d, 0x4d, 0x85,
	0x9a, 0x29, 0x6f, 0x42, 0x2f, 0x7e, 0xb9, 0x4f, 0xbf, 0x5b, 0xe0, 0xe8, 0x1b, 0xba, 0x22, 0x2c,
	0x53, 0x11, 0xba, 0x8e, 0x5a, 0x0d, 0x1d, 0xc9, 0xc9, 0x5e, 0x0b, 0x05, 0xd7, 0xd2, 0x3e, 0x8a,
	0x0b, 0x16, 0x8a, 0xac, 0xa8, 0xa6, 0x46, 0x85, 0x65, 0x92, 0x81, 0x5c, 0xc8, 0xc3, 0xdc, 0xed,
	0x8c, 0xda, 0x52, 0x25, 0x35, 0x63, 0xbc, 0x7c, 0xb7, 0x31, 0x71, 0xbe, 0x59, 0x30, 0xc4, 0x72,
	0x95, 0xd9, 0xae, 0xfb, 0x4e, 0xc8, 0xd9, 0x1f, 0x66, 0x05, 0x2b, 0xab, 0xa4, 0x00, 0x79, 0x06,
	0x5b, 0xa1, 0x3a, 0x7b, 0x98, 0xa5, 0x42, 0x16, 0x51, 0xc5, 0xdd, 0x60, 0xd7, 0xcd, 0x3d, 0x7d,
	0x1c, 0x77, 0xcc, 0x71, 0xbc, 0xdd, 0x08, 0x8f, 0xe7, 0x7b, 0x3f, 0xdb, 0x25, 0x77, 0xc6, 0x8a,
	0xdb, 0x38, 0x64, 0xaf, 0xee, 0x04, 0x39, 0x00, 0xa8, 0x3f, 0x7c, 0x64, 0x5b, 0xbe, 0x90, 0xf1,
	0x75, 0xf5, 0x9a, 0x14, 0xcf, 0xe9, 0x03, 0xf9, 0xb2, 0xe5, 0x89, 0x41, 0x3d, 0x07, 0xa4, 0xb5,
	0x0e, 0xd1, 0x72, 0x02, 0x9b, 0xd5, 0x98, 0x23, 0x38, 0x54, 0xb5, 0xb9, 0xe9, 0x99, 0x04, 0xda,
	0xbf, 0x00, 0x7b, 0x39, 0xdd, 0x08, 0x0a, 0x46, 0x1f, 0x8b, 0x5e, 0x83, 0xc1, 0x23, 0x87, 0x30,
	0x6c, 0x8c, 0x2d, 0xb2, 0x2b, 0xcd, 0x56, 0xa7, 0xa2, 0x77, 0x2f, 0x8f, 0x4e, 0x0e, 0x00, 0xea,
	0xef, 0x82, 0xaa, 0x83, 0xf1, 0x41, 0xf1, 0x9a, 0x14, 0x9e, 0x9a, 0xc1, 0xf6, 0x4a, 0xf7, 0x10,
	0xb7, 0xb2, 0x6c, 0x76, 0xa1, 0xf7, 0x87, 0x1d, 0x74, 0xf5, 0x12, 0x1c, 0xfd, 0xbd, 0xc8, 0x23,
	0x69, 0xdb, 0x10, 0x98, 0xb7, 0x4a, 0xca, 0xb3, 0x17, 0x5d, 0xfc, 0x3b, 0xda, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0x18, 0x50, 0xa6, 0x30, 0x09, 0x00, 0x00,
}
