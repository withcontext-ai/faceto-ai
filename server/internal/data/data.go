package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/ent"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewTransaction,
	NewHttpRepo,
	NewLinkRepo,
	NewAuthRepo,
	NewVoiceRepo,
	NewStorageRepo,
	NewRoomRepo,
	NewRoomMsgRepo,
	NewRoomWebhookRepo,
	NewRoomVodRepo,
	NewLiveGPTWebhook,
)

type txCtxKey struct{}

// Data .
type Data struct {
	db *ent.Client
}

// NewTransaction .
func NewTransaction(data *Data) biz.Transaction {
	return data
}

func NewTxContext(parent context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(parent, txCtxKey{}, tx)
}

func TxFromContext(ctx context.Context) *ent.Tx {
	tx, _ := ctx.Value(txCtxKey{}).(*ent.Tx)
	return tx
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	client := ent.NewClient(ent.Driver(drv))

	if conf.Database.Debug {
		client = client.Debug()
	}

	d := &Data{
		db: client,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func (d *Data) DB(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(txCtxKey{}).(*ent.Tx)
	if ok {
		return tx.Client()
	}
	return d.db
}

func (d *Data) WithTx(ctx context.Context, fn func(context.Context) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return fn(ctx)
	}

	tx, err := d.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(NewTxContext(ctx, tx)); err != nil {
		if _err := tx.Rollback(); _err != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, _err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
