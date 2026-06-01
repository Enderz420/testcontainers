IF OBJECT_ID('dbo.[User]', 'U') IS NULL
BEGIN
    CREATE TABLE [User] (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        username NVARCHAR(255) NOT NULL UNIQUE,
        email NVARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL DEFAULT GETDATE(),
        updated_at DATETIME NULL
    )   
END