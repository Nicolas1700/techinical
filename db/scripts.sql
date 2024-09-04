CREATE TABLE IF NOT EXISTS users
(
    id_user character varying(50) NOT NULL PRIMARY KEY,
    name_user character varying(50) NOT NULL,
    cell_phone bigint NOT NULL
);

COMMENT ON COLUMN users.id_user IS 'Identificador único del usuario';
COMMENT ON COLUMN users.name_user IS 'Nombre del usuario';
COMMENT ON COLUMN users.cell_phone IS 'Número de teléfono del usuario';

CREATE TABLE IF NOT EXISTS video
(
    id_video varchar(50) NOT NULL PRIMARY KEY,
    id_user varchar(50) NOT NULL,
    name_video varchar(50) NOT NULL,
    url_video varchar(255) NOT NULL,

    CONSTRAINT fk_video_user FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

COMMENT ON COLUMN video.id_video IS 'Identificador único del video';
COMMENT ON COLUMN video.id_user IS 'Identificador del usuario que subió el video';
COMMENT ON COLUMN video.name_video IS 'Nombre del video';
COMMENT ON COLUMN video.url_video IS 'URL del video';

CREATE TABLE IF NOT EXISTS challenges
(
    id_challenge varchar(50) NOT NULL PRIMARY KEY,
    id_video varchar(50) NOT NULL,
    name_challenge varchar(50) NOT NULL,
    number_participants BIGINT NOT NULL,

    CONSTRAINT fk_challenge_video FOREIGN KEY (id_video) REFERENCES video(id_video) ON DELETE CASCADE
);

COMMENT ON COLUMN challenges.id_challenge IS 'Identificador único del desafío';
COMMENT ON COLUMN challenges.id_video IS 'Identificador del video asociado al desafío';
COMMENT ON COLUMN challenges.name_challenge IS 'Nombre del desafío';
COMMENT ON COLUMN challenges.number_participants IS 'Número de participantes en el desafío';
