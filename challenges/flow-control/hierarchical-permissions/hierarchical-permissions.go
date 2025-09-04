package hierarchicalpermissions

import "fmt"

type (
	USER_ROLE   string
	PERMISSIONS string
)

func Permissions(role USER_ROLE) {
	switch role {
	case "admin":
		fmt.Println("Delete Access")
		fallthrough
	case "member":
		fmt.Println("Write Access")
		fallthrough
	case "guest":
		fmt.Println("Read Access")
	}
}

func HierarchicalPermissions() {
	var role USER_ROLE = "admin"

	Permissions(role)

}
