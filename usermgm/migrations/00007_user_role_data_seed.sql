-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO roles (id, name, description, status)
VALUES 
        ('d8f68815-442c-4962-b92c-9cb378042eae','Super Admin', 'Super Admin',1),  
        ('aad8c4f9-1fbe-4ae9-a537-b85258ae9161','Admin', 'Admin',1),  
        ('c6b4d4ba-9eab-4842-90f9-58806dbb0719','User', 'User',1);

INSERT INTO user_role (user_id, role_id)
VALUES 
        ('b6ddbe32-3d7e-4828-b2d7-da9927846e6b','d8f68815-442c-4962-b92c-9cb378042eae'),  
        ('bdda8d90-be00-46ee-a73f-98cd4b38d186','aad8c4f9-1fbe-4ae9-a537-b85258ae9161'),  
        ('662af210-5370-448c-96d7-d8c40d7b23d8','c6b4d4ba-9eab-4842-90f9-58806dbb0719');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
