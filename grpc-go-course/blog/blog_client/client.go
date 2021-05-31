package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/tetronoz/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client")

	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	// create Blog

	// fmt.Println("Creating a blog")
	// blog := &blogpb.Blog{
	// 	AuthorId: "Sergey",
	// 	Title: "My first blog",
	// 	Content: "Content of the first blog",
	// }

	// blogResponse, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	// if err != nil {
	// 	log.Fatalf("Unexpected error: %v", err)
	// } else {
	// 	fmt.Printf("Blog has been created: %v", blogResponse)
	// }

	// read Blog

	// fmt.Println("Reading a blog")

	// oid := "60b0e8dee3fe1d707fbbcc4d"

	// readBlogResponse, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: oid})
	// if err != nil {
	// 	log.Fatalf("Error reading blog: %v", err)
	// }

	// fmt.Println(readBlogResponse.Blog.Content)

	// fmt.Println("Update a blog")

	// newBlog := &blogpb.Blog{
	// 	Id: oid,
	// 	Title: "Updated title",
	// 	Content: "Updated content",
	// 	AuthorId: "Sergey Tolmachev",
	// }

	// _, err2 := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	// if err2 != nil {
	// 	log.Fatalf("Could not updated the blog: %v", err2)
	// }
	// fmt.Println("Blog was updated")

	// Delete a blog
	// fmt.Println("Delete a blog")
	// oid := "60b0e8dee3fe1d707fbbcc4d"

	// deleteBlogResponse, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: oid})
	// if err != nil {
	// 	log.Fatalf("Error deleting blog: %v", err)
	// }
	// fmt.Printf("Blog with id %v was deleted\n", deleteBlogResponse.BlogId)

	// List blog
	fmt.Println("List a blog")
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling list blog RPC: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something bad happend: %v\n", err)
		}
		fmt.Println(res.GetBlog())
	}
}