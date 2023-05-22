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

func (lc *LocationRepoistoryModule) GetCityByProvinceId(ctx context.Context, provinceId int64, name string) ([]entity.CityEntity, error) {
	var cities []entity.CityEntity
	sql := `
		SELECT
			id, 
			name,
			province_id,
			is_deleted,
			created_time,
			updated_time
		FROM 
			t_mst_city
		WHERE
			is_deleted = false
			%s
	`
	query, queryValue := helperGetCity(int16(provinceId), name)
	sql = fmt.Sprintf(sql, query)

	rows, err := lc.db.Query(ctx, sql, queryValue...)
	if err != nil {
		log.Errorf("[repository]: GetCity err: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		city := entity.CityEntity{}
		err := rows.Scan(&city.Id, &city.Name, &city.ProvinceId, &city.IsDeleted, &city.CreatedTime, &city.UpdatedTime)
		if err != nil {
			log.Errorf("[repository]: GetCity err scan rows: %v", err)
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}

func (lc *LocationRepoistoryModule) GetDistrictByCityId(ctx context.Context, cityId int64, name string) ([]entity.DistrictEntity, error) {
	var distrcits []entity.DistrictEntity
	sql := `
		SELECT
			id, 
			name,
			city_id,
			is_deleted,
			created_time,
			updated_time
		FROM 
			t_mst_district
		WHERE
			is_deleted = false
			%s
	`
	query, queryValue := helperGetDistrict(int16(cityId), name)
	sql = fmt.Sprintf(sql, query)

	rows, err := lc.db.Query(ctx, sql, queryValue...)
	if err != nil {
		log.Errorf("[repository]: GetDistrictByCityId err: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		district := entity.DistrictEntity{}
		err := rows.Scan(&district.Id, &district.Name, &district.CityId, &district.IsDeleted, &district.CreatedTime, &district.UpdatedTime)
		if err != nil {
			log.Errorf("[repository]: Getdistrict err scan rows: %v", err)
			return nil, err
		}
		distrcits = append(distrcits, district)
	}
	return distrcits, nil
}

func (lc *LocationRepoistoryModule) GetVillageByDistrictId(ctx context.Context, districtId int64, name string) ([]entity.VillageEntity, error) {
	var villages []entity.VillageEntity
	sql := `
		SELECT
			id, 
			name,
			district_id,
			is_deleted,
			created_time,
			updated_time
		FROM 
			t_mst_village
		WHERE
			is_deleted = false
			%s
	`
	query, queryValue := helperGetVillage(districtId, name)
	sql = fmt.Sprintf(sql, query)

	rows, err := lc.db.Query(ctx, sql, queryValue...)
	if err != nil {
		log.Errorf("[repository]: GetDistrictByCityId err: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		village := entity.VillageEntity{}
		err := rows.Scan(&village.Id, &village.Name, &village.DistrictId, &village.IsDeleted, &village.CreatedTime, &village.UpdatedTime)
		if err != nil {
			log.Errorf("[repository]: Getdistrict err scan rows: %v", err)
			return nil, err
		}
		villages = append(villages, village)
	}
	return villages, nil
}

func helperGetCity(id int16, name string) (query string, value []interface{}) {
	parameter := []string{}
	parameterValue := []interface{}{}

	var queryResult string

	if id != 0 {
		parameter = append(parameter, "province_id = $%d")
		parameterValue = append(parameterValue, id)
	}
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
	return queryResult, parameterValue

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

func helperGetDistrict(id int16, name string) (query string, value []interface{}) {
	parameter := []string{}
	parameterValue := []interface{}{}

	var queryResult string

	if id != 0 {
		parameter = append(parameter, "city_id = $%d")
		parameterValue = append(parameterValue, id)
	}
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
	return queryResult, parameterValue

}

func helperGetVillage(id int64, name string) (query string, value []interface{}) {
	parameter := []string{}
	parameterValue := []interface{}{}

	var queryResult string

	if id != 0 {
		parameter = append(parameter, "district_id = $%d")
		parameterValue = append(parameterValue, id)
	}
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
	return queryResult, parameterValue

}
