IF OBJECT_ID('core.user_group_junctions', 'U') IS NULL
BEGIN
    CREATE TABLE core.user_group_junctions (
        user_id  UNIQUEIDENTIFIER NOT NULL,
        group_id UNIQUEIDENTIFIER NOT NULL,
        CONSTRAINT PK_user_group_junctions PRIMARY KEY (user_id, group_id),
        CONSTRAINT FK_ugj_user  FOREIGN KEY (user_id)  REFERENCES core.users(id)  ON DELETE CASCADE,
        CONSTRAINT FK_ugj_group FOREIGN KEY (group_id) REFERENCES core.groups(id) ON DELETE CASCADE
    );
END
