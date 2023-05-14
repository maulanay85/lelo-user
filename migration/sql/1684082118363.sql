create table t_mst_village(
    id bigserial primary key,
    name varchar(50) not null,
    district_id integer not null,
    is_deleted boolean default false,
    created_by integer,
    updated_by integer,
    created_time timestamp default current_timestamp,
    updated_time timestamp default current_timestamp
);

create index idx_district_id_name on public.t_mst_village(district_id, name);

create table t_mst_district(
    id bigserial primary key,
    name varchar(50) not null,
    city_id integer not null,
    is_deleted boolean default false,
    created_by integer,
    updated_by integer,
    created_time timestamp default current_timestamp,
    updated_time timestamp default current_timestamp
);

create index idx_city_id_name on public.t_mst_district(city_id, name);

create table t_mst_city(
    id bigserial primary key,
    name varchar(50) not null,
    province_id integer not null,
    is_deleted boolean default false,
    created_by integer,
    updated_by integer,
    created_time timestamp default current_timestamp,
    updated_time timestamp default current_timestamp
);

create index idx_province_id_name on public.t_mst_city(province_id, name);


create table t_mst_province(
    id bigserial primary key,
    name varchar(50) not null,
    is_deleted boolean default false,
    created_by integer,
    updated_by integer,
    created_time timestamp default current_timestamp,
    updated_time timestamp default current_timestamp
);

create index idx_id_name on public.t_mst_province(name);

create table t_mst_user_address(
    id bigserial primary key,
    user_id integer not null,
    province_id integer not null,
    city_id integer not null,
    district_id integer not null,
    village_id integer not null,
    zip_code varchar(32),
    lat float,
    long float,
    is_main boolean default false,
    created_by integer,
    updated_by integer,
    created_time timestamp default current_timestamp,
    updated_time timestamp default current_timestamp
);

create index idx_user_id on public.t_mst_user_address(user_id);
