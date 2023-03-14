CREATE TABLE IF NOT EXISTS ticket_orders (
  id bigserial PRIMARY KEY,
  ticket_id BIGINT,
  user_id BIGINT,
  seat_id BIGINT,
  purchased_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);