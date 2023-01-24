CREATE TABLE TRACKER.EXERCISE_DETAILS (
    EXERCISE_NAME VARCHAR(50) NOT NULL REFERENCES TRACKER.EXERCISE(NAME) ON DELETE CASCADE,
    BODY_PART VARCHAR(50) NOT NULL,
    LEVEL VARCHAR(10) NOT NULL,
    CRET_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UPDT_TS TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY(EXERCISE_NAME, BODY_PART)
);