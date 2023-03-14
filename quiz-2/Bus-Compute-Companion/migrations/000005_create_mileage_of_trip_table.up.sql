CREATE TABLE IF NOT EXISTS mileage_of_trip (
  id bigserial PRIMARY KEY,
  beginning_location_id BIGINT,
  destination_location_id BIGINT,
  total_miles BIGINT
);