// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction_count_by_public_key.proto

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

type TransactionCountByPublicKeyORM struct {
	Count           uint64 `gorm:"index:transaction_count_by_address_idx_count"`
	PublicKey       string `gorm:"primary_key"`
	TransactionHash string
}

// TableName overrides the default tablename generated by GORM
func (TransactionCountByPublicKeyORM) TableName() string {
	return "transaction_count_by_public_keys"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TransactionCountByPublicKey) ToORM(ctx context.Context) (TransactionCountByPublicKeyORM, error) {
	to := TransactionCountByPublicKeyORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCountByPublicKeyWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	to.TransactionHash = m.TransactionHash
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TransactionCountByPublicKeyWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionCountByPublicKeyORM) ToPB(ctx context.Context) (TransactionCountByPublicKey, error) {
	to := TransactionCountByPublicKey{}
	var err error
	if prehook, ok := interface{}(m).(TransactionCountByPublicKeyWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	to.TransactionHash = m.TransactionHash
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TransactionCountByPublicKeyWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TransactionCountByPublicKey the arg will be the target, the caller the one being converted from

// TransactionCountByPublicKeyBeforeToORM called before default ToORM code
type TransactionCountByPublicKeyWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionCountByPublicKeyORM) error
}

// TransactionCountByPublicKeyAfterToORM called after default ToORM code
type TransactionCountByPublicKeyWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionCountByPublicKeyORM) error
}

// TransactionCountByPublicKeyBeforeToPB called before default ToPB code
type TransactionCountByPublicKeyWithBeforeToPB interface {
	BeforeToPB(context.Context, *TransactionCountByPublicKey) error
}

// TransactionCountByPublicKeyAfterToPB called after default ToPB code
type TransactionCountByPublicKeyWithAfterToPB interface {
	AfterToPB(context.Context, *TransactionCountByPublicKey) error
}

// DefaultCreateTransactionCountByPublicKey executes a basic gorm create call
func DefaultCreateTransactionCountByPublicKey(ctx context.Context, in *TransactionCountByPublicKey, db *gorm1.DB) (*TransactionCountByPublicKey, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountByPublicKeyORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountByPublicKeyORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionCountByPublicKeyORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountByPublicKeyORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTransactionCountByPublicKey patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionCountByPublicKey(ctx context.Context, patchee *TransactionCountByPublicKey, patcher *TransactionCountByPublicKey, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TransactionCountByPublicKey, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"PublicKey" {
			patchee.PublicKey = patcher.PublicKey
			continue
		}
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
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

// DefaultListTransactionCountByPublicKey executes a gorm list call
func DefaultListTransactionCountByPublicKey(ctx context.Context, db *gorm1.DB) ([]*TransactionCountByPublicKey, error) {
	in := TransactionCountByPublicKey{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountByPublicKeyORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TransactionCountByPublicKeyORM{}, &TransactionCountByPublicKey{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountByPublicKeyORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("public_key")
	ormResponse := []TransactionCountByPublicKeyORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionCountByPublicKeyORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TransactionCountByPublicKey{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionCountByPublicKeyORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountByPublicKeyORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionCountByPublicKeyORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TransactionCountByPublicKeyORM) error
}
