package location

import (
	"context"
	"fmt"

	"lelo-user/entity"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (l *LocationRepoistoryModule) GetProvince(ctx context.Context, name string) ([]entity.ProvinceEntity, error) {
	var provinces []entity.ProvinceEntity

	sql := `SELECT
				id,
				name,
				is_deleted,
				created_time,
				updated_time
			FROM
				t_mst_province
			WHERE
				is_deleted = false
				%s
			`
	query, queryValue := helperGetProvince(name)
	sql = fmt.Sprintf(sql, query)

	rows, err := l.db.Query(ctx, sql, queryValue...)
	if err != nil {
		log.Errorf("[repository]: GetProvince err: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		province := entity.ProvinceEntity{}
		err := rows.Scan(&province.Id, &province.Name, &province.IsDeleted, &province.CreatedTime, &province.UpdatedTime)
		if err != nil {
			log.Errorf("[repository]: GetProvince err scan rows: %v", err)
		}
		provinces = append(provinces, province)

	}

	return provinces, nil
}

func helperGetProvince(name string) (query string, value []interface{}) {
	parameter := []string{}
	parameterValue := []interface{}{}

	var queryResult string

	if name != "" {
		parameter = append(parameter, "name ilike $%d")
		parameterValue = append(parameterValue, "%"+name+"%")
	}

	// generate
	if len(parameter) > 0 {
		queryResult += " and "

		for i := 0; i < len(parameter); i++ {
			parameter[i] = fmt.Sprintf(parameter[i], i+1)
		}

		queryResult += strings.Join(parameter, " and ")
	}
	fmt.Print(queryResult)
	return queryResult, parameterValue

}
