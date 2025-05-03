CREATE DATABASE TestDB;
GO

USE TestDB;

CREATE TABLE Device (
    id INT IDENTITY(1,1) PRIMARY KEY,
    /* probably too big */
    name NVARCHAR(255),
    status NVARCHAR(255),
    totalMemory NVARCHAR(255),
    memoryUsage NVARCHAR(255),
    memWarnThres NVARCHAR(255),
    cpuCount NVARCHAR(255),
    cpuUsage NVARCHAR(255),
    cpuWarnThres NVARCHAR(255)
)

/* Device 1 - no warns */
INSERT INTO Device (name, status, totalMemory, memoryUsage, memWarnThres, cpuCount, cpuUsage, cpuWarnThres)
VALUES ('Device 1', 'Online', '8192MB', '42%', 'No', '8', '62%', 'No');
GO

/* Device 2 - mem warn */
INSERT INTO Device (name, status, totalMemory, memoryUsage, memWarnThres, cpuCount, cpuUsage, cpuWarnThres)
VALUES ('Device 2', 'Online', '16382MB', '99%', 'Yes', '16', '22%', 'No');
GO

/* Device 3 - cpu warn */
INSERT INTO Device (name, status, totalMemory, memoryUsage, memWarnThres, cpuCount, cpuUsage, cpuWarnThres)
VALUES ('Device 3', 'Online', '4092MB', '42%', 'No', '4', '82%', 'Yes');
GO

/* Device 4 - both warn */
INSERT INTO Device (name, status, totalMemory, memoryUsage, memWarnThres, cpuCount, cpuUsage, cpuWarnThres)
VALUES ('Device 4', 'Online', '8192MB', '98%', 'Yes', '8', '89%', 'Yes');
GO