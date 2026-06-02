IF OBJECT_ID('dbo.[users]', 'U') IS NULL
BEGIN
    CREATE TABLE [users] (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        username NVARCHAR(255) NOT NULL UNIQUE,
        email NVARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL DEFAULT GETDATE(),
        updated_at DATETIME NULL
    )   
END