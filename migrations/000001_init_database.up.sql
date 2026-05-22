IF OBJECT_ID('dbo.[User]', 'U') IS NULL
BEGIN
    CREATE TABLE [User] (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        created_at DATETIME2 NOT NULL DEFAULT GETDATE(),
        updated_at DATETIME2 NULL,
    )   
END