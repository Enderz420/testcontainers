IF OBJECT_ID('core.groups', 'U') IS NULL
BEGIN
    CREATE TABLE core.groups (
        id            UNIQUEIDENTIFIER PRIMARY KEY,
        name          NVARCHAR(255) NOT NULL,
        created_at    DATETIME NOT NULL DEFAULT GETDATE(),
        last_modified DATETIME NOT NULL DEFAULT GETDATE()
    );
END