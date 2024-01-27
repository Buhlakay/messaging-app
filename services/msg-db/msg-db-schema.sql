CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  sender_id INT NOT NULL,
  receiver_id INT NOT NULL,
  body TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (sender_id) REFERENCES users(id),
  FOREIGN KEY (receiver_id) REFERENCES users(id)
);
