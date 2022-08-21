package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
)

type StreetMarketRepository struct {
	db *sql.DB
}

var ()

func NewStreetMarketRepository(db *sql.DB) *StreetMarketRepository {
	return &StreetMarketRepository{db}
}

func (r *StreetMarketRepository) List(
	ctx context.Context,
	query domain.StreetMarketFilter,
) ([]domain.StreetMarket, error) {
	bq := "SELECT * FROM street_market"

	cls, vls, args := buildArgs(query)

	where := []string{}

	for i := 0; i < len(cls); i++ {
		where = append(where, fmt.Sprintf("%s = %s", cls[i], vls[i]))
	}

	if len(where) > 0 {
		bq = fmt.Sprintf("%s WHERE %s", bq, strings.Join(where, " AND "))
	}

	res, err := r.db.QueryContext(ctx, bq, args...)
	if err != nil {
		return nil, fmt.Errorf("%w", domain.ErrUnexpected)
	}

	rrs := []domain.StreetMarket{}
	for res.Next() {
		sm := domain.StreetMarket{}
		if err := res.Scan(
			&sm.ID,
			&sm.Long,
			&sm.Lat,
			&sm.SectCens,
			&sm.Area,
			&sm.IDdist,
			&sm.District,
			&sm.IDSubTH,
			&sm.SubTownHall,
			&sm.Region5,
			&sm.Region8,
			&sm.Name,
			&sm.Register,
			&sm.Street,
			&sm.Number,
			&sm.Neighborhood,
			&sm.AddrExtraInfo,
		); err != nil {
			return nil, fmt.Errorf("%w", domain.ErrUnexpected)
		}
		rrs = append(rrs, sm)
	}

	return rrs, nil
}

func (r *StreetMarketRepository) Create(ctx context.Context, streetMarket domain.StreetMarket) error {
	bq := "INSERT INTO street_market (%s) VALUES (%s)"

	cl, vls, args := buildArgs(streetMarket)

	q := fmt.Sprintf(bq, strings.Join(cl, ","), strings.Join(vls, ","))

	qr, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	if ra < 1 {
		return domain.ErrNothingCreated
	}

	return nil
}

func (r *StreetMarketRepository) Update(ctx context.Context, sm domain.StreetMarket) error {
	id := sm.ID
	sm.ID = ""

	cl, vls, args := buildArgs(sm)
	lArgs := len(cl)

	bq := "UPDATE street_market SET %s WHERE id = $%v"

	set := []string{}
	for i := 0; i < lArgs; i++ {
		set = append(set, fmt.Sprintf("%s = %s", cl[i], vls[i]))
	}

	q := fmt.Sprintf(bq, strings.Join(set, ","), lArgs+1)

	args = append(args, id)
	qr, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	if ra < 1 {
		return domain.ErrNothingUpdated
	}

	return nil
}

func (r *StreetMarketRepository) DeleteByID(ctx context.Context, ID string) error {
	q := "DELETE FROM street_market WHERE id = $1"

	qr, err := r.db.ExecContext(ctx, q, ID)
	if err != nil {
		return fmt.Errorf("%w", domain.ErrUnexpected)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w", domain.ErrNothingDeleted)
	}

	if ra < 1 {
		return domain.ErrNothingDeleted
	}

	return nil
}

func buildArgs(inp interface{}) (columns, placeHolders []string, values []interface{}) {
	v := reflect.ValueOf(inp)
	t := reflect.TypeOf(inp)

	phC := 1

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value != "" && value != 0 && value != 0.0 {
			values = append(values, value)
			placeHolders = append(placeHolders, fmt.Sprintf("$%v", phC))
			phC += 1
			cn := string(t.Field(i).Name)
			columns = append(columns, strings.ToLower(cn))
		}
	}

	return columns, placeHolders, values
}
