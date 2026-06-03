IF NOT EXISTS (SELECT 1 FROM core.groups)
BEGIN
    INSERT INTO core.groups (id, name, created_at, last_modified)
    VALUES
        (NEWID(), 'Admin',  GETDATE(), GETDATE()),
        (NEWID(), 'Editor', GETDATE(), GETDATE()),
        (NEWID(), 'Viewer', GETDATE(), GETDATE());
END
