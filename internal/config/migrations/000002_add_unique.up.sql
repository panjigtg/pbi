ALTER TABLE user
ADD UNIQUE INDEX ux_user_email (email),
ADD UNIQUE INDEX ux_user_notelp (notelp);
