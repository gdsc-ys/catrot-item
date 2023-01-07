package handlers

import (
	"fmt"
	"log"
	"context"
	"src/database"
	"src/models"
	"flag"
	"time"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "src/proto"
)

var (
	addr = flag.String("addr", "54.180.211.115:50051", "the address to connect to")
)

func ItemList(c *fiber.Ctx) error {
	var items []models.Item
	db := database.DB
	db.Find(&items)
	return c.JSON(fiber.Map{
		"success": true,
		"itmes":   items,
	})
}

func ItemDetail(c *fiber.Ctx) error {
	var item models.Item
	itemId := c.Params("itemId")
	db := database.DB
	db.First(&item, itemId)

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFunctionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id := int32(item.UserID)  
	// id := int32(1)
	r, err := client.GetInfo(ctx, &pb.UserRequest{Id: id})

	fmt.Printf("item %T", item)	

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return c.JSON(fiber.Map{
		"success": true,
		"item":    item,
		"user":	r,
	})
}

func ItemCreate(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return err
	}
	log.Println("item 정보", item.Category, item.Content, item.Price, item.Title)
	db := database.DB
	db.Create(&item)
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["img"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			if err := c.SaveFile(file, fmt.Sprintf("./media/%s", file.Filename)); err != nil {
				return err
			}
		}
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}



func ProtoTest(c *fiber.Ctx) error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFunctionsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id := 1 
	r, err := client.GetInfo(ctx, &pb.UserRequest{Id: int32(id) })

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"response": r,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
