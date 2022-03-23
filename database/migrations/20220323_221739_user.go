package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20220323_221739 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20220323_221739{}
	m.Created = "20220323_221739"

	migration.Register("User_20220323_221739", m)
}

// Run the migrations
func (m *User_20220323_221739) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("create table m_users (user_id int(11) not null auto_increment, user_uuid varchar(100) not null, user_noanggota varchar(100) not null, user_role varchar(50) null, user_nama varchar(100) not null, user_password varchar(200) null, user_authid varchar(100) not null, user_authtoken text null, user_refreshtoken text null, user_validasi_date timestamp(0) null, user_validasi_by int4 null, user_aktif char(1) NOT NULL DEFAULT 'y', created_by int4 not null, created_date timestamp(0) not null, updated_by int4 null, updated_date timestamp(0) null, revised int4 null default 0, disabled_by int4 null, disabled_alasan varchar(255) null, disabled_date timestamp(0) null, constraint m_users_pkey primary key (user_id))")
}

// Reverse the migrations
func (m *User_20220323_221739) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `m_users`")
}
