IF NOT EXISTS (SELECT 1 FROM core.groups)
BEGIN
    INSERT INTO core.groups (id, name, created_at, last_modified)
    VALUES
        (NEWID(), 'admin',  GETDATE(), GETDATE()),
        (NEWID(), 'moderator', GETDATE(), GETDATE()),
        (NEWID(), 'user', GETDATE(), GETDATE());
END
