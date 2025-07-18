-- name: CreateBrand :one
insert into brand (name)
values ($1)
returning *;

-- name: GetBrand :one
select * from brand
where brand_id = $1
and deleted_at is null;

-- name: UpdateBrand :one
update brand set name = $2
where brand_id = $1
returning *;

-- name: DeleteBrand :exec
update brand set deleted_at = now()
where brand_id = $1;
