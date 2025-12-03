CREATE TABLE bioskop (
     id SERIAL PRIMARY KEY,
     nama VARCHAR(50) NOT NULL,
     lokasi VARCHAR(100) NOT NULL,
     rating NUMERIC(3, 1) CONSTRAINT rating_check CHECK (rating >= 1.0 AND rating <= 10.0)
);