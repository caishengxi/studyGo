package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID      uint
	Content     string `gorm:"size:255"`
	Likes       int    `gorm:"default:0"`
	PublishedAt string `gorm:"type:datetime"`
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {

	if c.PostID > 0 {

		var count int64
		res := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

		if res.Error == nil {

			if count <= 0 {

				return tx.Model(&Post{}).
					Where("id = ?", c.PostID).
					Update("status", "无评论").Error
			}

		} else {

			return res.Error
		}

	}

	return nil
}

type Post struct {
	gorm.Model
	Title       string `gorm:"size:255"`
	Content     string `gorm:"size:255"`
	PublishedAt string `gorm:"type:datetime"`
	Likes       int    `gorm:"default:0"`
	UserID      uint
	Comments    []Comment `gorm:"foreignKey:PostID"`
	Status      string    `gorm:"size:255;default:默认"`
}

func (post *Post) AfterCreate(tx *gorm.DB) error {

	result := tx.Model(&User{}).
		Where("id = ?", post.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (post *Post) AfterDelete(tx *gorm.DB) error {

	result := tx.Model(&User{}).
		Where("id = ?", post.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count - ?", 1))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

type User struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255;uniqueIndex"`
	PostCount int    `gorm:"not null;default:0"`
	Posts     []Post `gorm:"foreignKey:UserID"`
}

func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(err)
	}
	return db

}

func ConnectDB() *gorm.DB {

	//连接到数据库文件，如果不存在则新建
	db, err := gorm.Open(sqlite.Open("myhomework.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db

}

func InitData(db *gorm.DB) {
	users := []User{
		{
			ID: 1, Name: "Alice", Email: "alice@sina.com",
			Posts: []Post{
				{
					ID:          1,
					Title:       "My First Post alice",
					Content:     "This is my first post.",
					PublishedAt: "2026-06-03 09:09:21",
					Likes:       12,
					UserID:      1,
					Comments: []Comment{
						{ID: 11, Content: "this is a Great post!", Likes: 5, PostID: 1, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 12, Content: "Thanks for sharing.", Likes: 3, PostID: 1, PublishedAt: "2026-07-01 10:42:49"},
						{ID: 13, Content: "aThanks for sharing.", Likes: 13, PostID: 1, PublishedAt: "2026-07-01 10:42:49"},
						{ID: 14, Content: "bThanks for sharing.", Likes: 3, PostID: 1, PublishedAt: "2026-07-01 10:42:49"},
					},
				},
			},
		},
		{
			ID: 2, Name: "Bob", Email: "bob@sina.com",
			Posts: []Post{
				{
					ID:          2,
					Title:       "My First Post",
					Content:     "This is my first post.",
					PublishedAt: "2026-06-03 09:09:21",
					Likes:       10,
					UserID:      2,
					Comments: []Comment{
						{ID: 21, Content: "this is a Great post!", Likes: 5, PostID: 2, PublishedAt: "2026-06-01 11:08:06"},
					},
				},
				{
					ID:          3,
					Title:       "My First Post",
					Content:     "This is my first post.",
					PublishedAt: "2026-06-03 09:09:21",
					Likes:       54,
					UserID:      2,
					Comments: []Comment{
						{ID: 22, Content: "评论牛啊1!", Likes: 5, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 23, Content: "很有趣的!", Likes: 5, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 24, Content: "真厉害!", Likes: 7, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 25, Content: "把散户挂在高处!", Likes: 0, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 26, Content: "这。。。。", Likes: 5, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 27, Content: "哈哈哈哈。", Likes: 3, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
						{ID: 28, Content: "这么说好吗，嘻嘻", Likes: 5, PostID: 3, PublishedAt: "2026-06-01 11:08:06"},
					},
				},
			},
		},
		{
			ID: 3, Name: "才生息", Email: "shengxi@qq.com",
			Posts: []Post{
				{
					ID:          4,
					Title:       "学习go语言1",
					Content:     "基本语法： 1 2 3",
					PublishedAt: "2026-06-01 19:39:01",
					Likes:       163,
					UserID:      3,
					Comments: []Comment{
						{ID: 31, Content: "很好!", Likes: 5, PostID: 4, PublishedAt: "2026-06-01 21:08:06"},
						{ID: 32, Content: "学到了很多.", Likes: 3, PostID: 4, PublishedAt: "2026-07-04 10:42:49"},
						{ID: 33, Content: "Thanks for sharing.", Likes: 13, PostID: 4, PublishedAt: "2026-07-06 10:42:49"},
					},
				},
				{
					ID:          5,
					Title:       "学习go语言2",
					Content:     "高并发， 锁",
					PublishedAt: "2026-06-02 09:09:21",
					Likes:       71,
					UserID:      3,
					Comments: []Comment{
						{ID: 34, Content: "高并发好啊!", Likes: 5, PostID: 5, PublishedAt: "2026-07-01 11:08:06"},
						{ID: 35, Content: "Thanks for sharing.", Likes: 3, PostID: 5, PublishedAt: "2026-07-01 18:42:49"},
					},
				},
				{
					ID:          6,
					Title:       "学习go语言3",
					Content:     "数据库操作",
					PublishedAt: "2026-06-03 21:00:49",
					Likes:       18,
					UserID:      3,
					Comments: []Comment{
						{ID: 36, Content: "这个数据库很有趣!", Likes: 15, PostID: 6, PublishedAt: "2026-06-04 11:08:06"},
					},
				},
			},
		},
	}

	db.Create(&users)

}
func FindUserAllPosts(db *gorm.DB, userid uint) {
	var user User
	if err := db.Preload("Posts.Comments").First(&user, userid).Error; err != nil {
		panic(err)
	}

	fmt.Printf("========== 用户信息 =========\n")
	fmt.Printf("ID: %d\n姓名: %s\n邮箱: %s\n文章数: %d\n", user.ID, user.Name, user.Email, user.PostCount)
	if len(user.Posts) == 0 {
		fmt.Println("该用户暂无文章")
		return
	}

	fmt.Printf("========== 用户文章（%d 篇） =========\n", len(user.Posts))
	for _, post := range user.Posts {
		fmt.Printf("文章 ID: %d\n标题: %s\n内容: %s\n点赞数: %d\n状态: %s\n",
			post.ID, post.Title, post.Content, post.Likes, post.Status)
		if len(post.Comments) == 0 {
			fmt.Println("评论: 暂无评论")
			continue
		}

		fmt.Printf("评论（%d 条）:\n", len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("  - 评论 ID: %d | 内容: %s | 点赞数: %d \n",
				comment.ID, comment.Content, comment.Likes)
		}
		fmt.Println("------------------------------")
	}

}

func FindMostCommentPost(db *gorm.DB) {
	var res struct {
		PostID       uint
		Title        string
		Content      string
		CommentCount int64
	}

	tx := db.Model(&Comment{}).
		Select("posts.id AS post_id, posts.title, posts.content, COUNT(comments.id) AS comment_count").
		Joins("JOIN posts ON posts.id = comments.post_id").
		Group("posts.id, posts.title, posts.content").
		Order("comment_count DESC").
		Limit(1).
		Scan(&res)
	if tx.Error != nil {
		panic(tx.Error)
	}
	if tx.RowsAffected == 0 {
		fmt.Println("没有带评论的文章")
		return
	}

	fmt.Printf("评论最多的文章：ID=%d, 标题=%s, 内容=%s, 评论数=%d\n",
		res.PostID, res.Title, res.Content, res.CommentCount)
}

func AddOnePost(db *gorm.DB, userid uint, postid uint) {
	newpost := Post{
		ID:          postid,
		Title:       "My new Post",
		Content:     "This is my new post.",
		PublishedAt: "2026-10-03 09:09:21",
		Likes:       1,
		UserID:      userid,
	}

	db.Create(&newpost)

}

func GetUserPostCount(db *gorm.DB, userid uint) int {
	var user User
	db.First(&user, userid)

	return user.PostCount
}

func GetPostCommentStatus(db *gorm.DB, postid uint) string {
	var post Post
	db.First(&post, postid)
	return post.Status
}

func RemoveOneComment(db *gorm.DB, commentid uint) {
	var comment Comment
	if err := db.First(&comment, commentid).Error; err != nil {
		panic(err)
	}
	if err := db.Delete(&comment).Error; err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("===========开始===========")
	db := InitDB()
	//初始化数据
	InitData(db)

	fmt.Println("===========数据初始化完成===========")

	fmt.Println("*********************************************************")
	fmt.Println("第二题：联合查询")

	fmt.Println("查询某用户的所有文章和评论   结果：")
	//查询某用户的所有文章和评论
	FindUserAllPosts(db, 3)
	fmt.Println("查询评论最多的文章			结果：")
	//查询评论最多的文章
	FindMostCommentPost(db)
	fmt.Println("*********************************************************")

	fmt.Println("第三题：钩子hook")
	//增加一篇文章，看user的post_count会不会自动+1
	fmt.Println("增加一篇文章，看user的post_count会不会自动+1. 结果：")

	fmt.Printf("用户 %d 增加前 文章数量:%d\n", 3, GetUserPostCount(db, 3))
	AddOnePost(db, 3, 10)
	fmt.Printf("用户 %d 增加后 文章数量:%d\n", 3, GetUserPostCount(db, 3))

	fmt.Println("-------------------------------------------------------")

	//删除一条评论，当文章评论数量为0 的时候 ，看文章的status 会不会有变化
	fmt.Println("删除一条评论，当文章评论数量为0 的时候 ，看文章的status 会不会有变化  结果：")

	fmt.Printf("文章 %d 删除前 状态为：%s\n", 2, GetPostCommentStatus(db, 2))
	RemoveOneComment(db, 21)
	fmt.Printf("文章%d 删除后 状态为：%s\n", 2, GetPostCommentStatus(db, 2))

	fmt.Println("=========执行完成==========")

	/*
		//db := ConnectDB()
		//db := InitDB()
		//删除数据
		// db.Delete(&Post{}, []int{1, 2})
		// db.Delete(&Comment{}, []int{1, 2, 3, 4})
		// user := User{ID: 3, Name: "Bob", Email: "bob@sina.com",
		// 	Posts: []Post{
		// 		{
		// 			ID:          6,
		// 			Title:       "My First Post",
		// 			Content:     "This is my first post.",
		// 			PublishedAt: "2026-06-03 09:09:21",
		// 			Likes:       10,
		// 			UserID:      3,
		// 			Comments: []Comment{
		// 				{ID: 31, Content: "this is a Great post!", Likes: 5, PostID: 6, PublishedAt: "2026-06-01 11:08:06"},
		// 				{ID: 32, Content: "Thanks for sharing.", Likes: 3, PostID: 6, PublishedAt: "2026-07-01 10:42:49"},
		// 				{ID: 33, Content: "aThanks for sharing.", Likes: 13, PostID: 6, PublishedAt: "2026-07-01 10:42:49"},
		// 				{ID: 34, Content: "bThanks for sharing.", Likes: 3, PostID: 6, PublishedAt: "2026-07-01 10:42:49"},
		// 			},
		// 		},
		// 	},
		// }
		//db.Model(&User{}).Where("ID = ?", 1).Update("Name", "才生息")
		//db.Create(&user)
		//db.Model(&user).Update()
		//db.Save(&user)
		// var user User
		// db.First(&user, 1)
		// user.Email = "390809746@qq.com"
		// user.Name = "才生息"
		// db.Save(&user)
		// err := db.Transaction(func(tx *gorm.DB) error {
		// 	var existing User
		// 	err := tx.Where("email = ?", user.Email).First(&existing).Error
		// 	if err == nil {
		// 		postIDs := tx.Model(&Post{}).Select("id").Where("user_id = ?", existing.ID)
		// 		if err := tx.Where("post_id IN (?)", postIDs).Delete(&Comment{}).Error; err != nil {
		// 			return err
		// 		}
		// 		if err := tx.Where("user_id = ?", existing.ID).Delete(&Post{}).Error; err != nil {
		// 			return err
		// 		}
		// 		if err := tx.Delete(&existing).Error; err != nil {
		// 			return err
		// 		}
		// 	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 		return err
		// 	}
		// 	return tx.Create(&user).Error
		// })
		// if err != nil {
		// 	panic(err)
		// }

		//fmt.Printf("用户及关联帖子、评论插入成功，用户ID: %d\n", user.ID)

		//用户1的所有文章和评论
		// var posts []Post
		// db.Preload("Comments").Where("User_ID = ?", 3).Find(&posts)
		// fmt.Println(posts)

		// //评论最多的文章
		// var res struct {
		// 	PostID uint
		// 	Count  int64
		// }

		// db.Model(&Comment{}).
		// 	Select("post_id ,COUNT(*) as count").
		// 	Group("post_id").
		// 	Order("count DESC").
		// 	Limit(1).
		// 	Scan(&res)

		// fmt.Println(res)
		// fmt.Printf("评论最多的文章id是 %d,评论数是 %d\n", res.PostID, res.Count)

		// type UserPost struct {
		// 	UserId    uint
		// 	PostCount int
		// }
		// var resarray []UserPost
		// db.Model(&Post{}).
		// 	Select("user_id,COUNT(*) as post_count").
		// 	Group("user_id").Scan(&resarray)

		// fmt.Println(resarray)
		// for _, userpost := range resarray {
		// 	db.Model(&User{}).Where("id = ?", userpost.UserId).Update("post_count", userpost.PostCount)
		// }

		// newpost := Post{
		// 	ID:          7,
		// 	Title:       "My new Post",
		// 	Content:     "This is my new post.",
		// 	PublishedAt: "2026-05-03 09:09:21",
		// 	Likes:       7,
		// 	UserID:      2,
		// }

		// db.Create(&newpost)

		//db.Delete(&Post{}, 7)
		// var comment Comment
		// if err := db.Unscoped().First(&comment, 32).Error; err != nil {
		// 	panic(err)
		// }
		// if err := db.Unscoped().Delete(&comment).Error; err != nil {
		// 	panic(err)
		// }
	*/

}
