package service

import (
	"fmt"
	"os"
	"os/exec"
	"context"
	"path/filepath"
	"runtime"

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
func (s *GreeterService) DelFile(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	var file string = "/var/www/html/myfile/"
	err1 := os.Remove(file+g.Hello)
	if err1 != nil{
		fmt.Println("del err :%s",err1)
	}
	return &v1.HelloReply{Msg: "Hello " + g.Hello}, nil
}

func (s *GreeterService) GetFile(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {

	reply := &v1.HelloReply{}
	reply.Status = 0
	reply.Msg = "qiuzhiguang"


	var scan = func (fp string, fi os.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !fi.IsDir() {
			fmt.Println(fp) 			
			//temp := v1.Item{Id: "123", Audio: "/myfile/Flowers.mp3"}
			temp := v1.Item{Id: fp[20:], Audio: fp[13:]}
			reply.Data = append(reply.Data, &temp)
		}
		return nil
	}

	root := "/var/www/html/myfile" // 指定要遍历的目录
    err := filepath.WalkDir(root, scan)
    if err != nil {
        fmt.Printf("error walking the path %v: %v\n", root, err)
    }

	return reply, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) CreateFile(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})

	cmd := exec.Command("yt-dlp", "-o","/var/www/html/myfile/%(id).150B.%(ext)s","-f" ,"140","-x","--audio-format","mp3",g.Hello)	

	numCPU := runtime.NumCPU()
	if numCPU > 4 { // linode machine		
		cmd = exec.Command("echo", g.Hello)
	}	

    out, err1 := cmd.Output()

    if err1 != nil {
		fmt.Println("error exit: ")		
        fmt.Println(err1)
    } else {
        fmt.Println(string(out))
    }
	
	if err != nil {
		return nil, err
	}

	return &v1.HelloReply{Msg: "Create " + g.Hello}, nil
}
