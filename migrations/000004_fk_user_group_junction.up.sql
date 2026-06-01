IF OBJECT_ID('dbo.user_group_junctions', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.user_group_junctions (
        user_id  UNIQUEIDENTIFIER NOT NULL,
        group_id UNIQUEIDENTIFIER NOT NULL,
        CONSTRAINT PK_user_group_junctions PRIMARY KEY (user_id, group_id),
        CONSTRAINT FK_ugj_user  FOREIGN KEY (user_id)  REFERENCES dbo.[User](id)  ON DELETE CASCADE,
        CONSTRAINT FK_ugj_group FOREIGN KEY (group_id) REFERENCES dbo.Groups(id) ON DELETE CASCADE
    );
END
