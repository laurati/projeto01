SELECT 'CREATE DATABASE default_database'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'default_database')\gexec