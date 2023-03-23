package objectvalue

//import pb "GRPC-AUTH/internal/infra/proto"

type ResponseDocuments struct {
	Title   string
	Message string
	IsOk    bool
	Status  int32
	//Value   []*pb.TypeDocument
}
