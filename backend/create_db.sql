DO
$$
BEGIN
    -- Create role if it doesn't exist
    IF NOT EXISTS (
        SELECT 1 FROM pg_catalog.pg_roles WHERE rolname = 'iithpool'
    ) THEN
        EXECUTE format(
            'CREATE ROLE %I WITH LOGIN PASSWORD %L',
            'iithpool', 'password123'
        );
END IF;

    -- Create database if it doesn't exist
    IF NOT EXISTS (
        SELECT 1 FROM pg_database WHERE datname = 'CanteenPool'
    ) THEN
        EXECUTE format(
            'CREATE DATABASE %I OWNER %I',
            'CanteenPool', 'iithpool'
        );
END IF;
END
$$ LANGUAGE plpgsql;
