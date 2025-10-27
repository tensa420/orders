-- +goose Up

CREATE TABLE Orders (
    order_uuid PRIMARY KEY,
    user_uuid UUID not null,
    part_uuids []TEXT not null,
    total_price DECIMAL(10,3) NOT NULL,
    transaction_uuid UUID NOT NULL,
    payment_method TEXT NOT NULL,
    STATUS TEXT NOT NULL
);
-- +goose Down
drop table NOTE;