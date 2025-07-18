-- name: CreateProduct :one
insert into product (name, price, stock, brand_id)
values ($1, $2, $3, $4)
returning *;

-- name: GetProduct :one
select * from product
where product_id = $1
and deleted_at is null;

-- name: UpdateProduct :one
update product
set name = $2, price = $3, stock = $4, brand_id = $5
where product_id = $1
returning *;

-- name: DeleteProduct :exec
update product
set deleted_at = now()
where product_id = $1;
