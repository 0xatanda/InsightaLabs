CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    gender VARCHAR(10),
    gender_probability FLOAT,
    age INT,
    age_group VARCHAR(20),
    country_id VARCHAR(2),
    country_name VARCHAR(255),
    country_probability FLOAT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_gender ON profiles(gender);
CREATE INDEX idx_age ON profiles(age);
CREATE INDEX idx_country ON profiles(country_id);
CREATE INDEX idx_age_group ON profiles(age_group);
CREATE INDEX idx_created_at ON profiles(created_at);