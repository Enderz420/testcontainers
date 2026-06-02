IF OBJECT_ID('dbo.groups', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.groups (
        id            UNIQUEIDENTIFIER PRIMARY KEY,
        name          NVARCHAR(255) NOT NULL,
        created_at    DATETIME NOT NULL DEFAULT GETDATE(),
        last_modified DATETIME NOT NULL DEFAULT GETDATE()
    );
END