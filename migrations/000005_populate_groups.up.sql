IF NOT EXISTS (SELECT 1 FROM core.groups)
BEGIN
    INSERT INTO core.groups (id, name, created_at, last_modified)
    VALUES
        (NEWID(), 'admin',  GETUTCDATE(), GETUTCDATE()),
        (NEWID(), 'moderator', GETUTCDATE(), GETUTCDATE()),
        (NEWID(), 'user', GETUTCDATE(), GETUTCDATE());
END
