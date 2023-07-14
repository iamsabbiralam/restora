-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO users (id, username, email, password, image, status, created_at, created_by, updated_at, updated_by)
VALUES 
        ('b6ddbe32-3d7e-4828-b2d7-da9927846e6b','superadmin', 'superadmin@gmail.com','$2a$10$JfrJhMaA34LPMbKl8G6Kiu0Q3EtKZnyVvehwHSn8mFX2eLXi7cVgy', 'default.jpg', 1, '2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b','2021-12-29 06:30:19.7526', 'b6ddbe32-3d7e-4828-b2d7-da9927846e6b'),
        ('662af210-5370-448c-96d7-d8c40d7b23d8', 'user', 'user@gmail.com','$2a$10$JfrJhMaA34LPMbKl8G6Kiu0Q3EtKZnyVvehwHSn8mFX2eLXi7cVgy', 'default.jpg', 1, '2021-12-29 06:30:19.7526', '662af210-5370-448c-96d7-d8c40d7b23d8','2021-12-29 06:30:19.7526', 'b662af210-5370-448c-96d7-d8c40d7b23d8') ON CONFLICT (username) DO NOTHING;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
