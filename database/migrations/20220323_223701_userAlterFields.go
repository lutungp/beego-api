package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserAlterFields_20220323_223701 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserAlterFields_20220323_223701{}
	m.Created = "20220323_223701"

	migration.Register("UserAlterFields_20220323_223701", m)
}

// Run the migrations
func (m *UserAlterFields_20220323_223701) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE m_users ADD user_notelp varchar(50) NULL")
	m.SQL("ALTER TABLE m_users ADD user_email varchar(50) NULL")

}

// Reverse the migrations
func (m *UserAlterFields_20220323_223701) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
