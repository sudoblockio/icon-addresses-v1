package crud

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/geometry-labs/icon-addresses/models"
)

// TransactionCountByAddressModel - type for address table model
type TransactionCountByAddressModel struct {
	db            *gorm.DB
	model         *models.TransactionCountByAddress
	modelORM      *models.TransactionCountByAddressORM
	LoaderChannel chan *models.TransactionCountByAddress
}

var transactionCountByAddressModel *TransactionCountByAddressModel
var transactionCountByAddressModelOnce sync.Once

// GetAddressModel - create and/or return the addresss table model
func GetTransactionCountByAddressModel() *TransactionCountByAddressModel {
	transactionCountByAddressModelOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionCountByAddressModel = &TransactionCountByAddressModel{
			db:            dbConn,
			model:         &models.TransactionCountByAddress{},
			LoaderChannel: make(chan *models.TransactionCountByAddress, 1),
		}

		err := transactionCountByAddressModel.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionCountByAddressModel: Unable migrate postgres table: ", err.Error())
		}

		StartTransactionCountByAddressLoader()
	})

	return transactionCountByAddressModel
}

// Migrate - migrate transactionCountByAddresss table
func (m *TransactionCountByAddressModel) Migrate() error {
	// Only using TransactionCountByAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

// Insert - Insert transactionCountByAddress into table
func (m *TransactionCountByAddressModel) Insert(transactionCountByAddress *models.TransactionCountByAddress) error {
	db := m.db

	// Set table
	db = db.Model(&models.TransactionCountByAddress{})

	db = db.Create(transactionCountByAddress)

	return db.Error
}

// Select - select from transactionCountByAddresss table
func (m *TransactionCountByAddressModel) SelectOne(transactionHash string, publicKey string) (models.TransactionCountByAddress, error) {
	db := m.db

	// Set table
	db = db.Model(&models.TransactionCountByAddress{})

	// Transaction Hash
	db = db.Where("transaction_hash = ?", transactionHash)

	// Public Key
	db = db.Where("public_key = ?", publicKey)

	transactionCountByAddress := models.TransactionCountByAddress{}
	db = db.First(&transactionCountByAddress)

	return transactionCountByAddress, db.Error
}

func (m *TransactionCountByAddressModel) SelectLargestCountByPublicKey(publicKey string) (uint64, error) {
	// TODO
	return 0, nil
	db := m.db

	// Set table
	db = db.Model(&models.TransactionCountByAddress{})

	db = db.Where("public_key = ?", publicKey)

	// Get max id
	count := uint64(0)
	row := db.Select("max(count)").Row()
	row.Scan(&count)

	return count, db.Error
}

// StartTransactionCountByAddressLoader starts loader
func StartTransactionCountByAddressLoader() {
	go func() {

		for {
			// Read transactionCountByAddress
			newTransactionCountByAddress := <-GetTransactionCountByAddressModel().LoaderChannel

			// Insert
			_, err := GetTransactionCountByAddressModel().SelectOne(
				newTransactionCountByAddress.TransactionHash,
				newTransactionCountByAddress.PublicKey,
			)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Last count
				lastCount, err := GetTransactionCountByAddressModel().SelectLargestCountByPublicKey(
					newTransactionCountByAddress.PublicKey,
				)
				if err != nil {
					zap.S().Fatal(err.Error())
				}
				newTransactionCountByAddress.Count = lastCount + 1

				// Insert
				err = GetTransactionCountByAddressModel().Insert(newTransactionCountByAddress)
				if err != nil {
					zap.S().Warn(err.Error())
				}

				zap.S().Debug("Loader=TransactionCountByAddress, Address=", newTransactionCountByAddress.PublicKey, " - Insert")
			} else if err != nil {
				// Error
				zap.S().Fatal(err.Error())
			}

			///////////////////////
			// Force enrichments //
			///////////////////////
			err = reloadAddress(newTransactionCountByAddress.PublicKey)
			if err != nil {
				// Postgress error
				zap.S().Fatal(err.Error())
			}
		}
	}()
}
