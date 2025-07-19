-- name: CreateProduct :one
insert into product (name, description, price, stock, brand_id)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetProduct :one
select * from product
where product_id = $1
and deleted_at is null;

-- name: UpdateProduct :one
update product
set name = $2, description = $3, price = $4, stock = $5, brand_id = $6
where product_id = $1
returning *;

-- name: DeleteProduct :exec
update product
set deleted_at = now()
where product_id = $1;

-- name: IsBrandUsed :one
select exists(
    select 1 from product
    where brand_id = $1
    and deleted_at is null
);
