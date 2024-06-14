CREATE TABLE os (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_os_name
ON os (name);

INSERT INTO os VALUES
(0, 'ubuntu');

CREATE TABLE architecture (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_architecture_name
ON architecture (name);

INSERT INTO architecture VALUES
(0, 'amd64'),
(1, 'arm64'),
(2, 'ppc64el'),
(3, 's390x'),
(4, 'riscv64');
