IF OBJECT_ID('dbo.Blogpost', 'U') IS NULL
BEGIN
    CREATE TABLE Blogpost (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        title NVARCHAR(255) NOT NULL,
        content NVARCHAR(MAX),
        created_by UNIQUEIDENTIFIER NOT NULL,
        created_at DATETIME2 NOT NULL DEFAULT GETDATE(),
        updated_at DATETIME2 NULL,

        CONSTRAINT FK_Blogpost_User FOREIGN KEY (created_by) REFERENCES [User](id)
    )
END