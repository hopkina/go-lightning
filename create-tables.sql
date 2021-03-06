DROP TABLE lightning_strikes;

CREATE TABLE lightning_strikes (
  id           SERIAL PRIMARY KEY,
  strike_time  TIMESTAMP NOT NULL,
  x_coord      NUMERIC NOT NULL,
  y_coord      NUMERIC NOT NULL
);

INSERT INTO 
	lightning_strikes(strike_time, x_coord, y_coord)
VALUES
  	('2021-11-10T23:00:00Z', -1.193025, 51.615335),
	('2021-11-10T21:00:00Z', -3.922570, 56.099230),
	('2021-11-10T16:00:00Z', -4.638846, 52.060492);
