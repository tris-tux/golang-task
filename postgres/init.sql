DROP TABLE IF EXISTS task;
CREATE SEQUENCE task_id START 1;
CREATE TABLE task (
    TASK_ID serial PRIMARY KEY,
    TITLE VARCHAR NOT NULL,
    ACCTION_TIME INT NOT NULL,
    CREATE_TIME INT NOT NULL,
    UPDATE_TIME INT NOT NULL,
    IS_FINISHED BIT NOT NULL
);

DROP TABLE IF EXISTS detail;
CREATE SEQUENCE detail_id START 1;
CREATE TABLE detail (
    DETAIL_ID serial PRIMARY KEY,
    OBJECT_TASK_FK INT NOT NULL,
    OBJECT_NAME VARCHAR NOT NULL,
    IS_FINISHED BIT NOT NULL
);