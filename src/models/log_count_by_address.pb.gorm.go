// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log_count_by_address.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type LogCountByAddressORM struct {
	Count           uint64 `gorm:"index:log_count_by_address_idx_count"`
	LogIndex        uint64 `gorm:"primary_key"`
	PublicKey       string `gorm:"index:log_count_by_address_idx_public_key"`
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (LogCountByAddressORM) TableName() string {
	return "log_count_by_addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *LogCountByAddress) ToORM(ctx context.Context) (LogCountByAddressORM, error) {
	to := LogCountByAddressORM{}
	var err error
	if prehook, ok := interface{}(m).(LogCountByAddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.PublicKey = m.PublicKey
	to.Count = m.Count
	if posthook, ok := interface{}(m).(LogCountByAddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *LogCountByAddressORM) ToPB(ctx context.Context) (LogCountByAddress, error) {
	to := LogCountByAddress{}
	var err error
	if prehook, ok := interface{}(m).(LogCountByAddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.PublicKey = m.PublicKey
	to.Count = m.Count
	if posthook, ok := interface{}(m).(LogCountByAddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type LogCountByAddress the arg will be the target, the caller the one being converted from

// LogCountByAddressBeforeToORM called before default ToORM code
type LogCountByAddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *LogCountByAddressORM) error
}

// LogCountByAddressAfterToORM called after default ToORM code
type LogCountByAddressWithAfterToORM interface {
	AfterToORM(context.Context, *LogCountByAddressORM) error
}

// LogCountByAddressBeforeToPB called before default ToPB code
type LogCountByAddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *LogCountByAddress) error
}

// LogCountByAddressAfterToPB called after default ToPB code
type LogCountByAddressWithAfterToPB interface {
	AfterToPB(context.Context, *LogCountByAddress) error
}

// DefaultCreateLogCountByAddress executes a basic gorm create call
func DefaultCreateLogCountByAddress(ctx context.Context, in *LogCountByAddress, db *gorm1.DB) (*LogCountByAddress, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountByAddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountByAddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type LogCountByAddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountByAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskLogCountByAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskLogCountByAddress(ctx context.Context, patchee *LogCountByAddress, patcher *LogCountByAddress, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*LogCountByAddress, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
		if f == prefix+"PublicKey" {
			patchee.PublicKey = patcher.PublicKey
			continue
		}
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListLogCountByAddress executes a gorm list call
func DefaultListLogCountByAddress(ctx context.Context, db *gorm1.DB) ([]*LogCountByAddress, error) {
	in := LogCountByAddress{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountByAddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &LogCountByAddressORM{}, &LogCountByAddress{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountByAddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []LogCountByAddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountByAddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*LogCountByAddress{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type LogCountByAddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountByAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountByAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]LogCountByAddressORM) error
}
