-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO users (id, username, email, password, status, created_at, created_by, updated_at, updated_by)
VALUES 
        ('b6ddbe32-3d7e-4828-b2d7-da9927846e6b','superadmin', 'superadmin@gmail.com','$2a$10$JfrJhMaA34LPMbKl8G6Kiu0Q3EtKZnyVvehwHSn8mFX2eLXi7cVgy', 1, '2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b','2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b'),
        ('662af210-5370-448c-96d7-d8c40d7b23d8', 'user', 'user@gmail.com','$2a$10$JfrJhMaA34LPMbKl8G6Kiu0Q3EtKZnyVvehwHSn8mFX2eLXi7cVgy', 1, '2021-12-29 06:30:19.7526', '662af210-5370-448c-96d7-d8c40d7b23d8','2021-12-29 06:30:19.7526', 'b662af210-5370-448c-96d7-d8c40d7b23d8') ON CONFLICT (username) DO NOTHING;

INSERT INTO user_information (id, user_id, first_name, last_name, image, mobile, gender, dob, address, city, country, created_at, created_by, updated_at, updated_by)
VALUES 
        ('b6ddbe32-3d7e-4828-b2d7-ju73he7s73er', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b', 'Super', 'Admin', 'default.jpg','+8801715039303', 1, '1995-08-28', 'farazipara', 'khulna', 'bangladesh', '2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b','2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b'),
        ('662af210-5370-448c-96d7-d8c40d7bv34s', '662af210-5370-448c-96d7-d8c40d7b23d8', 'John', 'Dou', 'default.jpg','+8801715039309', 1, '1995-08-28', 'farazipara', 'khulna', 'bangladesh', '2021-12-29 06:30:19.7526', '662af210-5370-448c-96d7-d8c40d7b23d8','2021-12-29 06:30:19.7526', 'b662af210-5370-448c-96d7-d8c40d7b23d8');


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
