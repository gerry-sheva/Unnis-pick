-- +goose Up
-- +goose StatementBegin
create table product (
    product_id uuid primary key default uuid_generate_v4(),
    name text collate "case_insensitive" unique not null,
    description text not null,
    price numeric(10, 2) not null,
    stock int not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz,

    brand_id uuid references brand(brand_id) not null
);

select trigger_updated_at('product');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists product;
-- +goose StatementEnd
