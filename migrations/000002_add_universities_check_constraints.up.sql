ALTER TABLE
    universities
ADD
    CONSTRAINT universities_date_founded_check CHECK (
        extract (
            year
            FROM
                founded
        ) BETWEEN 1589
        AND EXTRACT(
            YEAR
            FROM
                CURRENT_DATE
        )
    );