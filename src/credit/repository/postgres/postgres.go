package repository

import (
	"context"
	"kredit-plus/src/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	dbPool *pgxpool.Pool
}

func New(dbPool *pgxpool.Pool) domain.Database {
	return &postgres{
		dbPool: dbPool,
	}
}

func (p *postgres) InsertCustomer(ctx context.Context, customer domain.Customer) error {

	query := `
		INSERT INTO customer (nik, fullname, legalname, birthplace, birthdate, salary, photoktp, photoselfie)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := p.dbPool.Exec(ctx, query, customer.NIK, customer.FullName, customer.LegalName, customer.BirthPlace, customer.BirthDate, customer.Salary, customer.PhotoKTP, customer.PhotoSelfie)
	if err != nil {
		return err
	}
	return nil
}