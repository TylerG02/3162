CREATE TABLE IF NOT EXISTS bus_schedule (
  id bigserial PRIMARY KEY AUTO_INCREMENT,
  company_id BIGINT,
  beginning_location_id BIGINT,
  destination_location_id BIGINT,
  bus_departure_time TIME,
  bus_arrival_time TIME
);