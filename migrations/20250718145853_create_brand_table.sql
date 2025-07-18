-- +goose Up
-- +goose StatementBegin
create table brand (
    brand_id uuid primary key default uuid_generate_v1mc(),
    name text collate "case_insensitive" unique not null,
    created_at timestamptz default current_timestamp not null,
    updated_at timestamptz,
    deleted_at timestamptz
);
select trigger_updated_at('brand');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists brand;
-- +goose StatementEnd
