IF OBJECT_ID('core.blogpost', 'U') IS NULL
BEGIN
    CREATE TABLE core.blogpost (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        title NVARCHAR(255) NOT NULL,
        content NVARCHAR(MAX),
        created_by NVARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL DEFAULT GETUTCDATE(),
        updated_at DATETIME NOT NULL DEFAULT GETUTCDATE(),

        -- CONSTRAINT FK_Blogpost_User FOREIGN KEY (created_by) REFERENCES [User](username)
    )
END
