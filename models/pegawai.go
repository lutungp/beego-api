package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type MPegawai struct {
	PegawaiId     int64  `orm:"auto;" json:"-"`
	PegawaiUuid   string `orm:"size(128)"`
	PegawaiNo     string `orm:"size(128)"`
	PegawaiNama   string `orm:"size(128)"`
	PegawaiAlamat string `orm:"type(longtext)"`
	PegawaiNotelp string `orm:"size(128)"`
	PegawaiEmail  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(MPegawai))
}

// AddPegawai insert a new MPegawai into database and returns
// last inserted PegawaiUuid on success.
func AddPegawai(m *MPegawai) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPegawaiById retrieves MPegawai by PegawaiUuid. Returns error if
// PegawaiUuid doesn't exist
func GetPegawaiById(id string) (v *MPegawai, err error) {
	o := orm.NewOrm()
	v = &MPegawai{PegawaiUuid: id}

	if err = o.QueryTable(new(MPegawai)).
		Filter("PegawaiUuid", id).
		RelatedSel().
		One(v); err == nil {

		return v, nil
	}
	return nil, err
}

// GetAllPegawai retrieves all MPegawai matches certain condition. Returns empty list if
// no records exist
func GetAllPegawai(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MPegawai))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MPegawai
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePegawai updates MPegawai by PegawaiUuid and returns error if
// the record to be updated doesn't exist
func UpdatePegawaiById(m *MPegawai) (err error) {
	o := orm.NewOrm()
	v := MPegawai{PegawaiUuid: m.PegawaiUuid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePegawai deletes MPegawai by PegawaiUuid and returns error if
// the record to be deleted doesn't exist
func DeletePegawai(id string) (err error) {
	o := orm.NewOrm()
	v := MPegawai{PegawaiUuid: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MPegawai{PegawaiUuid: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
