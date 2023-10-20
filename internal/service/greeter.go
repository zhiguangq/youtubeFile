package service

import (
	"fmt"
	"os/exec"
	"context"

	v1 "linode/api/helloworld/v1"
	"linode/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) CreateFile(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	
	if err != nil {
		return nil, err
	}
	
	cmd := exec.Command("yt-dlp", "-o","/var/www/html/%(id).150B.%(ext)s","-f" ,"140","-x","--audio-format","mp3",g.Hello)
		
        out, err1 := cmd.Output()
        if err1 != nil {
             fmt.Println(err1)
        }    else {
            fmt.Println(string(out))
        }
	return &v1.HelloReply{Message: "Create " + g.Hello}, nil
}
