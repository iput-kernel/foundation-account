package domain

import (
	"regexp"
	"strings"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
)

func DetectRole(email string) *db.Role {
	// 大学のドメインではない場合nilを返す
	if !strings.HasSuffix(email, "@tks.iput.ac.jp") {
		return nil
	}

	//　学籍番号の場合は学生と判断してRoleを割り振る
	re := regexp.MustCompile(`^tk\d{6}@tks\.iput\.ac\.jp$`)
	if re.MatchString(email) {
		role := db.RoleStudent
		return &role
	}

	// それ以外は教師
	role := db.RoleTeacher
	return &role
}
