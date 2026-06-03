IF OBJECT_ID('core.permissions', 'U') IS NULL
BEGIN
    CREATE TABLE core.permissions (
        id            UNIQUEIDENTIFIER PRIMARY KEY,
        group_id      UNIQUEIDENTIFIER NOT NULL,
        name          NVARCHAR(255) NOT NULL,
        created_at    DATETIME NOT NULL DEFAULT GETDATE(),
        last_modified DATETIME NOT NULL DEFAULT GETDATE(),
        CONSTRAINT FK_permissions_group FOREIGN KEY (group_id) REFERENCES core.groups(id) ON DELETE CASCADE
    );
END
