IF OBJECT_ID('dbo.blogpost', 'U') IS NULL
BEGIN
    CREATE TABLE blogpost (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        title NVARCHAR(255) NOT NULL,
        content NVARCHAR(MAX),
        created_by NVARCHAR(255) NOT NULL,
        created_at DATETIME NOT NULL DEFAULT GETDATE(),
        updated_at DATETIME NULL,

        -- CONSTRAINT FK_Blogpost_User FOREIGN KEY (created_by) REFERENCES [User](username)
    )
END