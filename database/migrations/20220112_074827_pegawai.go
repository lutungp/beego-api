package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Pegawai_20220112_074827 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Pegawai_20220112_074827{}
	m.Created = "20220112_074827"

	migration.Register("Pegawai_20220112_074827", m)
}

// Run the migrations
func (m *Pegawai_20220112_074827) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE m_pegawai(`pegawai_id` int(11) NOT NULL AUTO_INCREMENT,`pegawai_uuid` varchar(128) NOT NULL,`pegawai_no` varchar(128) NOT NULL,`pegawai_nama` varchar(128) NOT NULL,`pegawai_alamat` longtext  NOT NULL,`pegawai_notelp` varchar(128) NOT NULL,`pegawai_email` varchar(128) NOT NULL,`pegawai_aktif` char(1) NOT NULL DEFAULT 'y',`created_by` int(11) NOT NULL,`created_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),`updated_by` int(11) DEFAULT NULL,`updated_date` timestamp NULL DEFAULT NULL,`disabled_by` int(11) DEFAULT NULL,`disabled_date` timestamp NULL DEFAULT NULL,`disabled_reason` varchar(100) DEFAULT NULL,`revised` smallint(6) DEFAULT 0,PRIMARY KEY (`pegawai_id`))")
}

// Reverse the migrations
func (m *Pegawai_20220112_074827) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `m_pegawai`")
}
