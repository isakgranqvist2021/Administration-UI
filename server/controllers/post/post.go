package post

import (
	"admin/controllers"
	"admin/models/post"
	"admin/models/user"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PublishController(c *fiber.Ctx) error {
	var publishData post.Post
	user := user.User{Email: c.Get("email")}

	if err := user.Find(); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	if err := c.BodyParser(&publishData); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	publishData.UserID = user.ID
	if err := publishData.Save(); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(controllers.HTTPResponse{
		Message: "",
		Success: true,
		Data:    nil,
	})
}

func UpdateOnePostController(c *fiber.Ctx) error {
	var post post.Post

	if err := c.BodyParser(&post); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	if err := post.Update(); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(controllers.HTTPResponse{
		Message: "Post Updated",
		Success: true,
		Data:    post,
	})
}

func DeleteOnePostController(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	post := post.Post{ID: id}
	if err := post.Delete(); err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(controllers.HTTPResponse{
		Message: "Post Deleted",
		Success: true,
		Data:    nil,
	})
}

func FindAllPostsController(c *fiber.Ctx) error {
	posts, err := post.All()

	if err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	return c.JSON(controllers.HTTPResponse{
		Message: "",
		Success: true,
		Data:    posts,
	})
}

func FindOnePostController(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		return c.JSON(controllers.HTTPResponse{
			Message: fmt.Sprintf("%v", err),
			Success: false,
			Data:    nil,
		})
	}

	post := post.Post{ID: id}
	post.Find()

	return c.JSON(controllers.HTTPResponse{
		Message: "",
		Success: true,
		Data:    post,
	})
}
