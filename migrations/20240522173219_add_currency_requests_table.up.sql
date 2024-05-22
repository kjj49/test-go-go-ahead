CREATE TABLE IF NOT EXISTS currency_requests(
    id SERIAL PRIMARY KEY,
    val VARCHAR(3) NOT NULL,
    date DATE NOT NULL,
    request_date DATE NOT NULL
);
