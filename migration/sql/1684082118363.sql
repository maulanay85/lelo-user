create table t_mst_village(
    id bigserial,
    name varchar(50) not null,
    district_id integer not null,
    is_deleted boolean default false,
    create_by integer,
    update_by integer,
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp
);

create index idx_district_id_name on public.t_mst_village(district_id, name);

create table t_mst_district(
    id bigserial,
    name varchar(50) not null,
    city_id integer not null,
    is_deleted boolean default false,
    create_by integer,
    update_by integer,
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp
);

create index idx_city_id_name on public.t_mst_district(city_id, name);

create table t_mst_city(
    id bigserial,
    name varchar(50) not null,
    province_id integer not null,
    is_deleted boolean default false,
    create_by integer,
    update_by integer,
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp
);

create index idx_province_id_name on public.t_mst_city(province_id, name);


create table t_mst_province(
    id bigserial,
    name varchar(50) not null,
    is_deleted boolean default false,
    create_by integer,
    update_by integer,
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp
);

create index idx_id_name on public.t_mst_province(name);

create table t_mst_user_address(
    id bigserial,
    user_id integer not null,
    province_id integer not null,
    city_id integer not null,
    district_id integer not null,
    village_id integer not null,
    zip_code varchar(32),
    lat float,
    long float,
    is_main boolean default false,
    create_by integer,
    update_by integer,
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp
);

create index idx_user_id on public.t_mst_user_address(user_id);