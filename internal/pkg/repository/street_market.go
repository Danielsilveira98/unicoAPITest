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
	pg domain.Pagination,
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

	q := fmt.Sprintf("%s ORDER BY createdat DESC OFFSET %v LIMIT %v", bq, pg.Offset, pg.Limit)
	res, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("[QueryContext] %w", err)
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
			&sm.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("[Scan] %w", domain.ErrUnexpected)
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
		return fmt.Errorf("[ExecContext] %w", err)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("[RowsAffected] %w", err)
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
		return fmt.Errorf("[ExecContext] %w", err)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("[RowsAffected] %w", err)
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
		return fmt.Errorf("[ExecContext] %w", err)
	}

	ra, err := qr.RowsAffected()
	if err != nil {
		return fmt.Errorf("[RowsAffected] %w", err)
	}

	if ra < 1 {
		return domain.ErrNothingDeleted
	}

	return nil
}

func buildArgs(inp interface{}) (columns, placeHolders []string, values []interface{}) {
	v := reflect.ValueOf(inp)
	t := reflect.TypeOf(inp)

	asEmpty := []any{"", 0, 0.0, nil}

	blacklist := []any{"createdat"}

	phC := 1
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		cn := strings.ToLower(string(t.Field(i).Name))
		if !ignore(value, asEmpty) && !ignore(cn, blacklist) {
			values = append(values, value)
			placeHolders = append(placeHolders, fmt.Sprintf("$%v", phC))
			phC += 1
			columns = append(columns, cn)
		}
	}

	return columns, placeHolders, values
}

func ignore(value any, list []any) bool {
	for _, cl := range list {
		if cl == value {
			return true
		}
	}

	return false
}
