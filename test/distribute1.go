package main

import "fmt"

type User struct {
	ID       int
	Username string
	Orders   int
}

func distributeOrders1(users []*User, totalOrders int) {
	// 每个普通用户分配的单数
	normalUserOrders := 50

	// 分配给普通用户
	for i := 0; i < len(users) && totalOrders > 0; i++ {
		if totalOrders >= normalUserOrders {
			users[i].Orders += normalUserOrders
			totalOrders -= normalUserOrders
		} else {
			users[i].Orders += totalOrders
			totalOrders = 0
		}
	}

	// 如果还有剩余单子，则分配给特殊用户
	if totalOrders > 0 {
		// 查找特殊用户
		specialUser := &User{}
		for _, user := range users {
			if user.Username == "Special" {
				specialUser = user
				break
			}
		}

		// 分配给特殊用户
		specialUser.Orders += totalOrders
	}

	// 输出结果
	for _, user := range users {
		fmt.Printf("%s (ID: %d) 分配的单数：%d\n", user.Username, user.ID, user.Orders)
	}
}

func main() {
	// 初始化用户
	user1 := &User{ID: 1, Username: "User1"}
	user2 := &User{ID: 2, Username: "User2"}
	user3 := &User{ID: 3, Username: "User3"}
	specialUser := &User{ID: 4, Username: "Special"}

	users := []*User{user1, user2, user3, specialUser}

	// 调用分单算法
	distributeOrders1(users, 500)
}
